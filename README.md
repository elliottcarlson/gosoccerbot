# Soccer Bot

Go based Slack bot to show live scores from the World Cup

## Features:

* Threads per match
* Instant replays of goals

## Setup

### Slack

Visit https://api.slack.com/apps/ and "Create New App". Select your target workspace and then
choose "From an app manifest". Select YAML as the manifest type, and paste the contents of
`slack_manifest.yaml` in to the editor. Review the settings and "Create" the bot.

On the Basic Information page, scroll down to "App-Level Tokens", and Generate an app-level
token. This token needs the scope `connections:write`. Once created, copy the xapp- token
and paste it in to your .env file (modeled after the .env.example file) as the `SLACK_APP_TOKEN`

Next visit the "OAuth & Permissions" section in the left hand menu. Click "Install to Workspace".
Once installed, you will be redirected to the "OAuth & Permissions" page, and now an OAuth token
will be available for the workspace - copy this xoxb- token and paste it in your .env file as
the `SLACK_BOT_TOKEN`.

Next you want to configure the .env files values for where to send match results to; In Slack
right click the channel from your Channels list that you want to send match data to, select 
"View channel details", and in the popup scroll down and you can find the Channel ID; paste this
value in the .env file under `SLACK_OUTPUT_CHANNEL`. For the `SLACK_ADMIN_USER_ID`, you will
want to visit your Slack profile (reached by right clicking your name), and under the three dot
menu, you can Copy member ID.

Make sure you invite the bot to the intended channel for it to be able to post there!

Finally, you can upload the images in the "images" directory as emoji using the default names of
the files to your Slack instance to ensure the bot has all the appropriate emojis in place.

### Reddit

Optional: Create a new reddit account since you need to store your `REDDIT_CLIENT_USERNAME` and
`REDDIT_CLIENT_PASSWORD` directly in the file.

Visit https://www.reddit.com/prefs/apps. At the bottom of the page "Create application". Give it
an appropriate name, and make sure to select "script" as the type. Neither the about or redirect
URLs need to be valid URLs. Once created you will have your `REDDIT_CLIENT_ID` (which is listed
immediately under your app name and "persoinal use script" - it does not denote it in any clear
way... damn you reddit.), and your secret that goes in `REDDIT_CLIENT_SECRET` in the .env file.

### Running the bot

With go 1.17+ installed, simply run
```
go run .
```

Any errors should be available on the terminal.