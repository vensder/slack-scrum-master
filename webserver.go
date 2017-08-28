package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	separator = `------------------------------`
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go!")
}

func handlerSlack(w http.ResponseWriter, r *http.Request) {

	//Read the Request Parameter "command"
	token := r.FormValue("token") // TODO: after turning on ssl on the end point
	// teamID := r.FormValue("team_id")
	// teamDomain := r.FormValue("team_domain")
	// channelID := r.FormValue("channel_id")
	// channelName := r.FormValue("channel_name")
	// userID := r.FormValue("user_id")
	// userName := r.FormValue("user_name")
	command := r.FormValue("command")
	// text := r.FormValue("text")
	// responseURL := r.FormValue("response_url")

	// TODO: move participants to config file
	members := []string{
		"vensder",
		"dev1",
		"dev2",
		"dev3",
		"dev4",
		"devops1",
		"qa1",
		"admin1",
	}

	//Ideally do other checks for tokens/username/channel/etc - after SSL

	if (token == "ToKeNtOkEnTOkEN") && (command == "/scrum") {
		rand.Seed(time.Now().UnixNano())
		for i := range members {
			j := rand.Intn(i + 1)
			members[i], members[j] = members[j], members[i]
		}
		outputString := strings.Join(members, "\n")
		fmt.Fprintf(w, "%v\n%v", outputString, separator)

	} else {
		fmt.Fprint(w, "I do not understand your command.")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/slack", handlerSlack)
	http.ListenAndServe(":8080", nil)
}
