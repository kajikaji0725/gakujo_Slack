package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kajikaji0725/gakujo_Slack/slack_bot"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("please set .env on ./..", err)
	}
	fmt.Println(os.Getenv("J_USERNAME"))
	fmt.Println(os.Getenv("J_PASSWORD"))
	fmt.Println(os.Getenv("BOT_TOKEN"))
	fmt.Println(os.Getenv("BOT_CHANNEL"))
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
