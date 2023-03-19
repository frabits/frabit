# DO NOT run docker build against this file directly. Instead using ./build_docker.sh as that
# one sets the various ARG used in the Dockerfile

# After build

# $ docker run --init --rm --name bytebase --publish 8080:8080 --volume ~/.bytebase/data:/var/opt/bytebase bytebase/bytebase

FROM node:18 as frontend

ARG RELEASE="release"

RUN npm i -g pnpm

WORKDIR /frontend-build

# Install build dependency (e.g. vite)
COPY ./frontend/package.json ./frontend/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

COPY ./frontend/ .
# Copy the SQL review config files to the frontend
COPY ./backend/plugin/advisor/config/ ./src/types

# Build frontend
RUN pnpm "${RELEASE}-docker"

FROM golang:1.20.1 as backend

WORKDIR /backend-build

COPY . .

# Copy frontend asset
COPY --from=frontend /frontend-build/dist ./server/dist

RUN GOOS=linux GOARCH=amd64 go build -o frabit-server ./cmd/frabit-server/main.go

# Use debian because mysql requires glibc.
FROM debian:bullseye-slim as monolithic

# Create user "frabit" for running server.
RUN addgroup --gid 113 --system frabit && adduser --uid 113 --system frabit && adduser frabit frabit

COPY ./scripts /usr/local/bin/

COPY --from=backend /backend-build/frabit-server /usr/local/bin/

CMD ["--port", "80", "--data", "/var/opt/bytebase"]

ENTRYPOINT ["frabit-server"]
