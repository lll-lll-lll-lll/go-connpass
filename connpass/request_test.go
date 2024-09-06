package connpass

import (
	"fmt"
	"testing"
)

func Test_Event_ToURLVal(t *testing.T) {
	t.Parallel()
	t.Run("Test_Event_ToURLVal", func(t *testing.T) {
		t.Parallel()
		e := &EventRequest{}
		e.EventIDList = []int{1, 2, 3}
		e.Keyword = []string{"golang", "python"}
		e.KeywordOR = []string{"golang", "python"}
		e.YM = []string{"202201", "202202"}
		e.YMD = []string{"20220101", "20220202"}
		e.NickName = []string{"Shun_Pei", "Shun_Pei"}
		e.OwnerNickName = []string{"Shun_Pei", "Shun_Pei"}
		e.SeriesID = 1
		e.Start = 1
		e.Order = 1
		e.Count = 1
		q := e.ToQueryParameter()
		if len(q) != 12 {
			fmt.Println(len(q))
			t.Errorf("failed to ToURLVal. %v", q)
		}
	})
}
