package gapi

import (
	db "github.com/Dee-Dee-Tiger-Hacks/Backend-Server/db/sqlc"
	util "github.com/Dee-Dee-Tiger-Hacks/Backend-Server/db/utils"
	mongodb "github.com/Dee-Dee-Tiger-Hacks/Backend-Server/mongodb/repositories"
	"github.com/Dee-Dee-Tiger-Hacks/Backend-Server/pb"
	"github.com/Dee-Dee-Tiger-Hacks/Backend-Server/token"
)

type AuthServer struct {
	pb.UnimplementedDeeDeeServer
	config       util.Config
	store        db.Store
	tokenMaker   token.Maker
	mongoDBStore mongodb.Store
}

func NewAuthServer(config util.Config, store db.Store, mongoDBStore mongodb.Store, tokenMaker token.Maker) (*AuthServer, error) {
	server := &AuthServer{
		config:       config,
		store:        store,
		mongoDBStore: mongoDBStore,
		tokenMaker:   tokenMaker,
	}

	return server, nil
}
