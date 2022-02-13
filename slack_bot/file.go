package slack_bot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sort"

	"github.com/szpp-dev-team/gakujo-api/model"
)

type SeisekiSubject struct {
	Subject string
}

type ByAge []*model.SeisekiRow

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Year < a[j].Year }

// JSONBytesEqual compares the JSON in two byte slices.
// True => 一致
// False => 不一致
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

func UpdateSeisekiFile(rows []*model.SeisekiRow) error {
	Seiseki := make([]SeisekiSubject, 0)
	sort.Slice(rows, func(i, j int) bool { return rows[j].Date.After(rows[i].Date) })

	for _, subjectname := range rows {
		var seiseki SeisekiSubject
		seiseki.Subject = subjectname.SubjectName
		Seiseki = append(Seiseki, seiseki)
	}

	subjectnameJson, err := json.MarshalIndent(Seiseki, "", " ")
	if err != nil {
		return err
	}

	f, err := os.Open("seiseki.json")
	defer f.Close()
	if err != nil {
		new, err := os.Create("seiseki.json")
		if err != nil {
			return err
		}
		defer new.Close()
		new.WriteString(string(subjectnameJson))
	} else {
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		if diff, _ := JSONBytesEqual(subjectnameJson, b); diff {
			BotSame()
		} else {
			var pastSeiseki []SeisekiSubject
			err := json.Unmarshal(b, &pastSeiseki)
			if err != nil {
				return err
			}

			updata, err := os.Create("seiseki.json")
			if err != nil {
				return err
			}
			updata.WriteString(string(subjectnameJson))
			index := 0
			for i, row := range rows {
				if row.Year == 2021 {
					index = i
					break
				}
			}
			row := rows[index:]
			change := Seiseki[len(pastSeiseki):]
			changeRows := rows[len(pastSeiseki):]
			log.Println(change)
			BotNew(row, change, changeRows)
		}
	}
	return nil
}
