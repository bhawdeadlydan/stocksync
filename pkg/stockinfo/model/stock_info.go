package model

import (
	"github.com/google/uuid"
	"time"
)

type StockInfo struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt       time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	Fsym            string    `gorm:"column:fsym;primaryKey" json:"fsym"`
	Tsym            string    `gorm:"column:tsym;primaryKey" json:"tsym"`
	DisplayFsym     string    `gorm:"column:display_fsym;" json:"display_fsym"`
	DisplayTsym     string    `gorm:"column:display_tsym;" json:"display_tsym"`
	Change24Hour    string    `gorm:"column:change_24_hour" json:"CHANGE24HOUR"`
	ChangePct24Hour string    `gorm:"column:change_pct_24_hour" json:"CHANGEPCT24HOUR"`
	Open24Hour      string    `gorm:"column:open24_hour" json:"OPEN24HOUR"`
	Volume24Hour    string    `gorm:"column:volume24_hour" json:"VOLUME24HOUR"`
	Volume24Hourto  string    `gorm:"column:volume_24_hour_to" json:"VOLUME24HOURTO"`
	Low24Hour       string    `gorm:"column:low_24_hour" json:"LOW24HOUR"`
	High24Hour      string    `gorm:"column:high_24_hour" json:"HIGH24HOUR"`
	Price           string    `gorm:"column:price" json:"PRICE"`
	Supply          string    `gorm:"column:supply" json:"SUPPLY"`
	MktCap          string    `gorm:"column:mkt_cap" json:"MKTCAP"`
}

func (StockInfo) TableName() string {
	return "stock_info"
}
