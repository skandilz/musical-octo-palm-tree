package main

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"math/rand"
	"os"
	"twitterdm/api"
)

func main() {

	var twitter *api.Twitter

	tokenPrompt := promptui.Prompt{
		Label: "Auth Token",
	}

	userPrompt := promptui.Prompt{
		Label: "Username",
	}

	messagesPrompt := promptui.Prompt{
		Label:   "Messages Path",
		Default: "messages",
		Validate: func(s string) error {
			entries, _ := os.ReadDir(s)
			if len(entries) < 1 {
				return errors.New("no files are found")
			}
			return nil
		},
	}

	for twitter == nil || twitter.Id == "" {
		res, _ := tokenPrompt.Run()
		twitter = api.NewTwitter(res)
	}

	fmt.Printf("✔ Logged in as {%s} (id: %s)", twitter.Name, twitter.Id)

	targetId := ""
	for targetId == "" {
		res, _ := userPrompt.Run()
		targetId = twitter.GetIdFromName(res)
	}

	path, _ := messagesPrompt.Run()

	entries, _ := os.ReadDir(path)
	followers := twitter.GetFollowersFromId(targetId)

	for i := 0; i < len(followers); i++ {
		follower := followers[i]
		messageIndex := rand.Intn(len(entries))
		message, _ := os.ReadFile(path + "/" + entries[messageIndex].Name())

		status := twitter.SendMessage(follower.Id, string(message))
		if status {
			fmt.Printf("Message sent successfully to @%s (id: %s) (message: %s)\n", follower.Name, follower.Id, entries[messageIndex].Name())
		}
	}
}
