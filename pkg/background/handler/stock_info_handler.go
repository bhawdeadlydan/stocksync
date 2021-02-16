package handler

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"stocksync/pkg/client"
	"stocksync/pkg/config"
	"stocksync/pkg/stockinfo"
)

type StockInfoBackgroundHandler struct {
	lgr    *zap.Logger
	svc    stockinfo.Service
	client *client.StockClient
	drc    config.DataRefresherConfig
}

func NewStockInfoBackgroundHandler(lgr *zap.Logger, svc stockinfo.Service, client *client.StockClient, drc config.DataRefresherConfig) *StockInfoBackgroundHandler {
	return &StockInfoBackgroundHandler{
		lgr:    lgr,
		svc:    svc,
		client: client,
		drc:    drc,
	}
}

func (sih *StockInfoBackgroundHandler) UpdateStockInfo() error {
	ctx := context.Background()
	sih.lgr.Sugar().Infof("Updating stock data..")

	for _, fsym := range sih.drc.GetFsyms() {
		for _, tsym := range sih.drc.GetTsyms() {
			sih.lgr.Sugar().Infof("Updating stock data for fsym: %s, tsym: %s", fsym, tsym)

			stockInfo, err := sih.client.GetPriceData(ctx, fsym, tsym)
			if err != nil {
				return fmt.Errorf("StockInfoBackgroundHandler.GetStockPrice . error %v", err)
			}
			sih.lgr.Sugar().Infof("Received stock Info for fsym: %s, tsym: %s. StockInfo: %v", fsym, tsym, stockInfo)

			err = sih.svc.CreateOrUpdateStockInfo(ctx, stockInfo)
			if err != nil {
				return fmt.Errorf("StockInfoBackgroundHandler.CreateStockInfo . error %v", err)
			}
		}
	}

	sih.lgr.Debug("msg", zap.String("eventCode", "STOCK_UPDATED"))
	return nil
}
