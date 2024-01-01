package spinnykitty

import (
	"testing"
)

// Test Setting Frame Size to int 5
func TestSetMaximum(t *testing.T) {

	TestChatFrame := new(ChatFrame)
	TestChatFrame.SetMaximum(5)

	want := 5
	got := TestChatFrame.MaxSize

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test Setting Frame ML API URL
func TestSetMLApiUrl(t *testing.T) {

	TestChatFrame := new(ChatFrame)

	TestChatFrame.SetMLApiUrl("http://127.0.0.1:5000/spinny")

	want := "http://127.0.0.1:5000/spinny"
	got := TestChatFrame.MLApiUrl

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test Setting Frame Bot Name
func TestSetBotName(t *testing.T) {

	TestChatFrame := new(ChatFrame)

	TestChatFrame.SetBotName("fossabot")

	want := "fossabot"
	got := TestChatFrame.BotName

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test Empty Feature Frame Bot Name
func TestFeatures(t *testing.T) {

	TestChatFrame := new(ChatFrame)

	frameFeatures, _ := TestChatFrame.Features()

	want := "{\"time\":-9223372036.854776,\"distinctUserCount\":0,\"meanTokenCount\":0,\"meanSpinnyCount\":0}"
	got := string(frameFeatures)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsReady(t *testing.T) {

	TestChatFrame := new(ChatFrame)
	TestChatFrame.SetMaximum(5)

	want := false
	got := TestChatFrame.IsReady()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsSpinny(t *testing.T) {

	TestChatFrame := new(ChatFrame)
	TestChatFrame.SetMaximum(5)

	want := false
	got, _ := TestChatFrame.IsSpinny()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
