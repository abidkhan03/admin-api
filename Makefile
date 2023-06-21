build:
	go build -o bin/admin-api cmd/backend/*.go

test:
	go test -count=1 -p 1 ./...

migrate:
	go build -o bin/migrate cmd/dbmigrate/*.go

analyze:
	go build -o bin/analyze cmd/analyze/*.go

gpt-analyze:
	go build -o bin/analyze cmd/gpt-analyze/*.go

gpt-examples:
	go build -o bin/analyze cmd/gpt-examples/*.go
