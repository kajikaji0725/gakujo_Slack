package slack_bot

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/szpp-dev-team/gakujo-api/model"
)

func BotNew(seiseki []*model.SeisekiRow) {
	fmt.Println(os.Getenv("BOT"))
	messeages := ""
	for _, row := range seiseki {
		messeage := fmt.Sprintf("%v\n", *row)
		messeages += messeage
	}
	messeages += "成績が更新されましたよ　のみ"
	api := slack.New(os.Getenv("BOT"))
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL"),
		slack.MsgOptionText(messeages, false),
	)
}

func BotSame() {
	fmt.Println(os.Getenv("BOT"))
	api := slack.New(os.Getenv("BOT"))
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL"),
		slack.MsgOptionText("成績に変更はありません　のみ", false),
	)
}
