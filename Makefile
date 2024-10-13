local:
	go build -o bin/verse ./src

dist:
	docker build -t jyking99/verse:built .

release:
	make dist
	docker tag jyking99/verse:built jyking99/verse:release

