FROM golang:1.17-alpine AS build

WORKDIR /app/bitwyre
COPY . .

RUN CGO_ENABLED=0 go build -o /app/bitwyre/app /app/bitwyre/main.go

###
FROM gcr.io/distroless/base-debian11

ENV GIN_MODE=release
COPY --from=build /app/bitwyre/app .

# Copy env template
COPY --from=build /app/bitwyre/docker/.template.env.dev ./.env

EXPOSE 3000

CMD ["/app", "serve",  "--env", "dev"]