package spinnykitty

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"utilities"

	"github.com/gempir/go-twitch-irc/v4"
)

// Chat Emote Constants
const EmojiSpinnyKitty string = "hermitSpinnyKitty"
const EmojiDance string = "hermitDance"
const EmojiYo string = "hermitYoYoYo"
const EmojiWarsh string = "hermitWarsh"
const EmojiEttie string = "hermitEttie"
const EmojiMilk string = "hermitMilk"
const EmojiHeart string = "hermitBitHeart"
const EmojiCrown string = "hermitCrown"
const EmojiCrabRave string = "hermitCrabRave"
const EmojiLoveMilk string = "hermitMilkhermitBitHeart"
const EmojiLissen string = "hermitLISSEN"
const EmojiJerkbag string = "hermitJerkbag"
const EmojiRage string = "hermitBitRage"
const EmojiDerp string = "hermitDerp"
const EmojiMariah string = "hermitMariahCarey"
const EmojiClap string = "hermitClap"
const EmojiFreshMeat string = "hermitMeat"

// Define sliding Chat Message Frame
type ChatFrame struct {
	MaxSize      int
	MessageFrame []twitch.PrivateMessage
	MLApiUrl     string
	BotName      string
}

type SpinnyPrediction struct {
	Prediction int `json:"prediction"`
}

// ML Features of a Chat Frame
type ChatFrameFeatures struct {
	TimeRangeSeconds  float64 `json:"time"`
	DistinctUserCount int     `json:"distinctUserCount"`
	MeanTokenCount    float64 `json:"meanTokenCount"`
	MeanSpinnyCount   float64 `json:"meanSpinnyCount"`
}

// Set maximum chat frame size
func (frame *ChatFrame) SetMaximum(size int) {
	frame.MaxSize = size
}

// Set maximum ML API URL
func (frame *ChatFrame) SetBotName(botName string) {
	frame.BotName = botName
}

// Set maximum ML API URL
func (frame *ChatFrame) SetMLApiUrl(url string) {
	frame.MLApiUrl = url
}

// Is the Chat Frame Ready for Use
func (frame ChatFrame) IsReady() bool {
	if len(frame.MessageFrame) != frame.MaxSize {
		return false
	} else {
		return true
	}
}

// Add new chat Message to the frame buffer
func (frame *ChatFrame) AddMessage(m twitch.PrivateMessage) {

	// Ignore messages from channel bot
	if strings.EqualFold(m.User.Name, frame.BotName) {
		return
	}

	// append to full frame contents if frame not max size
	if !frame.IsReady() {
		frame.MessageFrame = append(frame.MessageFrame, m)
	} else {
		// append dropping leading message if frame at max size
		frame.MessageFrame = append(frame.MessageFrame[1:], m)
	}
}

// Calc Mean number of specified tokens in the Chat Frame
func (frame ChatFrame) meanTokenCount(token string) float64 {

	returnMean := 0.0

	if !frame.IsReady() {
		return returnMean
	}

	var tokenCounts []int

	for _, msg := range frame.MessageFrame {
		count := len(strings.SplitAfter(msg.Message, token))
		if count == 1 {
			count = 0
		}
		tokenCounts = append(tokenCounts, count)
	}
	return utilities.CalcMean(tokenCounts)
}

// Calculate the Distinct number of users in the Chat Frame
func (frame ChatFrame) distinctUserCount() int {

	if !frame.IsReady() {
		return 0
	}

	counter := make(map[string]int)

	for _, msg := range frame.MessageFrame {
		counter[msg.User.Name]++
	}
	return len(counter)
}

// Calculate the time range across the Chat Frame in seconds
func (frame ChatFrame) timeRangeSeconds() float64 {

	if !frame.IsReady() {
		return 0
	}

	var min_time time.Time
	var max_time time.Time

	min_time = time.Now()

	for _, msg := range frame.MessageFrame {
		if msg.Time.After(max_time) {
			max_time = msg.Time
		}
		if msg.Time.Before(min_time) {
			min_time = msg.Time
		}
	}
	return max_time.Sub(min_time).Seconds()
}

// Output the ML Features of the Chat Frame
func (frame ChatFrame) Features() ([]byte, error) {
	features := ChatFrameFeatures{frame.timeRangeSeconds(), frame.distinctUserCount(), frame.meanTokenCount(" "), frame.meanTokenCount(EmojiSpinnyKitty)}
	features_json, err := json.Marshal(features)
	if err != nil {
		return nil, err
	}
	return features_json, nil
}

func (frame ChatFrame) IsSpinny() (bool, error) {

	if !frame.IsReady() {
		return false, nil
	}

	feature_body, _ := frame.Features()

	r, err := http.NewRequest("POST", frame.MLApiUrl, bytes.NewBuffer(feature_body))
	if err != nil {
		log.Println(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	post := &SpinnyPrediction{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		log.Println(err)
	}

	if res.StatusCode != http.StatusCreated {
		log.Println(res.Status)
	}
	if post.Prediction == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
