package server

import (
	"database/sql"

	"github.com/Nyarum/noterius/core"
	log "github.com/Sirupsen/logrus"
	"go.uber.org/zap"

	"net"
)

type Server struct {
	config   core.Config
	database *sql.DB
	logger   *zap.SugaredLogger
}

func NewServer(config core.Config, database *sql.DB, logger *zap.SugaredLogger) *Server {
	return &Server{
		config:   config,
		database: database,
		logger:   logger,
	}
}

func (s *Server) Run() error {
	listen, err := net.Listen("tcp", s.config.Common.Host)
	if err != nil {
		return err
	}

	s.logger.Infow("Started server", "host", s.config.Common.Host)

	for {
		client, err := listen.Accept()
		if err != nil {
			s.logger.Errorw("Error accept connection", "err", err)
			continue
		}

		log.Println(client)
	}
}
