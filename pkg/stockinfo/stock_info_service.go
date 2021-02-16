package stockinfo

import (
	"context"
	"fmt"
	"stocksync/pkg/repository"
	"stocksync/pkg/stockinfo/dto"
	"stocksync/pkg/stockinfo/mapper"
	"stocksync/pkg/stockinfo/model"
)

type Service interface {
	CreateOrUpdateStockInfo(ctx context.Context, info *model.StockInfo) error
	GetStocksFor(ctx context.Context, stockQuery *dto.StockQuery) ([]*dto.StockResponse, error)
}

type stockInfoService struct {
	repository repository.StockRepository
}

func (sis *stockInfoService) CreateOrUpdateStockInfo(ctx context.Context, info *model.StockInfo) error {
	err := sis.repository.CreateOrUpdateStockEntry(ctx, info)
	if err != nil {
		return fmt.Errorf("Service.CreateOrUpdateStockInfo failed. Error: %w", err)
	}

	return nil
}

func (sis *stockInfoService) GetStocksFor(ctx context.Context, stockQuery *dto.StockQuery) ([]*dto.StockResponse, error) {
	stockInfos, err := sis.repository.GetStockData(ctx, stockQuery)
	if err != nil {
		return nil, fmt.Errorf("Service.GetStocksFor", err)
	}

	return mapper.GetFormattedResponseFor(stockInfos), nil
}

func NewStockInfoService(repository repository.StockRepository) Service {
	return &stockInfoService{
		repository: repository,
	}
}
