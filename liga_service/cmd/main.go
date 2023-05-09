package main

import (
	"fmt"
	"net"

	"github.com/PentaGol/liga_service/config"
	p "github.com/PentaGol/liga_service/genproto/liga"
	"github.com/PentaGol/liga_service/pkg/db"
	"github.com/PentaGol/liga_service/pkg/logger"
	"github.com/PentaGol/liga_service/service"
	grpcclient "github.com/PentaGol/liga_service/service/grpc_client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "golang")
	defer logger.Cleanup(log)

	connDb, err := db.ConnectToDB(cfg)
	if err != nil {
		fmt.Println("failed connect database", err)
	}

	grpcClient, err := grpcclient.New(cfg)
	if err != nil {
		fmt.Println("failed while grpc client", err.Error())
	}

	ligaService := service.NewLigaService(connDb, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.LigaServicePort)
	if err != nil {
		log.Fatal("failed while listening port: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	p.RegisterLigaServiceServer(s, ligaService)

	log.Info("main: server running",
		logger.String("port", cfg.LigaServicePort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed while listening: %v", logger.Error(err))
	}

}
