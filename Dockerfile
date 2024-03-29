FROM golang:1.20 AS build
WORKDIR /go/src

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -buildvcs=false -a -installsuffix cgo -o site .

FROM busybox AS runtime
WORKDIR /go/app

COPY --from=build /go/src/site .
COPY --from=build /go/src/static ./static
COPY --from=build /go/src/templates ./templates
EXPOSE 8080/tcp
ENTRYPOINT ["./site"]
