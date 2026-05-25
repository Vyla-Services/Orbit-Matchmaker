package models

type Session struct {
	ID       string   `json:"id"`
	Playlist string   `json:"playlist"`
	Region   string   `json:"region"`
	Players  []string `json:"players"`
	ServerIP string   `json:"serverIp"`
}
