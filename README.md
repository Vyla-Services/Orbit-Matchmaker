# Orbit Matchmaker

Orbit is a matchmaker!

---

## Setup

### 1. Install Go
Go 1.22 or newer is required.

https://go.dev/dl/

### 2. Clone the repository
git clone <your-repo-url>
cd Orbit-Matchmaker

Code

### 3. Configure environment variables
Copy the example file:

cp .env.example .env

Code

Edit `.env`:

ORBIT_ADDR=:8080
ORBIT_REGIONS=NA,EU
ORBIT_PLAYLISTS=Playlist_Showdown_Solo.Playlist_Showdown_Solo:100

Code

### 4. Install dependencies
cd orbit
go mod tidy

Code

### 5. Run Orbit
go run ./cmd/orbit

Code

Orbit will start on the address defined in `ORBIT_ADDR`.

---

## Usage

### Queue a player
POST /queue
Content-Type: application/json

{
"accountId": "123",
"playlist": "Playlist_Showdown_Solo.Playlist_Showdown_Solo",
"region": "NA"
}

Code

### Check session status
GET /session/123

Code

If the player has been matched, Orbit returns:

{
"id": "20260525010101",
"playlist": "Playlist_Showdown_Solo.Playlist_Showdown_Solo",
"region": "NA",
"players": ["123", "..."],
"serverIp": "127.0.0.1"
}

Code

### Register a server
POST /server/register
Content-Type: application/json

{
"id": "server1",
"ip": "127.0.0.1",
"port": 7777,
"region": "NA",
"capacity": 100
}

Code

### Send server heartbeat
POST /server/heartbeat
Content-Type: application/json

{
"id": "server1"
}

Code

---

## Matchmaking Loop

Orbit automatically:
- pulls players from queues
- groups them by playlist + region
- assigns the lowest-load server
- creates a session
- exposes it via `/session/:accountId`



---

## Build a production binary

go build -o orbit ./cmd/orbit

Code

Run it:

./orbit