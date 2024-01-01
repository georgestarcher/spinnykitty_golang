package spinnykitty

import (
	"log"
	"testing"
)

// Test Setting Frame Size to int 5
func TestChatFrameSetValidMax(t *testing.T) {

	TestChatFrame := new(ChatFrame)
	TestChatFrame.SetMaximum(5)

	want := 5
	got := TestChatFrame.MaxSize

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	log.Println(got)
}

// Test Setting Frame ML API URL
func TestChatFrameSetValidAPIUrl(t *testing.T) {

	TestChatFrame := new(ChatFrame)

	TestChatFrame.SetMLApiUrl("http://127.0.0.1:5000/spinny")

	want := "http://127.0.0.1:5000/spinny"
	got := TestChatFrame.MLApiUrl

	log.Println(TestChatFrame.MLApiUrl)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	log.Println(got)
}

// Test Setting Frame Bot Name
func TestChatFrameSetValidBotName(t *testing.T) {

	TestChatFrame := new(ChatFrame)

	TestChatFrame.SetBotName("fossabot")
	log.Println(TestChatFrame)

	want := "fossabot"
	got := TestChatFrame.BotName

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	log.Println(got)
}

// Test Empty Feature Frame Bot Name
func TestChatFrameEmptyFeatures(t *testing.T) {

	TestChatFrame := new(ChatFrame)

	frameFeatures, _ := TestChatFrame.Features()

	want := "{\"time\":-9223372036.854776,\"distinctUserCount\":0,\"meanTokenCount\":0,\"meanSpinnyCount\":0}"
	got := string(frameFeatures)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	log.Println(got)
}

func TestChatFrameEmptyReady(t *testing.T) {

	TestChatFrame := new(ChatFrame)
	TestChatFrame.SetMaximum(5)

	want := false
	got := TestChatFrame.IsReady()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	log.Println(got)
}

func TestChatFrameEmptyKitty(t *testing.T) {

	TestChatFrame := new(ChatFrame)
	TestChatFrame.SetMaximum(5)

	want := false
	got, _ := TestChatFrame.IsSpinny()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	log.Println(got)
}
