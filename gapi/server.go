package gapi

import (
	"fmt"

	db "github.com/MidNight91119/simplebank/db/sqlc"
	"github.com/MidNight91119/simplebank/db/util"
	"github.com/MidNight91119/simplebank/pb"
	"github.com/MidNight91119/simplebank/token"
)

// Server serves GRPC requests for our banking services
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	tokenMaker token.Maker
	store      db.Store
}

// NewServer creates a new GRPC server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
