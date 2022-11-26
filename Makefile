.PHONY: build clean clean-bin deploy gomodgen print print-prod

build: clean-bin gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/add_default_review default_review/handlers/add/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_all_default_reviews default_review/handlers/get_all/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/remove_default_review default_review/handlers/remove/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/add_review review/handlers/add/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_batch_reviews review/handlers/get_batch/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_user_reviews review/handlers/get_user/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/remove_review review/handlers/remove/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/vote_review review/handlers/vote/main.go

clean:
	rm -rf ./bin ./vendor go.sum

clean-bin:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

print:
	sls print

print-prod:
	sls print -s prod
