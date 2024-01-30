#!/bin/bash

# Build script for ucsd-reservation-maker: 
# - builds a Go binary specific to each platform, placing each in the proper directory within the builds/ directory. 
# - copies over current config templates and instructions to each platform's directory.
# - creates zip file for each platform's directory, placing them in builds/dist/

BUILD_DIR=builds

LINUX_DIR=linux-64
MAC_ARM_DIR=mac-arm
MAC_64_DIR=mac-x64
WIN_64_DIR=win-64

PLATFORM_DIRS=($LINUX_DIR $MAC_ARM_DIR $MAC_64_DIR $WIN_64_DIR)

# -- Build to each platform --
PROGRAM_NAME=ucsd-reservation-maker

echo "-- Building Program --"
# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/$LINUX_DIR/$PROGRAM_NAME

# Build for macOS ARM
echo "Building for macOS ARM..."
GOOS=darwin GOARCH=arm64 go build -o $BUILD_DIR/$MAC_ARM_DIR/$PROGRAM_NAME

# Build for macOS x64 (Intel)
echo "Building for macOS x64..."
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/$MAC_64_DIR/$PROGRAM_NAME

# Build for Windows x64
echo "Building for Windows x64..."
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/$WIN_64_DIR/$PROGRAM_NAME


# -- Copy over files: config templates, .env template, instructions --
TEMPLATES_DIR=templates

CONFIG_TEMPLATES_DIR=$TEMPLATES_DIR/config
DEST_CONFIG_DIR=config

ENV_FILE=$TEMPLATES_DIR/.env

INSTRUCTIONS=README.md
DEST_INSTRUCTIONS=INSTRUCTIONS.md

echo -e "\n-- Copying Files --"
for DEST in ${PLATFORM_DIRS[@]}; do
  echo "Copying files to $DEST..."

  # Copy over config templates
  rm -rf $BUILD_DIR/$DEST/$DEST_CONFIG_DIR # clear out old config
  cp -r $CONFIG_TEMPLATES_DIR $BUILD_DIR/$DEST/

  # Copy over .env file
  cp $ENV_FILE $BUILD_DIR/$DEST/

  # Copy over new instructions
  cp $INSTRUCTIONS $BUILD_DIR/$DEST/
  mv $BUILD_DIR/$DEST/$INSTRUCTIONS $BUILD_DIR/$DEST/$DEST_INSTRUCTIONS # rename
done

# -- Create zip files --
cd $BUILD_DIR
DIST_DIR=dist

echo -e "\n-- Creating Distributions --"
# Clear out dist directory
rm -rf $DIST_DIR/*

# Create zip files for each platform
for DEST in ${PLATFORM_DIRS[@]}; do
  echo "Creating zip file for $DEST..."
  zip -q -r $DEST.zip $DEST
  mv $DEST.zip $DIST_DIR/
done


