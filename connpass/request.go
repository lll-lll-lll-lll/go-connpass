package connpass

import (
	"net/url"
	"path"
	"strconv"
)

var (
	_ Request = (*UserRequest)(nil)
	_ Request = (*EventRequest)(nil)
)

type APIType string

const (
	EVENT_PATH APIType = "event"
	USER_PATH  APIType = "user"
)

type Request interface {
	// ToQueryParameter convert request to url.Values
	ToQueryParameter() url.Values
	// URL request url
	URL() string
}

type UserRequest struct {
	// リクエスト先URL
	Path APIType `json:"-"`
	// ニックネーム.指定したニックネームのユーザを検索します。複数指定可能です*
	NickName []string `json:"nickname[]"`
	// 検索の開始位置. 検索結果の何件目から出力するかを指定します。デフォルトは1です。
	Start int `json:"start"`
	// 検索結果の取得件数
	Count int `json:"count"`
	// 出力形式. json固定
	Format string `json:"format"`
}

func (u *UserRequest) URL() string {
	return path.Join(CONNPASSAPI_V1, string(u.Path))
}

func (u *UserRequest) ToQueryParameter() url.Values {
	q := url.Values{}
	if len(u.NickName) != 0 {
		q.Add("nickname", join(u.NickName))
	}
	if u.Start != 0 {
		q.Add("start", strconv.Itoa(u.Start))
	}
	if u.Count != 0 {
		q.Add("count", strconv.Itoa(u.Count))
	}
	if u.Format == "" {
		u.Format = "json"
		q.Add("format", u.Format)
	}
	return q
}

type EventRequest struct {
	// リクエスト先URL
	Path APIType `json:"-"`
	// イベント毎に割り当てられた番号で検索します。複数指定可能です*
	EventIDList []int `json:"event_id[]"`
	// キーワード. イベントのタイトル、キャッチ、概要、住所をAND条件部分一致で検索します。複数指定可能です*
	Keyword []string `json:"keyword[]"`
	// キーワード. イベントのタイトル、キャッチ、概要、住所をOR条件部分一致で検索します。複数指定可能です*
	KeywordOR []string `json:"keyword_or[]"`
	// イベント開催年月
	// 指定した年月に開催されているイベントを検索します。複数指定可能です*
	// yyyymm形式で指定してください。例: 201810
	YM []string `json:"ym[]"`
	// イベント開催年月日
	// 指定した年月日に開催されているイベントを検索します。複数指定可能です*
	YMD []string `json:"ymd[]"`
	// 指定したニックネームのユーザが参加しているイベントを検索します。複数指定可能です*
	NickName []string `json:"nickname[]"`
	// 指定したニックネームのユーザが管理しているイベントを検索します。複数指定可能です*
	OwnerNickName []string `json:"owner_nickname[]"`
	// グループ 毎に割り当てられた番号で、ひもづいたイベントを検索します。複数指定可能です*
	SeriesID int `json:"series_id"`
	// 検索結果の何件目から出力するかを指定します。
	Start int `json:"start"`
	// 検索結果の表示順
	// 検索結果の表示順を、更新日時順、開催日時順、新着順で指定します。
	// 1: 更新日時順
	// 2: 開催日時順
	// 3: 新着順
	// (初期値: 1)
	Order int `json:"order"`
	// 検索結果の最大出力データ数を指定します。
	// 初期値: 10、最小値：1、最大値：100
	Count int `json:"count"`
	// 出力形式. json固定
	Format string `json:"format"`
}

func join(reqQuery []string) string {
	var res string
	for i, v := range reqQuery {
		if i == 0 {
			res = v
		} else {
			res += "," + v
		}
	}
	return res
}

func (e *EventRequest) joinInt(sliceData []int) string {
	var res string
	for i, v := range sliceData {
		if i == 0 {
			res = strconv.Itoa(v)
		} else {
			res += "," + strconv.Itoa(v)
		}
	}
	return res
}

func (e *EventRequest) URL() string {
	return path.Join(CONNPASSAPI_V1, string(e.Path))
}

// ToURLVal リクエストに詰め込まれている値をURL.Valuesに変換する
func (e *EventRequest) ToQueryParameter() url.Values {
	q := url.Values{}
	if len(e.EventIDList) != 0 {
		q.Add("event_id", e.joinInt(e.EventIDList))
	}
	if len(e.Keyword) != 0 {
		q.Add("keyword", join(e.Keyword))
	}
	if len(e.KeywordOR) != 0 {
		q.Add("keyword_or", join(e.KeywordOR))
	}
	if len(e.YM) != 0 {
		q.Add("ym", join(e.YM))
	}
	if len(e.YMD) != 0 {
		q.Add("ymd", join(e.YMD))
	}
	if len(e.NickName) != 0 {
		q.Add("nickname", join(e.NickName))
	}
	if len(e.OwnerNickName) != 0 {
		q.Add("owner_nickname", join(e.OwnerNickName))
	}
	if e.SeriesID != 0 {
		q.Add("series_id", strconv.Itoa(e.SeriesID))
	}
	if e.Start != 0 {
		q.Add("start", strconv.Itoa(e.Start))
	}
	if e.Order != 0 {
		q.Add("order", strconv.Itoa(e.Order))
	}
	if e.Count != 0 {
		q.Add("count", strconv.Itoa(e.Count))
	}
	if e.Format == "" {
		e.Format = "json"
		q.Add("format", e.Format)
	}
	return q
}
