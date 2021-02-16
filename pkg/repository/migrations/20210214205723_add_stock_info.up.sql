create table if not exists stock_info (
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fsym varchar(100) not null ,
    tsym varchar(100) not null ,
    change_24_hour varchar(100) not null,
    change_pct_24_hour varchar(100) not null,
    open24_hour varchar(100) not null,
    volume24_hour varchar(100) not null,
    volume_24_hour_to varchar(100) not null,
    low_24_hour varchar(100) not null,
    high_24_hour varchar(100) not null,
    price varchar(100) not null,
    supply varchar(100) not null,
    mkt_cap varchar(100) not null,
    PRIMARY KEY (fsym, tsym)
)