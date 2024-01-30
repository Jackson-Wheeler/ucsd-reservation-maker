#!/bin/bash

# Build script for ucsd-reservation-maker: 
# - builds a Go binary specific to each platform, placing each in the proper directory within the builds/ directory. 
# - copies over current config templates and instructions to each platform's directory.
# - creates zip file for each platform's directory, placing them in builds/dist/

BUILD_DIR=builds

LINUX_DIR=$BUILD_DIR/linux-64
MAC_ARM_DIR=$BUILD_DIR/mac-arm
MAC_64_DIR=$BUILD_DIR/mac-x64
WIN_64_DIR=$BUILD_DIR/win-64

DESTINATIONS=($LINUX_DIR $MAC_ARM_DIR $MAC_64_DIR $WIN_64_DIR)

# -- Build to each platform --
PROGRAM_NAME=ucsd-reservation-maker

echo "-- Building Program --"
# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o $LINUX_DIR/$PROGRAM_NAME

# Build for macOS ARM
echo "Building for macOS ARM..."
GOOS=darwin GOARCH=arm64 go build -o $MAC_ARM_DIR/$PROGRAM_NAME

# Build for macOS x64 (Intel)
echo "Building for macOS x64..."
GOOS=darwin GOARCH=amd64 go build -o $MAC_64_DIR/$PROGRAM_NAME

# Build for Windows x64
echo "Building for Windows x64..."
GOOS=windows GOARCH=amd64 go build -o $WIN_64_DIR/$PROGRAM_NAME


# -- Copy over files: config templates, .env template, instructions --
TEMPLATES_DIR=templates

CONFIG_TEMPLATES_DIR=$TEMPLATES_DIR/config
DEST_CONFIG_DIR=config

ENV_FILE=$TEMPLATES_DIR/.env

INSTRUCTIONS=README.md
DEST_INSTRUCTIONS=INSTRUCTIONS.md

echo -e "\n-- Copying Files --"
for DEST in ${DESTINATIONS[@]}; do
  echo "Copying files to $(basename $DEST)..."

  # Clear out config directory & copy over config templates
  rm -rf $DEST/$DEST_CONFIG_DIR/*
  cp -r $CONFIG_TEMPLATES_DIR $DEST/$DEST_CONFIG_DIR/

  # Copy over .env file
  cp $ENV_FILE $DEST/

  # Copy over new instructions
  cp $INSTRUCTIONS $DEST/
  mv $DEST/$INSTRUCTIONS $DEST/$DEST_INSTRUCTIONS # rename
done

# -- Create zip files --
DIST_DIR=$BUILD_DIR/dist

echo -e "\n-- Creating Distributions --"
# Clear out dist directory
echo "Clearing out $DIST_DIR directory..."
rm -rf $DIST_DIR/*

# Create zip files for each platform
for DEST in ${DESTINATIONS[@]}; do
  DEST_NAME=$(basename $DEST)
  echo "Creating zip file for $DEST_NAME..."
  zip -q -r $DIST_DIR/$DEST_NAME.zip $DEST
done


