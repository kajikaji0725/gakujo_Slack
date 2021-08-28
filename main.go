package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kajikaji0725/gakujo_Slack/slack_bot"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
)

func main() {
	// flag.Parse()
	// fmt.Println(flag.Args())
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("please set .env on ./..", err)
	}
	c := gakujo.NewClient()
	if err := c.Login(os.Getenv("J_USERNAME"), os.Getenv("J_PASSWORD")); err != nil {
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
	slack_bot.UpdateSeisekiFile(rows)
}
