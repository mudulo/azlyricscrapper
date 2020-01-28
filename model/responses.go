package model

type LyricsResponse struct {
	Code int `json:"code"` 
	Lyrics Lyric `json:"lyrics"`
}

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

