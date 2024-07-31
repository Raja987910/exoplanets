FROM golang:alpine

ENV GO111MODULE=on
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o exoplanets ./cmd/server

# the application is going to listen on by default.
EXPOSE 8080

# Run
CMD ["/exoplanets"]