package contracts

import (
	"fmt"
	"stocksync/pkg/stockinfo/model"
)

type StockPrice struct {
	CHANGE24HOUR    float64
	CHANGEPCT24HOUR float64
	OPEN24HOUR      float64
	VOLUME24HOUR    float64
	VOLUME24HOURTO  float64
	LOW24HOUR       float64
	HIGH24HOUR      float64
	PRICE           float64
	SUPPLY          float64
	MKTCAP          float64
}

func (cdf *StockPrice) ToStockInfo(fsym string, tsym string) *model.StockInfo {
	stockInfo := &model.StockInfo{
		Fsym:            fsym,
		Tsym:            tsym,
		Change24Hour:    fmt.Sprintf("%f", cdf.CHANGE24HOUR),
		ChangePct24Hour: fmt.Sprintf("%f", cdf.CHANGEPCT24HOUR),
		Open24Hour:      fmt.Sprintf("%f", cdf.OPEN24HOUR),
		Volume24Hour:    fmt.Sprintf("%f", cdf.VOLUME24HOUR),
		Volume24Hourto:  fmt.Sprintf("%f", cdf.VOLUME24HOURTO),
		Low24Hour:       fmt.Sprintf("%f", cdf.LOW24HOUR),
		High24Hour:      fmt.Sprintf("%f", cdf.HIGH24HOUR),
		Price:           fmt.Sprintf("%f", cdf.PRICE),
		Supply:          fmt.Sprintf("%f", cdf.SUPPLY),
		MktCap:          fmt.Sprintf("%f", cdf.MKTCAP),
	}
	return stockInfo
}
