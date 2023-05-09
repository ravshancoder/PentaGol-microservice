package repo

import (
	p "github.com/PentaGol/liga_service/genproto/liga"
)

type LigaStorageI interface {
	CreateLiga(*p.LigaRequest) (*p.LigaResponse, error)
	GetLigaById(*p.IdRequest) (*p.LigaResponse, error)
	GetAllLigas(*p.AllLigaRequest) (*p.Ligas, error)
	DeleteLiga(*p.IdRequest) (*p.LigaResponse, error)
}

type GameStorageI interface {
	CreateGame(*p.GameRequest) (*p.GameResponse, error)
	GetGameById(*p.IdRequest) (*p.GameResponse, error)
	GetAllGames(*p.AllGameRequest) (*p.Games, error)
	DeleteGame(*p.IdRequest) (*p.GameResponse, error)
}

type ClubStorageI interface {
	CreateClub(*p.ClubRequest) (*p.ClubResponse, error)
	GetClubById(*p.IdRequest) (*p.ClubResponse, error)
	GetAllClubs(*p.AllClubRequest) (*p.Clubs, error)
}