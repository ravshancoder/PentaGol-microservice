package service

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"

	p "github.com/PentaGol/liga_service/genproto/liga"
	"github.com/PentaGol/liga_service/pkg/logger"
	grpcclient "github.com/PentaGol/liga_service/service/grpc_client"
	"github.com/PentaGol/liga_service/storage"
)

type LigaService struct {
	storage storage.IStorage
	Logger  logger.Logger
	Client  grpcclient.Clients
}

func NewLigaService(db *sqlx.DB, log logger.Logger, client grpcclient.Clients) *LigaService {
	return &LigaService{
		storage: storage.NewStoragePg(db),
		Logger:  log,
		Client:  client,
	}
}
// liga
func (s *LigaService) CreateLiga(ctx context.Context, req *p.LigaRequest) (*p.LigaResponse, error) {
	res, err := s.storage.Liga().CreateLiga(req)
	if err != nil {
		log.Println("failed to create Liga: ", err)
		return &p.LigaResponse{}, err
	}

	return res, nil
}

func (s *LigaService) GetLigaById(ctx context.Context, req *p.IdRequest) (*p.LigaResponse, error) {
	res, err := s.storage.Liga().GetLigaById(req)
	if err != nil {
		log.Println("failed to get Liga by id: ", err)
		return &p.LigaResponse{}, err
	}

	return res, nil
}

func (s *LigaService) GetAllLigas(ctx context.Context, req *p.AllLigaRequest) (*p.Ligas, error) {
	res, err := s.storage.Liga().GetAllLigas(req)
	if err != nil {
		log.Println("failed to get all Liga: ", err)
		return &p.Ligas{}, err
	}

	return res, nil
}

func (s *LigaService) DeleteLiga(ctx context.Context, req *p.IdRequest) (*p.LigaResponse, error) {
	res, err := s.storage.Liga().DeleteLiga(req)
	if err != nil {
		log.Println("failed to delete Liga: ", err)
		return &p.LigaResponse{}, err
	}

	return res, err
}


// game
func (s *LigaService) CreateGame(ctx context.Context, req *p.GameRequest) (*p.GameResponse, error) {
	res, err := s.storage.Game().CreateGame(req)
	if err != nil {
		log.Println("failed to create Game: ", err)
		return &p.GameResponse{}, err
	}

	return res, nil
}

func (s *LigaService) GetGameById(ctx context.Context, req *p.IdRequest) (*p.GameResponse, error) {
	res, err := s.storage.Game().GetGameById(req)
	if err != nil {
		log.Println("failed to get Game by id: ", err)
		return &p.GameResponse{}, err
	}

	return res, nil
}

func (s *LigaService) GetAllGames(ctx context.Context, req *p.AllGameRequest) (*p.Games, error) {
	res, err := s.storage.Game().GetAllGames(req)
	if err != nil {
		log.Println("failed to get all Game: ", err)
		return &p.Games{}, err
	}

	return res, nil
}

func (s *LigaService) DeleteGame(ctx context.Context, req *p.IdRequest) (*p.GameResponse, error) {
	res, err := s.storage.Game().DeleteGame(req)
	if err != nil {
		log.Println("failed to delete Game: ", err)
		return &p.GameResponse{}, err
	}

	return res, err
}

// club
func (s *LigaService) CreateClub(ctx context.Context, req *p.ClubRequest) (*p.ClubResponse, error) {
	res, err := s.storage.Club().CreateClub(req)
	if err != nil {
		log.Println("failed to create Club: ", err)
		return &p.ClubResponse{}, err
	}

	return res, nil
}

func (s *LigaService) GetClubById(ctx context.Context, req *p.IdRequest) (*p.ClubResponse, error) {
	res, err := s.storage.Club().GetClubById(req)
	if err != nil {
		log.Println("failed to get Club by id: ", err)
		return &p.ClubResponse{}, err
	}

	return res, nil
}

func (s *LigaService) GetAllClubs(ctx context.Context, req *p.AllClubRequest) (*p.Clubs, error) {
	res, err := s.storage.Club().GetAllClubs(req)
	if err != nil {
		log.Println("failed to get all Club: ", err)
		return &p.Clubs{}, err
	}

	return res, nil
}