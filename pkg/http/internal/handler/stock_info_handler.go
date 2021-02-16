package handler

import (
	"context"
	"fmt"
	"net/http"
	"stocksync/pkg/http/contract"
	"stocksync/pkg/http/internal/utils"
	"stocksync/pkg/stockinfo"
	"stocksync/pkg/stockinfo/model"

	"go.uber.org/zap"
)

type StockInfoHandler struct {
	lgr *zap.Logger
	svc stockinfo.Service
}

func NewStockInfoHandler(lgr *zap.Logger, svc stockinfo.Service) *StockInfoHandler {
	return &StockInfoHandler{
		lgr: lgr,
		svc: svc,
	}
}

func (sih *StockInfoHandler) CreateStockInfo(resp http.ResponseWriter, req *http.Request) error {
	ctx := context.Background()
	var si model.StockInfo
	err := utils.ParseRequest(req, &si)
	if err != nil {
		return err
	}

	err = sih.svc.CreateOrUpdateStockInfo(ctx, &si)
	if err != nil {
		return fmt.Errorf("StockInfoHandler.CreateStockInfo . error %v", err)
	}

	sih.lgr.Debug("msg", zap.String("eventCode", utils.StockInfoAdded))
	utils.WriteSuccessResponse(resp, http.StatusCreated, contract.StockInfoCreationSuccess)
	return nil
}
