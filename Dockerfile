FROM golang:1.18 AS base

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

FROM base AS development

RUN apt-get update && apt-get install --no-install-recommends -y \
  make \
  && \
  apt-get clean && rm -rf /var/lib/apt/lists/*

COPY . .
RUN make out/liveserver
CMD ["./out/liveserver"]

FROM base AS builder

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static AS production

COPY --from=builder /go/bin/app /
CMD ["/app"]
