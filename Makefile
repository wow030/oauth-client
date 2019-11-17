run:
	DEDIRECT_URL=https://localhost:8080/redirect
	go run main.go

test:
	go test -count=1 -p 1 -v ./service/...