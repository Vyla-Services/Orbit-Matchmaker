package models

type Player struct {
	AccountID string `json:"accountId"`
	Playlist  string `json:"playlist"`
	Region    string `json:"region"`
}
