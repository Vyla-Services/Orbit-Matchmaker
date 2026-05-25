package models

type Server struct {
	ID            string `json:"id"`
	IP            string `json:"ip"`
	Port          int    `json:"port"`
	Region        string `json:"region"`
	Capacity      int    `json:"capacity"`
	CurrentLoad   int    `json:"currentLoad"`
	LastHeartbeat int64  `json:"lastHeartbeat"`
}
