#! /bin/sh

ARGS=$1

if [ "$ARGS" = "hot-serve" ]; then
	air
elif [ "$ARGS" == "serve" ]; then
    go run main.go serve - $@
elif [ "$ARGS" == "seeder" ]; then
    go run main.go - $@
elif [ "$ARGS" == "build" ]; then
    go build -o build/main main.go
elif [ "$ARGS" = "setup" ]; then
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
else
    echo "Command is not found"
fi