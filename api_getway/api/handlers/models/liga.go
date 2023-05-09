package models

// liga
// ---------------------------------------
type LigaRequest struct {
	Name string `json:"name"`

}

type LigaResponse struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Ligas struct {
	Ligas []LigaResponse `json:"ligas"`
}

// game
// ---------------------------------------
type GameRequest struct {
	Time             string `json:"time"`
	Condtion         bool   `json:"condition"`
	FirstTeamId      int64  `json:"first_team_id"`
	SecondTeamId     int64  `json:"second_team_id"`
	ResultFirstTeam  int64  `json:"result_first_team"`
	ResultSecondTeam int64  `json:"result_second_team"`
	LigaId           int64  `json:"liga_id"`
}

type GameResponse struct {
	Id               int64  `json:"id"`
	Time             string `json:"time"`
	Condtion         bool   `json:"condition"`
	FirstTeamId      int64  `json:"first_team_id"`
	SecondTeamId     int64  `json:"second_team_id"`
	ResultFirstTeam  int64  `json:"result_first_team"`
	ResultSecondTeam int64  `json:"result_second_team"`
	FirstTeamPoint   int64  `json:"first_team_point"`
	SecondTeamPoint  int64  `json:"second_team_point"`
	LigaId           int64  `json:"liga_id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type Games struct {
	Games []GameResponse `json:"games"`
}


// Club
// ---------------------------------------
type ClubRequest struct {
	Name string `json:"name"`
	Points int `json:"points"`
}

type ClubResponse struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Points int `json:"points"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Clubs struct {
	Clubs []ClubResponse `json:"clubs"`
}