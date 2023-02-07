package model

import "time"

type Video struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Duration  int       `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Annotation struct {
	ID        int       `json:"id"`
	VideoID   int       `json:"video_id"`
	StartTime int       `json:"start_time"`
	EndTime   int       `json:"end_time"`
	Type      string    `json:"type"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
