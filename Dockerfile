FROM golang:1.18 AS build
WORKDIR /go/src

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o site .

FROM busybox AS runtime
WORKDIR /go/app

COPY --from=build /go/src/site .
COPY --from=build /go/src/public ./public
COPY --from=build /go/src/templates ./templates
EXPOSE 8080/tcp
ENTRYPOINT ["./site"]
