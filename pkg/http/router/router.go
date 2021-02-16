package router

import (
	"net/http"
	"stocksync/pkg/http/internal/handler"
	"stocksync/pkg/http/internal/middleware"
	stockinfo "stocksync/pkg/stockinfo"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	stockInfoPath = "/stock"
)

func NewRouter(lgr *zap.Logger, stockInfoService stockinfo.Service) http.Handler {
	router := mux.NewRouter()
	router.Use(handlers.RecoveryHandler())

	sih := handler.NewStockInfoHandler(lgr, stockInfoService)

	router.HandleFunc(stockInfoPath, withMiddlewares(lgr, middleware.WithErrorHandler(lgr, sih.CreateStockInfo))).Methods(http.MethodPost)

	return router
}

func withMiddlewares(lgr *zap.Logger, hnd http.HandlerFunc) http.HandlerFunc {
	return middleware.WithSecurityHeaders(middleware.WithReqResLog(lgr, hnd))
}
