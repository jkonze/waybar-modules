build:
	test -n "$(OWM_API_KEY)"
	test -n "$(OWM_LOCATION)"
	go build -ldflags "-X main.version=$(OWM_API_KEY) -X main.location=$(OWM_LOCATION)"