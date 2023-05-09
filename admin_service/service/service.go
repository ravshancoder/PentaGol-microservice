package service

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	u "github.com/PentaGol/admin_service/genproto/admin"
	"github.com/PentaGol/admin_service/pkg/logger"
	grpcclient "github.com/PentaGol/admin_service/service/grpc_client"
	"github.com/PentaGol/admin_service/storage"
)

type AdminService struct {
	storage storage.IStorage
	Logger  logger.Logger
	Client  grpcclient.Clients
}

func NewAdminService(db *sqlx.DB, log logger.Logger, client grpcclient.Clients) *AdminService {
	return &AdminService{
		storage: storage.NewStoragePg(db),
		Logger:  log,
		Client:  client,
	}
}


func (s *AdminService) CreateAdmin(ctx context.Context, req *u.AdminRequest) (*u.AdminResponse, error) {
	res, err := s.storage.Admin().CreateAdmin(req)
	if err != nil {
		log.Println("failed to creating Admin: ", err)
		return &u.AdminResponse{}, err
	}

	return res, nil
}


func (s *AdminService) GetAdminById(ctx context.Context, req *u.IdRequest) (*u.AdminResponse, error) {
	res, err := s.storage.Admin().GetAdminById(req)
	if err != nil {
		log.Println("failed to getting Admin: ", err)
		return &u.AdminResponse{}, err
	}

	return res, nil
}

func (s *AdminService) CheckField(ctx context.Context, req *u.CheckFieldReq) (*u.CheckFieldRes, error) {
	res, err := s.storage.Admin().CheckFiedld(req)
	if err != nil {
		s.Logger.Error("error delete", logger.Any("Error delete admins", err))
		return &u.CheckFieldRes{}, status.Error(codes.Internal, "something went wrong, please check admin info")
	}
	return res, nil
}

func (s *AdminService) UpdateToken(ctx context.Context, req *u.RequestForTokens) (*u.AdminResponse, error) {
	res, err := s.storage.Admin().UpdateToken(req)
	if err != nil {
		log.Println("failed to updating admin token: ", err)
		return &u.AdminResponse{}, err
	}

	return res, err
}

func (s *AdminService) GetByEmail(ctx context.Context, req *u.EmailReq) (*u.AdminResponse, error) {
	customer, err := s.storage.Admin().GetByEmail(req)
	if err != nil {
		s.Logger.Error("Error while getting admin info by email", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, "Something went wrong")
	}
	return customer, nil
}
