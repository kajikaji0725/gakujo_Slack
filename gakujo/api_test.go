package gakujo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var (
	begin    time.Time
	username string
	password string
	c        *Client
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("please set .env on ./..", err)
	}

	username = os.Getenv("J_USERNAME")
	password = os.Getenv("J_PASSWORD")
	begin = time.Now()
	c = NewClient()
	if err := c.Login(username, password); err != nil {
		log.Fatal("failed to login")
	}
	log.Println("[Info]Login succeeded(took:", time.Since(begin), "ms)")
}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

func TestLogin(t *testing.T) {
	inc := NewClient()
	if err := inc.Login(username, password); err != nil {
		t.Fatal(err)
	}
	t.Log("[Info]Login succeeded(took:", time.Since(begin), "ms)")
}

func TestHome(t *testing.T) {
	homeInfo, err := c.Home()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(homeInfo)
}

func TestNoticeDetail(t *testing.T) {
	c := NewClient()
	if err := c.Login(username, password); err != nil {
		t.Fatal(err)
	}
	t.Log("[Info]Login succeeded(took:", time.Since(begin), "ms)")
	index := "1"
	noticeDetail, err := c.NoticeDetail(index)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 11; i++ {
		switch {
		case i == 0:
			fmt.Print("連絡種別  ")
			fmt.Println(noticeDetail.ContactType)
		case i == 1:
			fmt.Print("タイトル  ")
			fmt.Println(noticeDetail.Title)
		case i == 2:
			fmt.Print("連絡内容  ")
			fmt.Println(noticeDetail.Detail)
		case i == 3:
			fmt.Print("ファイル  ")
			fmt.Println(noticeDetail.File)
		case i == 4:
			fmt.Print("ファイルリンク公開  ")
			fmt.Println(noticeDetail.FilelinkPublication)
		case i == 5:
			fmt.Print("参照URL  ")
			fmt.Println(noticeDetail.ReferenceURL)
		case i == 6:
			fmt.Print("重要度  ")
			fmt.Println(noticeDetail.Important)
		case i == 7:
			fmt.Print("日時  ")
			fmt.Println(noticeDetail.Date)
		case i == 8:
			fmt.Print("WEB返信要求  ")
			fmt.Println(noticeDetail.WebReturnRequest)
		}

	}
}

func TestClassNoticeRow(t *testing.T) {
	c := NewClient()
	if err := c.Login(username, password); err != nil {
		t.Fatal(err)
	}
	t.Log("[Info]Login succeeded(took:", time.Since(begin), "ms)")
	classNoticeRow, err := c.ClassNotice()
	if err != nil {
		t.Fatal(err)
	}
	for _, noticerow := range classNoticeRow {
		fmt.Print("授業科目  ")
		fmt.Println(noticerow.CourseName)
		fmt.Print("担当教員名  ")
		fmt.Println(noticerow.TeachersName)
		fmt.Print("タイトル  ")
		fmt.Println(noticerow.Title)
		fmt.Print("連絡種別  ")
		fmt.Println(noticerow.Type)
		fmt.Print("対象日  ")
		fmt.Println(noticerow.TargetDate)
		fmt.Print("連絡日時  ")
		fmt.Println(noticerow.Date)
		fmt.Println(" ")
	}
	e, err := json.MarshalIndent(classNoticeRow, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(e))
	classnoticerow := classNoticeRow
	// classnoticerow[10].CourseName = "hogehoge"
	e2, err := json.MarshalIndent(classnoticerow, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	if diff, _ := JSONBytesEqual(e, e2); diff == true {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
	f, err := os.Create("../hiniti.json")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString(string(e))
	f2, err := os.Create("../hiniti2.json")
	if err != nil {
		t.Fatal(err)
	}
	f2.WriteString(string(e2))
}

func TestSeisekiRows(t *testing.T) {
	kc, err := c.NewKyoumuClient()
	if err != nil {
		t.Fatal(err)
	}
	rows, err := kc.SeisekiRows()
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range rows {
		fmt.Println(*row)
	}
	e, err := json.MarshalIndent(rows, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(e))
	f, err := os.Create("seiseki.json")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString(string(e))
}
