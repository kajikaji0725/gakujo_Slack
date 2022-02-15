package slack_bot

import (
	"fmt"
	"os"
	"time"

	"github.com/slack-go/slack"
	"github.com/szpp-dev-team/gakujo-api/model"
)

func BotNew(seiseki []*model.SeisekiRow, change []SeisekiSubject) {
	messeages := ""
	changeMessages := ""

	for _, row := range seiseki {
		messeage := fmt.Sprintf("%v\n", *row)
		messeages += messeage
	}

	messeages += "成績が更新されましたよ hoge"

	for _, row := range change {
		messeage := fmt.Sprintf("%v\n", row)
		changeMessages += messeage
	}
	changeMessages += "以上の科目の成績が追加されました。"

	api := slack.New(os.Getenv("BOT_TOKEN_TEST"))
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL_TEST"),
		slack.MsgOptionText(messeages, false),
	)
	api = slack.New(os.Getenv("BOT_TOKEN"))
	_, _, _ = api.PostMessage(
		os.Getenv("BOT_CHANNEL"),
		slack.MsgOptionText(changeMessages, false),
	)
}

func BotSame() {
	fmt.Printf("%s 草\n", time.Now().String())
}
