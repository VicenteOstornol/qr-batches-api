FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

ARG GO_CMD
ENV GO_CMD=${GO_CMD}

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${GO_CMD} github.com/VicenteOstornol/lotesapi/cmd/${GO_CMD}


FROM golang:1.21-alpine
ARG GO_CMD
ENV GO_CMD=${GO_CMD}

WORKDIR /usr/local/bin
COPY --from=builder /app/internal/repository/postgres/migrations /usr/local/bin/internal/repository/postgres/migrations
COPY --from=builder /app/bin/${GO_CMD} /usr/local/bin/${GO_CMD}
CMD ${GO_CMD}
