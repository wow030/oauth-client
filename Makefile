run:
	REDIRECT_URL=https://localhost:8000/redirect \
	TOKEN_URL=https://localhost:8001/token \
	go run .

test:
	go test -count=1 -p 1 -v ./service/...