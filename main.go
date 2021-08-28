package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kajikaji0725/gakujo_Slack/slack_bot"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
)

var Api []string

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("please set .env on ./..", err)
	}
	Api = flag.Args()
	c := gakujo.NewClient()
	if err := c.Login(Api[0], Api[1]); err != nil {
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
