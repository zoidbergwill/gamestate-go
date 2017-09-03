package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Provider struct {
	Name      string `json:"name"`
	Appid     int    `json:"appid"`
	Version   int    `json:"version"`
	Steamid   string `json:"steamid"`
	Timestamp int    `json:"timestamp"`
}
type Map struct {
	Mode   string `json:"mode"`
	Name   string `json:"name"`
	Phase  string `json:"phase"`
	Round  int    `json:"round"`
	TeamCt struct {
		Score                int `json:"score"`
		TimeoutsRemaining    int `json:"timeouts_remaining"`
		MatchesWonThisSeries int `json:"matches_won_this_series"`
	} `json:"team_ct"`
	TeamT struct {
		Score                int `json:"score"`
		TimeoutsRemaining    int `json:"timeouts_remaining"`
		MatchesWonThisSeries int `json:"matches_won_this_series"`
	} `json:"team_t"`
	NumMatchesToWinSeries int `json:"num_matches_to_win_series"`
	CurrentSpectators     int `json:"current_spectators"`
	SouvenirsTotal        int `json:"souvenirs_total"`
}
type Round struct {
	Phase   string `json:"phase"`
	WinTeam string `json:"win_team"`
}
type PlayerState struct {
	Health      int  `json:"health"`
	Armor       int  `json:"armor"`
	Helmet      bool `json:"helmet"`
	Defusekit   bool `json:"defusekit"`
	Flashed     int  `json:"flashed"`
	Smoked      int  `json:"smoked"`
	Burning     int  `json:"burning"`
	Money       int  `json:"money"`
	RoundKills  int  `json:"round_kills"`
	RoundKillhs int  `json:"round_killhs"`
	EquipValue  int  `json:"equip_value"`
}

type PlayerMatchStats struct {
	Kills   int `json:"kills"`
	Assists int `json:"assists"`
	Deaths  int `json:"deaths"`
	Mvps    int `json:"mvps"`
	Score   int `json:"score"`
}

type Weapon struct {
	Name        string `json:"name"`
	Paintkit    string `json:"paintkit"`
	Type        string `json:"type"`
	AmmoClip    int    `json:"ammo_clip"`
	AmmoClipMax int    `json:"ammo_clip_max"`
	AmmoReserve int    `json:"ammo_reserve"`
	State       string `json:"state"`
}

type Player struct {
	Steamid      string            `json:"steamid"`
	Clan         string            `json:"clan"`
	Name         string            `json:"name"`
	ObserverSlot int               `json:"observer_slot"`
	Team         string            `json:"team"`
	Activity     string            `json:"activity"`
	State        PlayerState       `json:"state"`
	MatchStats   PlayerMatchStats  `json:"match_stats"`
	Weapons      map[string]Weapon `json:"weapons"`
	Position     string            `json:"position"`
}

type PhaseCountdowns struct {
	Phase       string `json:"phase"`
	PhaseEndsIn string `json:"phase_ends_in"`
}

type Auth struct {
	Token string `json:"token"`
}

type Message struct {
	Provider        Provider          `json:"provider"`
	Map             Map               `json:"map"`
	Round           Round             `json:"round"`
	Player          Player            `json:"player"`
	Players         map[string]Player `json:"allplayers"`
	PhaseCountdowns PhaseCountdowns   `json:"phase_countdowns"`
	Previously      interface{}       `json:"previously"`
	Auth            Auth              `json:"auth"`
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		fmt.Println(string(body))
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		var message Message
		err = json.Unmarshal(body, &message)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", message)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 8080. Go to http://127.0.0.1:8080/")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	log.Fatal(err)
}
