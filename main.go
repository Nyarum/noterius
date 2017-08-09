package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/server"
	"go.uber.org/zap"
)

func main() {
	// Wait to handle all exit functions
	defer func() {
		time.Sleep(3 * time.Second)
	}()

	configPath := flag.String("config", "resource/config.toml", "Path to config")
	flag.Parse()

	logger, _ := zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	configInstance := core.NewConfig()
	if err := configInstance.Load(*configPath); err != nil {
		sugarLogger.Fatalw("Error config load", "err", err)
	}

	databaseInstance := core.NewDatabase()
	if err := databaseInstance.Load(configInstance.Database.Dsn); err != nil {
		sugarLogger.Fatalw("Error database load", "err", err)
	}

	serverInstance := server.NewServer(*configInstance, databaseInstance.DB, sugarLogger)
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-ch:
				cancel()
			}
		}
	}()

	if err := serverInstance.Run(ctx); err != nil {
		sugarLogger.Fatalw("Error server run", "err", err)
	}
}
