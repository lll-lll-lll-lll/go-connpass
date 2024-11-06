package connpass

type UserResponse struct {
	// 含まれる検索結果の件数
	ResultsReturned int `json:"results_returned"`
	// 検索結果の総件数
	ResultsAvailable int `json:"results_available"`
	// 検索の開始位置
	ResultsStart int `json:"results_start"`
	// 検索結果のユーザーリスト
	Users []struct{}
}

type User struct {
	// ユーザーID
	UserID int `json:"user_id"`
	// ニックネーム
	Nickname string `json:"nickname"`
	// 表示名
	DisplayName string `json:"display_name"`
	// 自己紹介
	Description string `json:"description"`
	// connpass.com 上のURL
	UserURL string `json:"user_url"`
	// ユーザのサムネ画像のURL. このURLはある程度の時間で失効されます。外部サイトでの直接参照などはご遠慮ください。
	UserImageURL string `json:"user_image_url"`
	// 利用開始日時 (ISO-8601形式)
	CreatedAt string `json:"created_at"`
	// 参加イベント数
	AttendedEventCount int `json:"attended_event_count"`
	// 管理イベント数
	OrganizeEventCount int `json:"organize_event_count"`
	// 発表イベント数
	PresenterEventCount int `json:"presenter_event_count"`
	// ブックマークイベント数
	BookmarkEventCount int `json:"bookmark_event_count"`
}

type Event struct {
	Series struct {
		// グループID
		Id int `json:"id"`
		// グループのタイトル
		Title string `json:"title"`
		// グループのconnpass.com 上のURL
		Url string `json:"url"`
	} `json:"series"`
	// 管理者のニックネーム
	OwnerNickname string `json:"owner_nickname"`
	// キャッチコピー
	Catch string `json:"catch"`
	// 概要(HTML形式)
	Description string `json:"description"`
	// connpass.com 上のURL
	EventUrl string `json:"event_url"`
	// Twitterのハッシュタグ
	HashTag string `json:"hash_tag"`
	// イベント開催日時 (ISO-8601形式)
	StartedAt string `json:"started_at"`
	// イベント終了日時 (ISO-8601形式)
	EndedAt string `json:"ended_at"`
	// 管理者の表示名
	OwnerDisplayName string `json:"owner_display_name"`
	// イベント参加タイプ
	EventType string `json:"event_type"`
	// タイトル
	Title string `json:"title"`
	// 開催場所
	Address string `json:"address"`
	// 開催会場
	Place string `json:"place"`
	// 更新日時 (ISO-8601形式)
	UpdatedAt string `json:"updated_at"`
	// イベントID
	EventId int `json:"event_id"`
	// 管理者のID
	OwnerId int `json:"owner_id"`
	// 定員
	Limit int `json:"limit"`
	// 参加者数
	Accepted int `json:"accepted"`
	// 補欠者数
	Waiting int `json:"waiting"`
	// 開催会場の緯度
	Lat string `json:"lat"`
	// 開催会場の経度
	Lon string `json:"lon"`
}

// EventResponse connpass api response
// https://connpass.com/about/api/
type EventResponse struct {
	// レスポンスに含まれる検索結果の件数
	ResultsReturned int `json:"results_returned"`
	// 検索結果の総件数
	ResultsAvailable int `json:"results_available"`
	// 検索の開始位置
	ResultsStart int `json:"results_start"`
	// 検索結果のイベントリスト
	Events []Event `json:"events"`
}
