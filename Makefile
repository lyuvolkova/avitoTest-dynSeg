swagger:
	swagger generate spec -o ./swagger.json -c server
	swagger generate client -f ./swagger.json -t ./tests/client --skip-validation -A service

test:
	curl -o /dev/null http://localhost:8080
	go test ./tests/...
