# The main Build image to build all our binaries
FROM golang:1.21.3-alpine3.18 as build

WORKDIR /

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY ./pkg ./pkg
COPY ./app ./app
COPY ./cmd/api ./cmd/api


# Just build API
FROM build as build-api
# Generate swagger docs
WORKDIR /cmd/api
RUN swag init -g main.go -o /docs -d ./,/pkg,/app/api
# build binary
WORKDIR /
RUN go build -o chocApi /cmd/api/main.go


# Create API release image
FROM alpine:3.18 as api
# Copy our static executable
COPY --from=build-api /chocApi /chocApi

EXPOSE 8080
CMD [ "/chocApi" ]
