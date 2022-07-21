# This is a multi build Dockerfile, it builds the ui and the api and combines them into a minimal image

# Step 1: Build the UI
FROM node:16 AS build-ui

WORKDIR /build

# Install pnpm packet manager
RUN npm install -g pnpm

# Prefetch to store
COPY ./pfspell-ui/pnpm-lock.yaml /build/
RUN pnpm fetch

# Copy rest and install
COPY ./pfspell-ui/ ./
RUN pnpm install -r --offline --shamefully-hoist

# Generate the static files
RUN NODE_ENV=production pnpm run generate -o


# Step 2: Build the api 
FROM golang:1.18-alpine AS build-api

WORKDIR /api

# Download the dependencies
COPY ./pfspell-api/go.mod ./
COPY ./pfspell-api/ ./
RUN go mod download

# Copy the rest and build
COPY ./pfspell-api/ ./
RUN go build -o /pfspell-api ./cmd/server/main.go


# Last step, combine the ui and the api builds
FROM alpine:latest

WORKDIR /app

# Copy from build images
COPY --from=build-api /pfspell-api ./pfspell-api
COPY --from=build-ui /build/.output/public/ ./public/

ENTRYPOINT ["/app/pfspell-api"]

