package main

import (
	fifa "github.com/ImDevinC/go-fifa"
	"github.com/sirupsen/logrus"
)

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
	*/
	fifa.OwnGoal: doOwnGoal,
	/*
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

func doGoalScore(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doGoalScore: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		SendSlackMessage(match, ":soccer: *%s* - GOOOOOOOOOOOOOAAAAAAAAAAAAAAAAAAALLLLLLLLLLLLLLLL!!!! *%s* %s scores a goal. The score is now *%s* %d - %d *%s*.", event.MatchMinute, player.Name, player.Team.Flag, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam)
	} else {
		SendSlackMessage(match, ":soccer: *%s* - GOOOOOOOOOOOOOAAAAAAAAAAAAAAAAAAALLLLLLLLLLLLLLLL!!!! %s %s scores a goal. The score is now *%s* %d - %d *%s*.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam)
	}

	go FindInstantReplay(match, event)
}

func doOwnGoal(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doOwnGoal: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		SendSlackMessage(match, ":soccer: *%s* - OWN GOAL! *%s* %s scores in their own goal! The score is now *%s* %d - %d *%s*.", event.MatchMinute, player.Name, player.Team.Flag, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam)
	} else {
		SendSlackMessage(match, ":soccer: *%s* - OWN GOAL! %s %s scores in their own goal!. The score is now *%s* %d - %d *%s*.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam)
	}

	go FindInstantReplay(match, event)
}

func doYellowCard(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doYellowCard: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		SendSlackMessage(match, ":referee-with-yellow-card: *%s* - *%s* %s is booked with a yellow card.", event.MatchMinute, player.Name, player.Team.Flag)
	} else {
		SendSlackMessage(match, ":referee-with-yellow-card: *%s* - %s %s receives a yellow card.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag)
	}
}

func doRedCard(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doRedCard: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		SendSlackMessage(match, ":referee-with-red-card: *%s* - *%s* %s is sent off with a red card.", event.MatchMinute, player.Name, player.Team.Flag)
	} else {
		SendSlackMessage(match, ":referee-with-red-card: *%s* - *%s* %s receives a red card.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag)
	}
}

func doSubstitution(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doSubstitution: %s : %s\n", event.Id, event.Timestamp)
	var playerIn Player
	var playerOut Player
	var ok bool

	if playerIn, ok = players[event.PlayerId]; ok {
		if playerOut, ok = players[event.SubPlayerId]; ok {
			SendSlackMessage(match, ":arrows_counterclockwise: *%s* - *%s* %s comes in for *%s* %s.", event.MatchMinute, playerIn.Name, playerIn.Team.Flag, playerOut.Name, playerOut.Team.Flag)
		} else {
			SendSlackMessage(match, ":arrows_counterclockwise: *%s* - *%s* %s is substituted.", event.MatchMinute, playerIn.Name, playerIn.Team.Flag)
		}
	} else {
		SendSlackMessage(match, ":arrows_counterclockwise: *%s* - *%s* %s substitutes a player.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag)
	}
}

func doAwardPenalty(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doAwardPenalty: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		SendSlackMessage(match, ":referee: *%s* - *%s* %s receives a penalty.", event.MatchMinute, player.Name, player.Team.Flag)
	} else {
		SendSlackMessage(match, ":referee: *%s* - *%s* %s receives a penalty.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag)
	}
}

func doMatchStart(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doMatchStart: %s : %s\n", event.Id, event.Timestamp)
	switch event.Period {
	case fifa.FIRST:
		SendSlackMessage(match, ":referee: *%s* - The match has started!", event.MatchMinute)
	case fifa.SECOND:
		SendSlackMessage(match, ":referee: *%s* - Start of the second half.", event.MatchMinute)
	}
}

func doHalfEnd(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doHalfEnd: %s : %s\n", event.Id, event.Timestamp)
	switch event.Period {
	case fifa.FIRST:
		SendSlackMessage(match, ":referee: *%s* - End of the first half. The current score is *%s* %d - %d *%s*", event.MatchMinute, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam)
	case fifa.SECOND:
		SendSlackMessage(match, ":referee: *%s* - End of the second half. The current score is *%s* %d - %d *%s*", event.MatchMinute, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam)
	}
}

func doGoalAttempt(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doGoalAttempt: %s : %s\n", event.Id, event.Timestamp)
	var attacker Player
	var defender Player
	var ok bool

	if attacker, ok = players[event.PlayerId]; !ok {
		if defender, ok = players[event.SubPlayerId]; ok {
			SendSlackMessage(match, ":goal_net: *%s* - *%s* %s attempts a goal but is thwarted by *%s* %s.", event.MatchMinute, attacker.Name, attacker.Team.Flag, defender.Name, defender.Team.Flag)
		} else {
			SendSlackMessage(match, ":goal_net: *%s* - *%s* %s attempts a goal but fails.", event.MatchMinute, attacker.Name, attacker.Team.Flag)
		}
	} else {
		SendSlackMessage(match, ":goal_net: *%s* - %s %s attempts a goal but fails.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag)
	}
}

func doOffside(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doOffside: %s : %s\n", event.Id, event.Timestamp)

	if player, ok := players[event.PlayerId]; ok {
		SendSlackMessage(match, ":referee: *%s* - *%s* %s is ruled offside.", event.MatchMinute, player.Name, player.Team.Flag)
	} else {
		SendSlackMessage(match, ":referee: *%s* - %s %s is ruled offside.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag)
	}
}

func doCorner(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doCorner: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		SendSlackMessage(match, ":corner: *%s* - *%s* %s takes the corner kick.", event.MatchMinute, player.Name, player.Team.Flag)
	} else {
		SendSlackMessage(match, ":corner: *%s* - %s %s takes the corner kick.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag)
	}
}

func doFoul(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doFoul: %s : %s\n", event.Id, event.Timestamp)
	if player, ok := players[event.PlayerId]; ok {
		SendSlackMessage(match, ":referee: *%s* - *%s* %s commits a foul.", event.MatchMinute, player.Name, player.Team.Flag)
	} else {
		SendSlackMessage(match, ":referee: *%s* - %s %s commits a foul.", event.MatchMinute, teams[event.TeamId].Name, teams[event.TeamId].Flag)
	}
}

func doMatchEnd(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doMatchEnd: %s : %s\n", event.Id, event.Timestamp)
	SendSlackMessage(match, ":referee: *%s* - End of the match with a final score of *%s* %d - %d *%s*.", event.MatchMinute, match.HomeTeam, event.HomeGoals, event.AwayGoals, match.AwayTeam)
	match.IsOver = true
}

func doPenaltyMissed(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doPenaltyMissed: %s : %s\n", event.Id, event.Timestamp)
	attacker := players[event.PlayerId]
	goalie := players[event.SubPlayerId]

	SendSlackMessage(match, ":goal_net: *%s* - *%s* %s saves a penalty kick by *%s* %s.", event.MatchMinute, goalie.Name, goalie.Team.Flag, attacker.Name, attacker.Team.Flag)
}

func doVAR(match *Match, event fifa.EventResponse) {
	logrus.Debugf("doVar: %s : %s\n", event.Id, event.Timestamp)
	SendSlackMessage(match, ":robot_face: *%s* - VAR Event: %s.", event.MatchMinute, event.EventDescription[0].Description)
}
