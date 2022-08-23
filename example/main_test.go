package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/lll-lll-lll-lll/conread/connpass"
	"github.com/lll-lll-lll-lll/conread/format"
	"github.com/stretchr/testify/assert"
)

func TestCreateMDFile(t *testing.T) {
	n := "TestREADME.md"
	file, err := os.Create(n)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer os.Remove(n)
}

func TestTimeCompare(t *testing.T) {
	a := "2022-07-24T15:00:00+09:00"
	now := time.Now()
	t5, _ := time.Parse(time.RFC3339, a)
	fmt.Println(t5.After(now))
}

func TestChangeTimeFormat(t *testing.T) {
	weekdaymonthja := strings.NewReplacer(
		"Sunday", "日",
		"Monday", "月",
		"Tueday", "火",
		"Wednesday", "水",
		"Thursday", "木",
		"Friday", "金",
		"Saturday", "土",
		"January", "1月",
		"February", "2月",
		"March", "3月",
		"April", "4月",
		"May", "5月",
		"June", "6月",
		"July", "7月",
		"August", "8月",
		"September", "9月",
		"October", "10月",
		"November", "11月",
		"December", "12月",
	)
	tu := "2022-08-17T19:00:00+09:00"
	p, _ := time.Parse(time.RFC3339, tu)
	f := func(p time.Time) string {
		if p.Minute() == 0 {
			return "00"
		} else {
			return strconv.Itoa(p.Minute())
		}
	}
	str := fmt.Sprintf("%s月%d日(%s) %d:%s", p.Month().String(), p.Day(), p.Weekday(), p.Hour(), f(p))
	fmt.Println(weekdaymonthja.Replace(str))

}

func TestYYMM(t *testing.T) {
	weekdaymonthja := strings.NewReplacer(
		"Sunday", "日",
		"Monday", "月",
		"Tueday", "火",
		"Wednesday", "水",
		"Thursday", "木",
		"Friday", "金",
		"Saturday", "土",
		"January", "01",
		"February", "02",
		"March", "03",
		"April", "04",
		"May", "05",
		"June", "06",
		"July", "07",
		"August", "08",
		"September", "09",
		"October", "10",
		"November", "11",
		"December", "12",
	)
	p := time.Now()
	str := weekdaymonthja.Replace(fmt.Sprintf("%d%s", p.Year(), p.Month().String()))
	fmt.Println(str)
}

func TestForThreeMonthEvents(t *testing.T) {
	p := time.Now()
	yearmonthsja := strings.NewReplacer(
		"January", "01",
		"February", "02",
		"March", "03",
		"April", "04",
		"May", "05",
		"June", "06",
		"July", "07",
		"August", "08",
		"September", "09",
		"October", "10",
		"November", "11",
		"December", "12",
	)
	s := yearmonthsja.Replace(fmt.Sprintf("%d%s", p.Year(), p.Month().String()))
	fmt.Println(s)
	fmt.Println(format.GetForThreeMonthsEvent())
}

func TestConnpassAPI(t *testing.T) {
	connpass, _ := connpass.NewConnpass("Shun_Pei")
	fmt.Println(connpass.ConnpassUSER)
	assert.Equal(t, connpass.ConnpassUSER, "Shun_Pei")
}
