module stocksync

go 1.14

require (
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/google/uuid v1.1.2
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.9.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/otel v0.16.0
	go.uber.org/zap v1.16.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)
