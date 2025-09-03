#!/bin/bash

echo "Stopping and removing containers, networks, volumes..."
docker compose down -v

echo "Rebuilding images without cache..."
docker compose build --no-cache

echo "Starting containers..."
docker compose up
