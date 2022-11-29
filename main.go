package main

import (
	"encoding/json"
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

var eventMap = map[fifa.MatchEvent]func(match *Match, event fifa.EventResponse){
	fifa.GoalScore:  doGoalScore,
	fifa.Assist:     nil,
	fifa.YellowCard: doYellowCard,
	fifa.RedCard:    doRedCard,
	/*
		4:  "DoubleYellow",
	*/
	fifa.Substitution:   doSubstitution,
	fifa.PenaltyAwarded: doAwardPenalty,
	fifa.MatchStart:     doMatchStart,
	fifa.HalfEnd:        doHalfEnd,
	/*
		9:  "MatchPaused",
		10: "MatchResumed",
	*/
	fifa.GoalAttempt: doGoalAttempt,
	fifa.FoulUnknown: doFoul,
	fifa.Offside:     doOffside,
	fifa.CornerKick:  doCorner,
	/*
		17: "ShotBlocked",
	*/
	fifa.Foul:        doFoul,
	fifa.CoinToss:    nil,
	20:               nil, // Unknown?
	fifa.DroppedBall: nil,
	fifa.ThrowIn:     nil,
	fifa.Clearance:   nil,
	fifa.MatchEnd:    doMatchEnd,
	27:               nil, // Aeriel Duel
	/*
		32: "CrossBar",
		33: "CrossBar2",
		34: "OwnGoal",
		37: "HandBall",
	*/
	fifa.FreeKickGoal: doGoalScore,
	fifa.PenaltyGoal:  doGoalScore,
	/*
		44: "FreeKickCrossbar",
		49: "FreeKickPost",
	*/
	fifa.PenaltyMissed:  doPenaltyMissed,
	fifa.PenaltyMissed2: doPenaltyMissed,
	/*
		57: "GoalieSaved",
		72: "VARPenalty",

	*/
	71: doVAR,
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	socket := socketmode.New(slackapi)
	if _, err := slackapi.AuthTest(); err != nil {
		fmt.Fprintf(os.Stderr, "SLACK_BOT_TOKEN is invalid: %v", err)
		os.Exit(1)
	}

	go socket.Run()

	worker()
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

func worker() {
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
		fmt.Println(err)
	}
	for _, m := range current {
		if m.CompetitionId != "17" {
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

		fmt.Printf("Checking for updates for *%s* v *%s* (_%s_)\n", match.HomeTeam, match.AwayTeam, match.Competition)

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
			if event.Timestamp.After(match.LastEventTs) {
				if event.Type == 9999 {
					return
				}

				match.LastEventTs = event.Timestamp

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

func sendError(message string, event fifa.EventResponse) {
	data, _ := json.MarshalIndent(event, "", "  ")

	slackapi.PostMessage(os.Getenv("SLACK_ADMIN_USER_ID"),
		slack.MsgOptionText(fmt.Sprintf("%s\n```%s```", message, data), false),
	)
}

func doGoalScore(match *Match, event fifa.EventResponse) {
	fmt.Printf("doGoalScore: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":soccer: *%s* - GOOOOOOOOOOOOOAAAAAAAAAAAAAAAAAAALLLLLLLLLLLLLLLL!!!! *%s* %s scores a goal. The score is now *%s* %d - %d *%s*.", event.MatchMinute, player.Name, player.Team.Flag, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	} else {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":soccer: *%s* - GOOOOOOOOOOOOOAAAAAAAAAAAAAAAAAAALLLLLLLLLLLLLLLL!!!! %s %s scores a goal. The score is now *%s* %d - %d *%s*.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}

	go FindInstantReplay(match, event)
}

func doYellowCard(match *Match, event fifa.EventResponse) {
	fmt.Printf("doYellowCard: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee-with-yellow-card: *%s* - *%s* %s is booked with a yellow card.", event.MatchMinute, player.Name, player.Team.Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	} else {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee-with-yellow-card: *%s* - %s %s receives a yellow card.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}
}

func doRedCard(match *Match, event fifa.EventResponse) {
	fmt.Printf("doRedCard: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee-with-red-card: *%s* - *%s* %s is sent off with a red card.", event.MatchMinute, player.Name, player.Team.Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	} else {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee-with-red-card: *%s* - *%s* %s receives a red card.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}
}

func doSubstitution(match *Match, event fifa.EventResponse) {
	fmt.Printf("doSubstitution: %s : %s\n", event.Id, event.Timestamp)
	var playerIn Player
	var playerOut Player
	var ok bool

	if playerIn, ok = players[event.PlayerId]; ok {
		if playerOut, ok = players[event.SubPlayerId]; ok {
			slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
				slack.MsgOptionText(
					fmt.Sprintf(":arrows_counterclockwise: *%s* - *%s* %s comes in for *%s* %s.", event.MatchMinute, playerIn.Name, playerIn.Team.Flag, playerOut.Name, playerOut.Team.Flag),
					false,
				),
				slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
					ThreadTimestamp: match.SlackThreadTs,
				}),
			)
		} else {
			slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
				slack.MsgOptionText(
					fmt.Sprintf(":arrows_counterclockwise: *%s* - *%s* %s is substituted.", event.MatchMinute, playerIn.Name, playerIn.Team.Flag),
					false,
				),
				slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
					ThreadTimestamp: match.SlackThreadTs,
				}),
			)
		}

	} else {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":arrows_counterclockwise: *%s* - *%s* %s substitutes a player.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}
}

func doAwardPenalty(match *Match, event fifa.EventResponse) {
	fmt.Printf("doAwardPenalty: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - *%s* %s receives a penalty.", event.MatchMinute, player.Name, player.Team.Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	} else {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - *%s* %s receives a penalty.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}
}

