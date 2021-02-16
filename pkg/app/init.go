package app

import (
	"io"
	"log"
	"net/http"
	"os"
	"stocksync/pkg/background/handler"
	"stocksync/pkg/client"
	"stocksync/pkg/config"
	"stocksync/pkg/http/router"
	"stocksync/pkg/http/server"
	"stocksync/pkg/reporters"
	"stocksync/pkg/repository"
	"stocksync/pkg/stockinfo"
	"go.uber.org/zap"
)

func initHTTPServer(configFile string) {
	config := config.NewConfig(configFile)
	logger := initLogger(config)
	rt := initRouter(config, logger)
	bgh := initBackgroundHandler(config, logger)

	server.NewServer(config, logger, rt, bgh).Start()
}

func initRouter(cfg config.Config, logger *zap.Logger) http.Handler {
	stockPriceRepo := initRepository(cfg)
	stockInfoService := initService(stockPriceRepo)

	return router.NewRouter(logger, stockInfoService)
}

func initService(stockInfoRepository repository.StockRepository) stockinfo.Service {
	stockInfoService := stockinfo.NewStockInfoService(stockInfoRepository)

	return stockInfoService
}

func initBackgroundHandler(cfg config.Config, logger *zap.Logger) *handler.StockInfoBackgroundHandler {
	stockPriceRepo := initRepository(cfg)
	stockInfoService := initService(stockPriceRepo)
	stockClient := initClient(cfg.GetClientConfig())

	return handler.NewStockInfoBackgroundHandler(logger, stockInfoService, stockClient, cfg.GetDataRefresherConfig())
}

func initRepository(cfg config.Config) repository.StockRepository {
	dbConfig := cfg.GetDBConfig()
	dbHandler := repository.NewDBHandler(dbConfig)

	db, err := dbHandler.GetDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	return repository.NewStockInfoRepository(db)
}

func initLogger(cfg config.Config) *zap.Logger {
	return reporters.NewLogger(
		cfg.GetLogConfig().GetLevel(),
		getWriters(cfg.GetLogFileConfig())...,
	)
}

func getWriters(cfg config.LogFileConfig) []io.Writer {
	return []io.Writer{
		os.Stdout,
		reporters.NewExternalLogFile(cfg),
	}
}

func initClient(cfg config.ClientConfig) *client.StockClient {
	hc := client.NewHTTPClient(cfg.GetTimeout())

	stockClient := client.NewStockClient(hc, cfg.GetStockClientBaseURL())

	return stockClient
}
