package playlist

import (
	"os"
	"pkg/models"
	"strconv"
	"strings"
)

var playlists []models.Playlist
var regions []string

func Load() {
	rawPlaylists := os.Getenv("ORBIT_PLAYLISTS")
	rawRegions := os.Getenv("ORBIT_REGIONS")

	if rawPlaylists == "" {
		rawPlaylists = "Playlist_Showdown_Solo.Playlist_Showdown_Solo:100"
	}

	if rawRegions == "" {
		rawRegions = "NA,EU"
	}

	ps := strings.Split(rawPlaylists, ",")
	for _, p := range ps {
		parts := strings.Split(p, ":")
		if len(parts) != 2 {
			continue
		}
		n := parts[0]
		max, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		playlists = append(playlists, models.Playlist{Name: n, MaxPlayers: max})
	}

	regions = strings.Split(rawRegions, ",")
}

func Playlists() []models.Playlist {
	return playlists
}

func Regions() []string {
	return regions
}
