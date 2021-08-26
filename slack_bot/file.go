package slack_bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"

	"github.com/szpp-dev-team/gakujo-api/model"
)

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

func File(rows []*model.SeisekiRow) error {
	e, err := json.MarshalIndent(rows, "", " ")
	if err != nil {
		return err
	}

	f, err := os.Open("../seiseki.json")
	if err != nil {
		return err
	}
	defer f.Close()
	if err != nil {
		new, err := os.Create("../seiseki.json")
		if err != nil {
			return err
		}
		defer new.Close()
		new.WriteString(string(e))
	} else {
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		if diff, _ := JSONBytesEqual(e, b); diff {
			Bot_same()
		} else {
			updata, err := os.Create("../seiseki.json")
			if err != nil {
				return err
			}
			updata.WriteString(string(e))
			fmt.Printf("\n")
			sort.Slice(rows, func(i, j int) bool { return rows[i].Year < rows[j].Year })
			index := 0
			for i, row := range rows {
				if row.Year == 2021 {
					index = i
					break
				}
			}
			row := rows[index:]
			Bot_new(row)
		}
	}
	return nil
}
