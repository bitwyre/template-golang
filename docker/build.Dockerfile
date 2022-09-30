FROM golang:1.17-alpine AS build

RUN mkdir -p /app/teras
WORKDIR /app/teras
COPY . .

RUN CGO_ENABLED=0 go build -o /app/teras/app /app/teras/main.go

###
FROM gcr.io/distroless/base-debian11

COPY --from=build /app/teras/app .

# Copy env template
COPY --from=build /app/teras/docker/.template.env.dev ./.env

CMD ["/app", "serve",  "--env", "dev"]