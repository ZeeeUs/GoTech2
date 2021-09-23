package model

type Event struct {
	UserID int    `json:"user_id"`
	Text   string `json:"text"`
	Date   string `json:"date"`
}
