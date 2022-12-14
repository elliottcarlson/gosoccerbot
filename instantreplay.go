package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/avast/retry-go/v4"
	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	snoo "github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func FindInstantReplay(match *Match, event fifa.EventResponse) {
	retry.Do(
		func() error {
			var reddit, _ = snoo.NewClient(snoo.Credentials{
				ID:       os.Getenv("REDDIT_CLIENT_ID"),
				Secret:   os.Getenv("REDDIT_CLIENT_SECRET"),
				Username: os.Getenv("REDDIT_CLIENT_USERNAME"),
				Password: os.Getenv("REDDIT_CLIENT_PASSWORD"),
			})

			var query string
			if teams[event.TeamId].Name == match.HomeTeam {
				query = fmt.Sprintf("\"%s [%d] - %d %s\"", match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam)
			} else {
				query = fmt.Sprintf("\"%s %d - [%d] %s\"", match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam)
			}

			logrus.Debugf("Searching for replay video: %s\n", query)

			posts, _, err := reddit.Subreddit.SearchPosts(ctx, query, "soccer", &snoo.ListPostSearchOptions{
				ListPostOptions: snoo.ListPostOptions{
					Time: "hour",
				},
				Sort: "relevance",
			})

			if err != nil {
				return err
			}

			for _, post := range posts {
				if strings.HasPrefix(post.URL, "https://dubz.co/v") ||
					strings.HasPrefix(post.URL, "https://streamin.me/v") {
					c := colly.NewCollector(
						colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"),
					)

					c.OnHTML("video", func(e *colly.HTMLElement) {
						d := colly.NewCollector(
							colly.MaxBodySize(100 * 1024 * 1024),
						)
						d.OnResponse(func(r *colly.Response) {
							err := r.Save("media/" + r.FileName())
							if err != nil {
								logrus.Errorf("Unable to save %s: %v", r.FileName(), err)
							}

							var title string
							if player, ok := players[event.PlayerId]; ok {
								title = fmt.Sprintf("Replay of the goal at %s by %s for %s", event.MatchMinute, player.Name, player.Team.Name)
							} else {
								title = fmt.Sprintf("Replay of the goal at %s", event.MatchMinute)
							}

							logrus.Debugf("Uploading %s to Slack\n", title)

							_, err = slackapi.UploadFile(slack.FileUploadParameters{
								File:  "media/" + r.FileName(),
								Title: title,
								Channels: []string{
									os.Getenv("SLACK_OUTPUT_CHANNEL"),
								},
								ThreadTimestamp: match.SlackThreadTs,
							})

							if err != nil {
								logrus.Debugf("Error uploading %s to Slack: %v", r.FileName(), err)
							}
							os.Remove("media/" + r.FileName())
						})
						d.OnRequest(func(r *colly.Request) {
							logrus.Debugf("Downloading %s", r.URL)
						})
						d.Visit(e.Attr("src"))
					})

					c.OnRequest(func(r *colly.Request) {
						logrus.Debugf("Potential match found; Visiting %s", r.URL)
					})

					c.Visit(post.URL)
					return nil
				} else {
					logrus.Debugf("Unsupported download URL found: %s", post.URL)
				}
			}

			return errors.New("none found")
		},
		retry.OnRetry(func(n uint, err error) {
			logrus.Debugf("Search attempt #%d: %s\n", n, err)
		}),
		retry.Attempts(uint((10*time.Minute)/(15*time.Second))),
		retry.Delay(15*time.Second),
		retry.DelayType(retry.FixedDelay),
	)

}
