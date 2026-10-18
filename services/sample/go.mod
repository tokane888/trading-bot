// TODO: search and replace the module name below across the git repository
module github.com/tokane888/gtrading-bot/services/sample

go 1.27.2

require (
	github.com/joho/godotenv v1.5.1
	github.com/tokane888/trading-bot/pkg/logger v0.0.0
)

replace github.com/tokane888/trading-bot/pkg/logger => ../../pkg/logger
