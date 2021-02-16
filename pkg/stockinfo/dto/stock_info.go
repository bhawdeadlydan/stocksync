package dto

import (
	"stocksync/pkg/stockinfo/model"
	"strconv"
)

type StockQuery struct {
	Fsyms []string
	Tsyms []string
}

type RawStockInfo struct {
	CHANGE24HOUR    float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR    float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR       float64 `json:"LOW24HOUR"`
	HIGH24HOUR      float64 `json:"HIGH24HOUR"`
	PRICE           float64 `json:"PRICE"`
	LASTUPDATE      string     `json:"LASTUPDATE"`
	SUPPLY          int     `json:"SUPPLY"`
	MKTCAP          float64 `json:"MKTCAP"`
}

type DisplayStockInfo struct {
	CHANGE24HOUR    string `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      string `json:"OPEN24HOUR"`
	VOLUME24HOUR    string `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  string `json:"VOLUME24HOURTO"`
	HIGH24HOUR      string `json:"HIGH24HOUR"`
	PRICE           string `json:"PRICE"`
	FROMSYMBOL      string `json:"FROMSYMBOL"`
	TOSYMBOL        string `json:"TOSYMBOL"`
	LASTUPDATE      string `json:"LASTUPDATE"`
	SUPPLY          string `json:"SUPPLY"`
	MKTCAP          string `json:"MKTCAP"`
}

type StockResponse struct {
	RawStockInfo     RawStockInfo
	DisplayStockInfo DisplayStockInfo
	FromSymbol       string
	ToSymbol         string
}

// mapping and formatting happens here
func NewStockResponse(stockInfo model.StockInfo) *StockResponse {
	return &StockResponse{
		RawStockInfo: RawStockInfo{
			CHANGE24HOUR: parseFloat(stockInfo.Change24Hour),
			CHANGEPCT24HOUR: parseFloat(stockInfo.ChangePct24Hour),
			OPEN24HOUR: parseFloat(stockInfo.Open24Hour),
			VOLUME24HOUR: parseFloat(stockInfo.Volume24Hour),
			VOLUME24HOURTO: parseFloat(stockInfo.Volume24Hourto),
			LOW24HOUR: parseFloat(stockInfo.Low24Hour),
			HIGH24HOUR: parseFloat(stockInfo.High24Hour),
			PRICE: parseFloat(stockInfo.Price),
			LASTUPDATE: "",
			SUPPLY: parseInt(stockInfo.Supply),
			MKTCAP: parseFloat(stockInfo.MktCap),
		},
		DisplayStockInfo: DisplayStockInfo{
			CHANGE24HOUR: stockInfo.Change24Hour,
			CHANGEPCT24HOUR: stockInfo.ChangePct24Hour,
			OPEN24HOUR: stockInfo.Open24Hour,
			VOLUME24HOUR: stockInfo.Volume24Hour,
			VOLUME24HOURTO: stockInfo.Volume24Hourto,
			HIGH24HOUR: stockInfo.Low24Hour,
			PRICE: stockInfo.High24Hour,
			FROMSYMBOL: stockInfo.Fsym,
			TOSYMBOL: stockInfo.Tsym,
			LASTUPDATE: "",
			SUPPLY: stockInfo.Supply,
			MKTCAP: stockInfo.MktCap,
		},
		FromSymbol: stockInfo.Fsym,
		ToSymbol:   stockInfo.Tsym,
	}
}

func parseFloat(element string) float64 {
	floatVal, _ := strconv.ParseFloat(element, 64)
	return floatVal
}
func parseInt(element string) int {
	intVal, _ := strconv.ParseInt(element, 10, 16)
	return int(intVal)
}