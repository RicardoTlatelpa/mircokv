FROM golang:1.24.3 as builder

WORKDIR /app

# Copy source files

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o server ./cmd/server

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/server /server

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]