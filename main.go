package main

import (
	"fmt"
	"os"
	"time"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/joho/godotenv/autoload"

	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

var slackapi = slack.New(
	os.Getenv("SLACK_BOT_TOKEN"),
	slack.OptionAppLevelToken(os.Getenv("SLACK_APP_TOKEN")),
)

var fifaapi = &fifa.Client{}

var matches = map[string]*Match{}
var seen = map[string]bool{}

type Match struct {
	Competition   string
	CompetitionId string
	SeasonId      string
	StageId       string
	MatchId       string
	HomeTeam      string
	HomeFlag      string
	AwayTeam      string
	AwayFlag      string
	SlackThreadTs string
	LastEventTs   time.Time
	LastEventType int
	LastEventId   int
	IsOver        bool
}

var teams = map[string]Team{}

type Team struct {
	Id   string
	Name string
	Flag string
}

var players = map[string]Player{}

type Player struct {
	Id       string
	Name     string
	LastName string
	Team     Team
}

var competitionIdFilter string

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	socket := socketmode.New(slackapi)
	if _, err := slackapi.AuthTest(); err != nil {
		logrus.Errorf("SLACK_BOT_TOKEN is invalid: %v", err)
		os.Exit(1)
	}

	competitionIdFilter = os.Getenv("COMPETITION_ID_FILTER")

	go socket.Run()

	eventLoop()
}

func addTeam(data fifa.TeamResponse) {
	team := Team{
		Id:   data.Players[0].TeamId,
		Name: data.Name[0].Description,
		Flag: countries[data.CountryId].Flag,
	}

	for _, player := range data.Players {
		players[player.Id] = Player{
			Id:       player.Id,
			Name:     player.Name[0].Description,
			LastName: player.ShortName[0].Description,
			Team:     team,
		}
	}

	teams[team.Id] = team
}

func eventLoop() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	for {
		checkForNewMatches()
		checkMatchEvents()

		select {
		case <-ticker.C:
			// nothing
		case <-quit:
			return
		}
	}
}

func checkForNewMatches() {
	current, err := fifaapi.GetCurrentMatches()
	if err != nil {
		logrus.Debug(err)
	}
	for _, m := range current {
		if competitionIdFilter != "" && m.CompetitionId != competitionIdFilter {
			continue
		}

		hash := fmt.Sprintf("%s-%s-%s-%s", m.CompetitionId, m.SeasonId, m.StageId, m.Id)

		if _, ok := matches[hash]; ok {
			continue
		}

		_, respTs, _ := slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":thread: *%s* %s 0 - 0 %s *%s* (_%s_) - Starting soon!", m.HomeTeam.Name[0].Description, countries[m.HomeTeam.CountryId].Flag, countries[m.AwayTeam.CountryId].Flag, m.AwayTeam.Name[0].Description, m.Competition[0].Description),
				false,
			),
		)

		logrus.Debugf("New match found for *%s* v *%s* (_%s_ / %s)\n", m.HomeTeam.Name[0].Description, m.AwayTeam.Name[0].Description, m.Competition[0].Description, m.CompetitionId)

		matches[hash] = &Match{
			Competition:   m.Competition[0].Description,
			CompetitionId: m.CompetitionId,
			SeasonId:      m.SeasonId,
			StageId:       m.StageId,
			MatchId:       m.Id,
			HomeTeam:      m.HomeTeam.Name[0].Description,
			HomeFlag:      countries[m.HomeTeam.CountryId].Flag,
			AwayTeam:      m.AwayTeam.Name[0].Description,
			AwayFlag:      countries[m.AwayTeam.CountryId].Flag,
			SlackThreadTs: respTs,
		}

		if data, err := fifaapi.GetMatchData(&fifa.GetMatchDataOptions{
			CompetitionId: m.CompetitionId,
			SeasonId:      m.SeasonId,
			StageId:       m.StageId,
			MatchId:       m.Id,
		}); err == nil {
			addTeam(data.HomeTeam)
			addTeam(data.AwayTeam)
		}
	}
}

func checkMatchEvents() {
	for _, match := range matches {
		if match.IsOver {
			continue
		}

		if match.SlackThreadTs == "" {
			_, respTs, _ := slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
				slack.MsgOptionText(
					fmt.Sprintf(":thread: *%s* %s 0 - 0 %s *%s* (_%s_)", match.HomeTeam, match.HomeFlag, match.AwayFlag, match.AwayTeam, match.Competition),
					false,
				),
			)
			match.SlackThreadTs = respTs
		}

		logrus.Debugf("Checking for updates for *%s* v *%s* (_%s_ / %s)\n", match.HomeTeam, match.AwayTeam, match.Competition, match.CompetitionId)

		events, err := fifaapi.GetMatchEvents(&fifa.GetMatchEventOptions{
			CompetitionId: match.CompetitionId,
			SeasonId:      match.SeasonId,
			StageId:       match.StageId,
			MatchId:       match.MatchId,
		})

		if err != nil {
			continue
		}

		for _, event := range events.Events {
			if _, seen := seen[fmt.Sprintf("%s-%d-%s", event.Id, event.Type, event.Timestamp)]; seen {
				continue
			}

			if (event.Timestamp.After(match.LastEventTs) || event.Timestamp.Equal(match.LastEventTs)) && match.LastEventType != int(event.Type) {
				if event.Type == 9999 {
					return
				}

				seen[fmt.Sprintf("%s-%d-%s", event.Id, event.Type, event.Timestamp)] = true

				match.LastEventTs = event.Timestamp
				match.LastEventType = int(event.Type)

				if method, ok := eventMap[event.Type]; ok {
					if method != nil {
						method(match, event)
					}
				} else {
					spew.Dump(event)
				}

				var flavorMessage string
				switch event.Period {
				case fifa.FIRST:
					flavorMessage = "In First Half"
				case fifa.FIRST_EXTRA:
					flavorMessage = "In extended First Half"
				case fifa.SECOND:
					flavorMessage = "In Second Half"
				case fifa.SECOND_EXTRA:
					flavorMessage = "In extended Second Half"
				case fifa.SHOOTOUT:
					flavorMessage = "In shootout!"
				}

				if match.IsOver {
					flavorMessage = "Game has finished!"
				}

				slackapi.UpdateMessage(
					os.Getenv("SLACK_OUTPUT_CHANNEL"),
					match.SlackThreadTs,
					slack.MsgOptionText(
						fmt.Sprintf(":thread: *%s* %s %d - %d %s *%s* (_%s_) - %s", match.HomeTeam, match.HomeFlag, event.HomeGoals, event.AwayGoals, match.AwayFlag, match.AwayTeam, match.Competition, flavorMessage),
						false,
					),
				)
			}
		}
	}
}

func SendSlackMessage(match *Match, message string, args ...interface{}) {
	_, _, err := slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
		slack.MsgOptionText(
			fmt.Sprintf(message, args...),
			false,
		),
		slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
			ThreadTimestamp: match.SlackThreadTs,
		}),
	)

	if err != nil {
		logrus.Errorf("Unable to post message to Slack: %v", err)
	}
}
