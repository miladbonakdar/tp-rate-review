.PHONY: build clean clean-bin deploy gomodgen print print-prod

build: clean-bin gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/add_default_review handlers/default_review/add.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_all_default_reviews handlers/default_review/get_all.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/remove_default_review handlers/default_review/remove.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/add_review handlers/review/add.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_batch_reviews handlers/review/get_batch.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_user_reviews handlers/review/get_user.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/remove_review handlers/review/remove.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/vote_review handlers/review/vote.go

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
