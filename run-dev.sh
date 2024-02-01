#!/bin/bash

# Development build & run script for ucsd-reservation-maker
# note: this build assumes the .dev directory has been set up with the 
# needed resources and that currently running on MacOs x86_64 platform

DEV_DIR=.dev
PROGRAM_NAME=make-reservation

echo "dev: building for MacOs x64..."
GOOS=darwin GOARCH=amd64 go build -o $DEV_DIR/$PROGRAM_NAME

echo "dev: running program..."
cd $DEV_DIR; ./$PROGRAM_NAME $1
