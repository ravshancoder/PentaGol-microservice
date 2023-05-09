package grpcclient

import (
	"fmt"

	"github.com/PentaGol/post_service/config"
	cu "github.com/PentaGol/post_service/genproto/admin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients interface {
	User() cu.AdminServiceClient
}

type ServiceManager struct {
	Config         config.Config
	userService    cu.AdminServiceClient
}

func New(cfg config.Config) (*ServiceManager, error) {
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.AdminServiceHost, cfg.AdminServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("user service dial host:%s, port:%s", cfg.AdminServiceHost, cfg.AdminServicePort)
	}

	return &ServiceManager{
		Config:         cfg,
		userService:    cu.NewAdminServiceClient(connUser),
	}, nil
}

func (s *ServiceManager) User() cu.AdminServiceClient {
	return s.userService
}
