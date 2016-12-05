package client

import "github.com/nlopes/slack"
import "fmt"
import "os"
import "strconv"

func ListTeams() {
	slack_teams, _ := strconv.ParseInt(os.Getenv("SLACK_TEAMS"), 10, 64)
	i := int64(0)
	for {
		key := fmt.Sprintf("SLACK_TOKEN_%d", i)
		api := slack.New(os.Getenv(key))
		r := api.GetTeamInfo()
		fmt.Println("test ", r)

		i++
		if i >= slack_teams {
			break
		}
	}

}
