IMG=openwhisk/action-golang-v1.11:latest

chess.go.zip: main/main.go main/index.go
	zip -r - main |\
	docker run -i $(IMG) -compile main >$@

main/index.go: web/index.html
	echo "package main" >$@
	echo 'var indexHTML = `<!DOCTYPE html>' >>$@
	cat web/index.html | sed 's/ + "api\/default\/chess"//' >> $@
	echo '`;' >>$@

deploy: chess.go.zip
	nim action update chess chess.go.zip --docker $(IMG) --web true
	nim action get chess --url

clean:
	-rm chess chess.go.zip
