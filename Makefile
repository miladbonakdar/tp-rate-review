.PHONY: build clean clean-bin deploy gomodgen print print-prod

build: clean-bin gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/add_default_review defaultreview/handlers/add/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_default_reviews defaultreview/handlers/getall/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/remove_default_review defaultreview/handlers/remove/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/add_review review/handlers/add/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_batch_reviews review/handlers/getbatch/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_user_reviews review/handlers/getuser/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/remove_review review/handlers/remove/main.go

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
