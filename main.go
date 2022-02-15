package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kajikaji0725/gakujo_Slack/slack_bot"
	"github.com/robfig/cron/v3"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("please set .env on ./..", err)
	}
}

func main() {
	cr := cron.New()
	cr.AddFunc("@hourly", func() {
		if time.Now().Hour() != 2 || time.Now().Hour() != 3 || time.Now().Hour() != 4 || time.Now().Hour() != 5 {
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
			er := slack_bot.UpdateSeisekiFile(rows)
			if er != nil {
				log.Fatal(er)
			}
		}
	})

	cr.Start()

	for {
		time.Sleep(time.Hour * 24)
	}
}
