BINARY := cwd2color

.PHONY: all build clean run

all: build

build:
	go build -o $(BINARY)

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)
