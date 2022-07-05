package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/shomali11/slacker"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sithu-aung/go_bookstore/pkg/routes"
)

func printCommetEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Commands Event : ")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3752220142306-3752175756099-f0H6NIKXqTl9JkljC77bGJ0Y")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03N6JTSFFE-3752223032418-f6af86c3cd2b60ba552966ff83e04f09f2d426bfdb85c1f1f386f4075d16a028")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommetEvents(bot.CommandEvents())

	bot.Command("My Wife was born in <year>", &slacker.CommandDefinition{
		Description: "This command will tell you the year of your wife",
		Example: "My Wife was born in 1994",
		Handler: func(botCtx slacker.BotContext,request slacker.Request,response slacker.ResponseWriter){
			year := request.Param("year")
			yob,err := strconv.Atoi(year)
			if err != nil {
				println("Error")
			}	
			age := 2022 - yob
			r:= fmt.Sprintf("Your wife is %d years old",age)
			response.Reply(r)
	},
  })


	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	error := bot.Listen(ctx)
	if error != nil {
		log.Fatal(error)
	}
	


	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}