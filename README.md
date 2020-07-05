# rest-coffee


# migration
cd migrations
goose postgres "user=postgres password=postgres host=database sslmode=disable database=coffee" up

# running
go run cmd/main.go
