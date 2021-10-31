.PHONY = run
build: 
	go build
clean:
	rm med
run:
	go build && ./med
