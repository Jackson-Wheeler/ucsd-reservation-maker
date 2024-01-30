#!/bin/bash

# Development build & run script for ucsd-reservation-maker:
# note: this build assumes MacOs x64 (Intel) platform

# Check if an argument was provided
if [ $# -eq 0 ]; then
    echo "Usage ./dev-build.sh <config file name>"
    exit 1
fi

configFileName=$1

DEV_DIR=dev
CONFIG_DIR=config
PROGRAM_NAME=ucsd-reservation-maker

echo "building for MacOs x64..."
GOOS=darwin GOARCH=amd64 go build -o $DEV_DIR/$PROGRAM_NAME

echo "running program..."
cd $DEV_DIR; ./$PROGRAM_NAME $CONFIG_DIR/$configFileName
