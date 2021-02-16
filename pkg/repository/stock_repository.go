package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"stocksync/pkg/stockinfo/dto"
	"stocksync/pkg/stockinfo/model"
	"time"
)

type StockRepository interface {
	CreateOrUpdateStockEntry(ctx context.Context, stockInfo *model.StockInfo) error
	GetStockData(ctx context.Context, stockQuery *dto.StockQuery) ([]model.StockInfo, error)
}

type gormStockRepository struct {
	db *gorm.DB
}

func (gbr *gormStockRepository) CreateOrUpdateStockEntry(ctx context.Context, stockInfo *model.StockInfo) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	db := gbr.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "fsym"}, {Name: "tsym"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"change_24_hour",
				"change_pct_24_hour",
				"open24_hour",
				"volume24_hour",
				"volume_24_hour_to",
				"low_24_hour",
				"high_24_hour",
				"price",
				"supply",
				"mkt_cap",
				"updated_at",
			}),
		}).Create(&stockInfo)
	if db.Error != nil {
		return fmt.Errorf("create stock entry failed. error %w", db.Error)
	}

	return nil
}

func (gbr *gormStockRepository) GetStockData(ctx context.Context, stockQuery *dto.StockQuery) ([]model.StockInfo, error) {
	var res []model.StockInfo
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	db := gbr.db.WithContext(ctx).Where("fsym in ? and tsym in ?", stockQuery.Fsyms, stockQuery.Tsyms).Find(&res)
	if db.Error != nil {
		return nil, fmt.Errorf("get stock data failed: %w", db.Error)
	}

	return res, nil
}

func NewStockInfoRepository(db *gorm.DB) StockRepository {
	return &gormStockRepository{
		db: db,
	}
}
