FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/backend ./backend/cmd/main.go

FROM oven/bun:1-alpine AS frontend-builder

WORKDIR /app/frontend

COPY frontend/package.json frontend/bun.lock ./
RUN bun install --frozen-lockfile

COPY frontend/ .
RUN bun run build

FROM gcr.io/distroless/cc-debian12

COPY --from=builder /bin/backend .
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist

EXPOSE 8911

CMD ["./backend"]
