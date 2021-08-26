package slack_bot

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/szpp-dev-team/gakujo-api/model"
)

func Bot_new(seiseki []*model.SeisekiRow) {
	messeages := ""
	for _, row := range seiseki {
		messeage := fmt.Sprintf("%v\n", *row)
		messeages += messeage
	}
	messeages += "成績が更新されましたよ　のみ"
	api := slack.New("xoxb-2363568815571-2368123169266-ta18hwKdJxfxaXOlcY2zA0fr")
	_, _, _ = api.PostMessage(
		"botてすと",
		slack.MsgOptionText(messeages, false),
	)
}

func Bot_same() {
	api := slack.New("xoxb-2363568815571-2368123169266-ta18hwKdJxfxaXOlcY2zA0fr")
	_, _, _ = api.PostMessage(
		"botてすと",
		slack.MsgOptionText("成績に変更はありません　のみ", false),
	)
}
