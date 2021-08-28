package slack_bot

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/szpp-dev-team/gakujo-api/model"
)

func BotNew(seiseki []*model.SeisekiRow, change []*model.SeisekiRow) {
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
	api := slack.New(os.Getenv("BOT_TOKEN"))
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL"),
		slack.MsgOptionText(messeages, false),
	)
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL"),
		slack.MsgOptionText(changeMessages, false),
	)
	api = slack.New("xoxb-1491378823348-2286618846368-uIxUrd1gGmuRC2QBPe1euDjr")
	messeages = ""
	for _, row := range change {
		messeage := fmt.Sprintf("%v\n", row.SubjectName)
		messeages += messeage
	}
	messeages += "この科目の成績が追加されました。"
	_, _, _ = api.PostMessage(
		"random",
		slack.MsgOptionText(messeages, false),
	)
}

func BotSame() {
	api := slack.New(os.Getenv("BOT_TOKEN"))
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL"),
		slack.MsgOptionText("成績に変更はありません", false),
	)
}
