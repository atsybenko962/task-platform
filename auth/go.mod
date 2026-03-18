module auth

go 1.25.7

replace github.com/fin/tools => ../tools

require (
	github.com/fin/tools v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
	go.uber.org/zap v1.27.1
)

require (
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
)
