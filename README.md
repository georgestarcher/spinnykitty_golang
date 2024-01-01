# spinnykitty

This executable is intended to be used in conjunction with https://github.com/georgestarcher/spinnykitty_ml_api.

The ML API component is a simple Flask Python app that the GOLANG application queries for the SpinnyKitty predicition.

## Dependancies
THe ML Model API Instance. https://github.com/georgestarcher/spinnykitty_ml_api

## Configuration

You will need a folder config/ in the same location as the golang executable. 

### secrets.json

You will need to make this file to hold your Twitch username and oauth token. The format is:

```
{"twitchUser":"MyUserName","twitchOauth":"oauth:TOKENHERE"}
```

### config.json

This file holds the values such as the Twitch Channel, the URL of the ML API and the model we did requires 5 lines in the ChatFrame.

```
{"twitchChannel":"generikb","twitchBotName":"fossabot","mlApiURL":"http://127.0.0.1:5000/spinny","maxChatFrameSize":5}
```
## Usage

1. Setup the ML API from https://github.com/georgestarcher/spinnykitty_ml_api
2. Create the config folder where you want to run the golang app.
3. Setup the config files as above. You will need to reference https://github.com/georgestarcher/spinnykitty on making your Twitch Oauth token.
4. Make the golang app and copy the executable appropriate to your environment.
5. run the application and enjoy.


### Help

You're on your own. This was a learning exercise. Go through it all and good luck!  

