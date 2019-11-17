run:
	REDIRECT_URL=https://localhost:8080/redirect \
	TOKEN_URL=https://localhost:8081/token \
	go run .

test:
	go test -count=1 -p 1 -v ./service/...