FROM golang:latest as build-env

# All these steps will be cached
RUN mkdir /random-questions-backend
WORKDIR /random-questions-backend
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/random-questions-backend

FROM scratch
COPY --from=build-env /go/bin/random-questions-backend /go/bin/random-questions-backend
ENTRYPOINT ["/go/bin/random-questions-backend"]