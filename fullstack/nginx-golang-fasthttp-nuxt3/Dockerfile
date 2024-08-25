# Stage 1: Build Golang backend
FROM golang:1.21 AS backend-builder
WORKDIR /app/backend
COPY backend .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Build Nuxt frontends
FROM oven/bun:1 AS frontend-builder
WORKDIR /app/frontend

# Build first-frontend
COPY frontend/first-frontend/package.json frontend/first-frontend/bun.lockb* ./first-frontend/

WORKDIR /app/frontend/first-frontend
RUN bun install
COPY frontend/first-frontend .
RUN bun run build

FROM oven/bun:1 AS frontend2-builder
WORKDIR /app/frontend

# Build second-frontend
COPY frontend/second-frontend/package.json frontend/second-frontend/bun.lockb* ./second-frontend/

WORKDIR /app/frontend/second-frontend
RUN bun install
COPY frontend/second-frontend .
RUN bun run build

# Stage 3: Final stage
FROM node:18-alpine

# Install Nginx
RUN apk add --no-cache nginx

# Copy backend binary
COPY --from=backend-builder /app/backend/main /usr/local/bin/backend

# Copy frontend builds
COPY --from=frontend-builder /app/frontend/first-frontend/.output/server /app/frontend/first-frontend/server
COPY --from=frontend-builder /app/frontend/first-frontend/.output/public /app/frontend/first-frontend/public

# Copy second builds
COPY --from=frontend2-builder /app/frontend/second-frontend/.output/server /app/frontend/second-frontend/server
COPY --from=frontend2-builder /app/frontend/second-frontend/.output/public /app/frontend/second-frontend/public

# Copy Nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy start script
COPY start.sh /start.sh
RUN chmod +x /start.sh

EXPOSE 80
# EXPOSE 80 8080 3000

CMD ["/start.sh"]