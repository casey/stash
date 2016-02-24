serve:
	goapp serve

deploy:
	appcfg.py --oauth2 update .

format:
	gofmt -w *.go

lint:
	go vet
	#@golint | sed "/don't use underscores/d"
  golint

test:
	./test http://localhost:8080

test-deployed:
	./test http://rodarmor-stash.appspot.com
