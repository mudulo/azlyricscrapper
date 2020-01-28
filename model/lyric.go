package model

type Lyric struct {
	Title string `json:"title"` 
	Artist string `json:"artist"`
	Body string `json:"body"`
	Copyright string `json:"copyright"`
}