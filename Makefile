.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/add_default_review handlers/default_review/add/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_all_default_reviews handlers/default_review/get_all/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/remove_default_review handlers/default_review/remove/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/add_review handlers/review/add/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_batch_reviews handlers/review/get_batch/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_user_reviews handlers/review/get_user/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/remove_review handlers/review/remove/main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
