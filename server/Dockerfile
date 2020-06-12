# STAGE 1 - BUILD
FROM golang:1.14-alpine AS build

ENV GO114MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

# STAGE 2 - DEPLOY
FROM scratch

COPY --from=build /dist/main /

ENTRYPOINT ["./main"]