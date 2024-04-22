package gapi

import (
	"fmt"

	db "github.com/Dee-Dee-Tiger-Hacks/Backend-Server/db/sqlc"
	util "github.com/Dee-Dee-Tiger-Hacks/Backend-Server/db/utils"
	mongodb "github.com/Dee-Dee-Tiger-Hacks/Backend-Server/mongodb/repositories"
	"github.com/Dee-Dee-Tiger-Hacks/Backend-Server/pb"
	"github.com/Dee-Dee-Tiger-Hacks/Backend-Server/token"
	"github.com/Dee-Dee-Tiger-Hacks/Backend-Server/worker"
)

type Server struct {
	pb.UnimplementedDeeDeeServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
	mongoDBStore    mongodb.Store
}

func NewServer(config util.Config, store db.Store, mongoDBStore mongodb.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		mongoDBStore:    mongoDBStore,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
