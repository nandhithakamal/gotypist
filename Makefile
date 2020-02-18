run: build 
	./gotypist

build:
	@echo Compiling gotypist.
	go build -o gotypist ./src

test:
	@echo Testing gotypist.
	go test -v ./src

clean:
	@echo Deleting config.
	rm -r ~/.config/gotypist
