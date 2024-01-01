package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"spinnykitty"

	"github.com/gempir/go-twitch-irc/v4"
)

type Secrets struct {
	TwitchUser  string `json:"twitchUser"`
	TwitchOauth string `json:"twitchOauth"`
}

type Config struct {
	TwitchChatChannel string `json:"twitchChannel"`
	TwitchBotName     string `json:"twitchBotName"`
	MLApiUrl          string `json:"mlApiURL"`
	MaxChatFrameSize  int    `json:"maxChatFrameSize"`
}

func main() {

	secretsFile, err := ioutil.ReadFile("./config/secrets.json")
	if err != nil {
		log.Fatal("Error when opening twitch secrets file: ", err)
	}

	var configSecrets Secrets
	err = json.Unmarshal(secretsFile, &configSecrets)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	configFile, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal("Error when opening config file: ", err)
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
	client := twitch.NewClient(configSecrets.TwitchUser, configSecrets.TwitchOauth)

	// Setup Chat Frame
	var chatFrame spinnykitty.ChatFrame
	chatFrame.SetMaximum(config.MaxChatFrameSize)
	chatFrame.SetMLApiUrl(config.MLApiUrl)
	chatFrame.SetBotName(config.TwitchBotName)

	fmt.Println(chatFrame)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.User.Name + ":" + message.Message)
		// Add message to Chat Frame
		chatFrame.AddMessage(message)
		//chatFrameFeaturesJSON, _ := chatFrame.Features()
		chatIsSpinny, _ := chatFrame.IsSpinny()
		fmt.Printf("isSpinnyKitty:%v\n", chatIsSpinny)
	})

	client.Join(config.TwitchChatChannel)

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
