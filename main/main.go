package main

import (
	"context"
	"log"
	"os"

	"github.com/Debsnil24/Slack_AgeBot.git/controller"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	err := godotenv.Load("/Users/debsnilsamudra/Documents/Program/Slack_AgeBot/.env")
	if err != nil {
		log.Fatalf("Error: Unable to load .env file %v", err)
	}

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go controller.PrintCommandEvt(bot.CommandEvents())

	bot.Command("My YOB is <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		Examples: []string{"My YOB is 2002"},
		Handler: controller.Calc,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}