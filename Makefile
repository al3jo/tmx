NAME        := tmx
VERSION     := 0.0.1
SOURCES 	:= $(shell find . -type f -name *.go)

.PHONY: all clean test

all: $(NAME)

clean:
	@- $(RM) -rf $(NAME)

test: $(NAME)
	go test -cover -v ./...


$(NAME): $(SOURCES)
	GOARCH=amd64 GOOS=linux go build -o $(NAME) -ldflags="-X 'main.Version=v$(VERSION)'"
