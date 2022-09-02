package app

import (
	"github.com/coreyvan/backend-takehome/internal/transport"
	"go.uber.org/zap"
)

func Run(log *zap.Logger) error {
	srv := transport.NewHTTP(log, 3000)

	return srv.Listen()
}
