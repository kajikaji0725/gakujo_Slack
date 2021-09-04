package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kajikaji0725/gakujo_Slack/slack_bot"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
)

func main() {
	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal("please set .env on ./..", err)
	// }
	c := gakujo.NewClient()
	if err := c.Login(os.Getenv("J_USERNAME"), os.Getenv("J_PASSWORD")); err != nil {
		fmt.Println("hoeg")
		log.Fatal(err)
	}

	kc, err := c.NewKyoumuClient()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := kc.SeisekiRows()
	if err != nil {
		log.Fatal(err)
	}
	er := slack_bot.UpdateSeisekiFile(rows)
	if er != nil {
		log.Fatal(er)
	}
}
