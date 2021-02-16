package handler

import (
	"context"
	"fmt"
	"net/http"
	"stocksync/pkg/http/contract"
	"stocksync/pkg/http/internal/utils"
	"stocksync/pkg/stockinfo"
	"stocksync/pkg/stockinfo/dto"
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

func (sih *StockInfoHandler) GetStockInfo(resp http.ResponseWriter, req *http.Request) error {
	ctx := context.Background()
	fsymKeys, ok := req.URL.Query()["fsyms"]
	if !ok || len(fsymKeys[0]) < 1 {
		return fmt.Errorf("Url Param 'fsym' is missing")
	}

	tsymKeys, ok := req.URL.Query()["tsyms"]
	if !ok || len(tsymKeys[0]) < 1 {
		return fmt.Errorf("Url Param 'tsym' is missing")
	}

	stockQuery := &dto.StockQuery{
		Fsyms: fsymKeys,
		Tsyms: tsymKeys,
	}

	stockResponses, err := sih.svc.GetStocksFor(ctx, stockQuery)
	if err != nil {
		return fmt.Errorf("error occurred while fetching stock Infos: %v", err)
	}
	sf := &contract.StockFormatter{stockResponses}
	utils.WriteSuccessResponse(resp, http.StatusOK, sf.FormatStockInfoResponse())
	return nil
}
