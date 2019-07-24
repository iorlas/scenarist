FROM golang AS build-env
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM alpine
WORKDIR /app
COPY --from=build-env /src/app .
ENTRYPOINT ./app
