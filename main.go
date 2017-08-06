package main

import (
	"flag"

	"go.uber.org/zap/zapcore"

	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/server"
	"go.uber.org/zap"
)

func main() {
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
	if err := serverInstance.Run(); err != nil {
		sugarLogger.Fatalw("Error server run", "err", err)
	}
}
