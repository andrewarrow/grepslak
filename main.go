package main

import "github.com/andrewarrow/localsdata/client"
import "os"
import "fmt"

func main() {
	if len(os.Args) == 1 {
		client.ListTeams()
		return
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "clean" {
			client.Clean()
			return
		} else if os.Args[1] == "search" {
			client.Search()
			return
		} else if os.Args[1] == "-v" {
			fmt.Println("0.8.3")
			return
		} else if os.Args[1] == "stats" {
			client.Stats("hdt")
			return
		}
		client.ListRooms(os.Args[1])
		return
	}
	if len(os.Args) == 3 {
		client.SaveHistory(os.Args[1], os.Args[2])
		return
	}
	if len(os.Args) == 4 {
		client.Say(os.Args[1], os.Args[2], os.Args[3])
		return
	}
}
