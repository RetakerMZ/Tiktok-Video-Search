package domain

type SearchResult struct {
	Code          int              `json:"code"`
	Msg           string           `json:"msg"`
	ProcessedTime float64          `json:"processed_time"`
	Data          SearchResultData `json:"data"`
}

type SearchResultData struct {
	Cursor  int                 `json:"cursor"`
	HasMore bool                `json:"hasMore"`
	Videos  []SearchResultVideo `json:"videos"`
}

type SearchResultVideo struct {
	AwemeID             string             `json:"aweme_id"`
	VideoID             string             `json:"video_id"`
	Region              string             `json:"region"`
	Title               string             `json:"title"`
	Cover               string             `json:"cover"`
	AiDynamicCover      string             `json:"ai_dynamic_cover"`
	OriginCover         string             `json:"origin_cover"`
	Duration            int                `json:"duration"`
	Play                string             `json:"play"`
	Wmplay              string             `json:"wmplay"`
	Size                int                `json:"size"`
	WmSize              int                `json:"wm_size"`
	Music               string             `json:"music"`
	MusicInfo           SearchResultMusic  `json:"music_info"`
	PlayCount           int                `json:"play_count"`
	DiggCount           int                `json:"digg_count"`
	CommentCount        int                `json:"comment_count"`
	ShareCount          int                `json:"share_count"`
	DownloadCount       int                `json:"download_count"`
	CreateTime          int                `json:"create_time"`
	IsAd                bool               `json:"is_ad"`
	CommercialVideoInfo string             `json:"commercial_video_info"`
	ItemCommentSettings int                `json:"item_comment_settings"`
	MentionedUsers      string             `json:"mentioned_users"`
	Author              SearchResultAuthor `json:"author"`
	IsTop               int                `json:"is_top"`
}

type SearchResultMusic struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Play     string `json:"play"`
	Cover    string `json:"cover"`
	Author   string `json:"author"`
	Original bool   `json:"original"`
	Duration int    `json:"duration"`
	Album    string `json:"album"`
}

type SearchResultAuthor struct {
	ID       string `json:"id"`
	UniqueID string `json:"unique_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type SearchService interface {
	GetSearch(query string, cursor string) (SearchResult, error)
}