func doMatchStart(match *Match, event fifa.EventResponse) {
	fmt.Printf("doMatchStart: %s : %s\n", event.Id, event.Timestamp)
	switch event.Period {
	case fifa.FIRST:
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - The match has started!", event.MatchMinute),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	case fifa.SECOND:
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - Start of the second half.", event.MatchMinute),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}
}

func doHalfEnd(match *Match, event fifa.EventResponse) {
	fmt.Printf("doHalfEnd: %s : %s\n", event.Id, event.Timestamp)
	switch event.Period {
	case fifa.FIRST:
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - End of the first half. The current score is *%s* %d - %d *%s*", event.MatchMinute, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)

		slackapi.UpdateMessage(
			os.Getenv("SLACK_OUTPUT_CHANNEL"),
			match.SlackThreadTs,
			slack.MsgOptionText(
				fmt.Sprintf(":thread: *%s* %s %d - %d %s *%s* (_%s_) - In Half Time!", match.HomeTeam, match.HomeFlag, event.HomeGoals, event.AwayGoals, match.AwayFlag, match.AwayTeam, match.Competition),
				false,
			),
		)
	case fifa.SECOND:
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - End of the second half. The current score is *%s* %d - %d *%s*", event.MatchMinute, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}
}

func doGoalAttempt(match *Match, event fifa.EventResponse) {
	fmt.Printf("doGoalAttempt: %s : %s\n", event.Id, event.Timestamp)
	var attacker Player
	var defender Player
	var ok bool

	if attacker, ok = players[event.PlayerId]; !ok {
		if defender, ok = players[event.SubPlayerId]; ok {
			slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
				slack.MsgOptionText(
					fmt.Sprintf(":goal_net: *%s* - *%s* %s attempts a goal but is thwarted by *%s* %s.", event.MatchMinute, attacker.Name, attacker.Team.Flag, defender.Name, defender.Team.Flag),
					false,
				),
				slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
					ThreadTimestamp: match.SlackThreadTs,
				}),
			)
		} else {
			slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
				slack.MsgOptionText(
					fmt.Sprintf(":goal_net: *%s* - *%s* %s attempts a goal but fails.", event.MatchMinute, attacker.Name, attacker.Team.Flag),
					false,
				),
				slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
					ThreadTimestamp: match.SlackThreadTs,
				}),
			)
		}
	} else {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":goal_net: *%s* - %s %s attempts a goal but fails.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)

	}
}

func doOffside(match *Match, event fifa.EventResponse) {
	fmt.Printf("doOffside: %s : %s\n", event.Id, event.Timestamp)

	if player, ok := players[event.PlayerId]; ok {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - *%s* %s is ruled offside.", event.MatchMinute, player.Name, player.Team.Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	} else {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - %s %s is ruled offside.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}
}

func doCorner(match *Match, event fifa.EventResponse) {
	fmt.Printf("doCorner: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":corner: *%s* - *%s* %s takes the corner kick.", event.MatchMinute, player.Name, player.Team.Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	} else {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":corner: *%s* - %s %s takes the corner kick.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}
}

func doFoul(match *Match, event fifa.EventResponse) {
	fmt.Printf("doFoul: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - *%s* %s commits a foul.", event.MatchMinute, player.Name, player.Team.Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	} else {
		slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
			slack.MsgOptionText(
				fmt.Sprintf(":referee: *%s* - %s %s commits a foul.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag),
				false,
			),
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
				ThreadTimestamp: match.SlackThreadTs,
			}),
		)
	}
}

func doMatchEnd(match *Match, event fifa.EventResponse) {
	fmt.Printf("doMatchEnd: %s : %s\n", event.Id, event.Timestamp)
	slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
		slack.MsgOptionText(
			fmt.Sprintf(":referee: *%s* - End of the match with a final score of *%s* %d - %d *%s*.", event.MatchMinute, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam),
			false,
		),
		slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
			ThreadTimestamp: match.SlackThreadTs,
		}),
	)

	match.IsOver = true
}

func doPenaltyMissed(match *Match, event fifa.EventResponse) {
	fmt.Printf("doPenaltyMissed: %s : %s\n", event.Id, event.Timestamp)
	var attacker Player
	var goalie Player
	var ok bool

	if attacker, ok = players[event.PlayerId]; !ok {
		sendError(fmt.Sprintf("Unable to lookup attacker for doGoalAttempt: %s", event.PlayerId), event)
		return
	}

	if goalie, ok = players[event.SubPlayerId]; !ok {
		sendError(fmt.Sprintf("Unable to lookup goalie for doGoalAttempt: %s", event.SubPlayerId), event)
		return
	}

	slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
		slack.MsgOptionText(
			fmt.Sprintf(":goal_net: *%s* - *%s* %s saves a penalty kick by *%s* %s.", event.MatchMinute, goalie.Name, goalie.Team.Flag, attacker.Name, attacker.Team.Flag),
			false,
		),
		slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
			ThreadTimestamp: match.SlackThreadTs,
		}),
	)
}

func doVAR(match *Match, event fifa.EventResponse) {
	fmt.Printf("doVar: %s : %s\n", event.Id, event.Timestamp)
	slackapi.PostMessage(os.Getenv("SLACK_OUTPUT_CHANNEL"),
		slack.MsgOptionText(
			fmt.Sprintf(":robot_face: *%s* - VAR Event: %s.", event.MatchMinute, event.EventDescription[0].Description),
			false,
		),
		slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{
			ThreadTimestamp: match.SlackThreadTs,
		}),
	)
}
