package contract

import "stocksync/pkg/stockinfo/dto"

const StockInfoCreationSuccess = "Stock Info created successfully"

type StockFormatter struct {
	StockResponses []*dto.StockResponse
}

func (sf *StockFormatter) FormatStockInfoResponse() interface{} {
	return map[string]interface{}{
		"RAW":     getRawResponse(sf.StockResponses),
		"DISPLAY": getDisplayResponse(sf.StockResponses),
	}
}

func getRawResponse(stockResponses []*dto.StockResponse) interface{} {
	rawResponse := make(map[string]interface{})
	groupedStockResponses := groupByFsym(stockResponses)

	for fsym, groupOfStockResponses := range groupedStockResponses {
		fsymStockResponse := make(map[string]dto.RawStockInfo)
		for _, stockResponse := range groupOfStockResponses {
			fsymStockResponse[stockResponse.ToSymbol] = stockResponse.RawStockInfo
		}
		rawResponse[fsym] = fsymStockResponse
	}

	return rawResponse
}

func getDisplayResponse(stockResponses []*dto.StockResponse) interface{} {
	rawResponse := make(map[string]interface{})
	groupedStockResponses := groupByFsym(stockResponses)

	for fsym, groupOfStockResponses := range groupedStockResponses {
		fsymStockResponse := make(map[string]dto.DisplayStockInfo)
		for _, stockResponse := range groupOfStockResponses {
			fsymStockResponse[stockResponse.ToSymbol] = stockResponse.DisplayStockInfo
		}
		rawResponse[fsym] = fsymStockResponse
	}

	return rawResponse
}

func groupByFsym(stockInfos []*dto.StockResponse) map[string][]*dto.StockResponse {
	groupedStocks := make(map[string][]*dto.StockResponse)

	for _, si := range stockInfos {
		groupedStocks[si.FromSymbol] = append(groupedStocks[si.FromSymbol], si)
	}

	return groupedStocks
}
