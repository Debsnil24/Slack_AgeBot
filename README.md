# Slack_AgeBot
Create a Simple Slack Bot for calculating age using Go

Update the slacker.go file at the line of error with the following code

```go
    botparams := slack.GetBotInfoParameters{
		Bot: messageEvent.BotID,
	}
	bot, err := s.apiClient.GetBotInfo(botparams)
```