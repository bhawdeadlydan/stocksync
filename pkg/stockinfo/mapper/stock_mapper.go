package mapper

import (
	"stocksync/pkg/stockinfo/dto"
	"stocksync/pkg/stockinfo/model"
)

func GetFormattedResponseFor(stockInfos []model.StockInfo) []*dto.StockResponse {
	stockResponses := make([]*dto.StockResponse, 0)
	for _, stockInfo := range stockInfos {
		stockResponses = append(stockResponses, dto.NewStockResponse(stockInfo))
	}
	return stockResponses
}
