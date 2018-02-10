package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/nlopes/slack"
)

type payloadIn struct {
	Action string `json:"action"`
	Sender sender `json:"sender"`
	Repo   repo   `json:"repository"`
}

type sender struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

type repo struct {
	Name            string `json:"name"`
	StarGazersCount int    `json:"stargazers_count"`
}

func main() {
	p := new(payloadIn)
	json.NewDecoder(os.Stdin).Decode(p)

	fmt.Printf("%+v \n", p)

	fmt.Println(p.Sender.Login)

	api := slack.New(os.Getenv("FIN_SLACK_KEY"))
	params := slack.PostMessageParameters{
		AsUser: true,
	}

	var b bytes.Buffer
	b.WriteString(":star: User *" + p.Sender.Login + "* starred " + p.Repo.Name + "\n")
	b.WriteString("        Total stars now *" + strconv.Itoa(p.Repo.StarGazersCount) + "*")

	_, _, err := api.PostMessage("demostream", b.String(), params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
