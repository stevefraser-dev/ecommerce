FROM golang:1.22 as base
WORKDIR /build/
COPY go.* ./
RUN go mod download
COPY . ./

FROM base as cli
WORKDIR /build/cmd/cli/
RUN CGO_ENABLED=0 go build -o cli
ENTRYPOINT ["/build/cmd/cli/cli"]

FROM base AS builder
WORKDIR /build/cmd/web/
RUN CGO_ENABLED=0 go build -o webapp

FROM alpine as web
WORKDIR /
COPY --from=cli /build/cmd/cli/cli /cli
COPY --from=builder /build/cmd/web/webapp /webapp
COPY layout layout
RUN chmod +x /webapp
CMD ["/webapp"]