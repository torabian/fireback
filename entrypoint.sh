#!/bin/sh

set -e

echo "Running migrations..."
./fireback migration apply

echo "Starting server..."
exec ./fireback start