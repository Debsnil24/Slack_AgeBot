package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func PrintCommandEvt(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func Calc(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	year := request.Param("year")
	yob, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("Error: Unable to convert year")
	}
	age := time.Now().Year() - yob
	r := fmt.Sprintf("Age is %d", age)
	response.Reply(r)
}