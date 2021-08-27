package slack_bot

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/szpp-dev-team/gakujo-api/model"
)

func BotNew(seiseki []*model.SeisekiRow, change []*model.SeisekiRow) {
	fmt.Println(os.Getenv("BOT"))
	messeages := ""
	changeMessages := ""
	for _, row := range seiseki {
		messeage := fmt.Sprintf("%v\n", *row)
		messeages += messeage
	}
	messeages += "成績が更新されましたよ"

	for _, row := range change {
		changemesseage := fmt.Sprintf("%v\n", *row)
		changeMessages += changemesseage
	}
	changeMessages += "これが追加されました"
	api := slack.New(os.Getenv("BOT"))
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL"),
		slack.MsgOptionText(messeages, false),
	)
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL"),
		slack.MsgOptionText(changeMessages, false),
	)
}

func BotSame() {
	fmt.Println(os.Getenv("BOT"))
	api := slack.New(os.Getenv("BOT"))
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL"),
		slack.MsgOptionText("成績に変更はありません", false),
	)
}
