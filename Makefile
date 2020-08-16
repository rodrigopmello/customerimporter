NAME=ci


all:	
	go build -o $(NAME)
	./$(NAME) $(CSV)

build:
	go build -o $(NAME)


tests:
	go test ./customerimporter -v


benchmark:
	go test -bench=. ./customerimporter




