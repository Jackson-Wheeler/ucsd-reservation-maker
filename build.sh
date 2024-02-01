#!/bin/bash

BUILD_DIR=dist

PLATFORMS=(linux64 mac-arm64 mac-x64 win64)
GOOS=(linux darwin darwin windows)
ARCH=(amd64 arm64 amd64 amd64)

PROJECT_NAME=UCSD-Reservation-Maker

BINARY_NAME=make-reservation

DRIVERS_DIR=drivers
DRIVER_DIR=driver

TEMPLATES_DIR=templates

# start new builds directory
rm -rf $BUILD_DIR
mkdir $BUILD_DIR

# for each platform: create a directory, build the program, copy over correct driver, copy over templates dir contents, create zip file & store in 
for ((i=0; i<${#PLATFORMS[@]}; i++)); do
  PLATFORM=${PLATFORMS[$i]}
  # setup
  echo -e "\n-- $PLATFORM --"
  mkdir $BUILD_DIR/$PLATFORM
  mkdir $BUILD_DIR/$PLATFORM/$PROJECT_NAME
  mkdir $BUILD_DIR/$PLATFORM/$PROJECT_NAME/$DRIVER_DIR

  # build binary
  echo "building binary..."
  GOOS=${GOOS[$i]} GOARCH=${ARCH[$i]} go build -o $BUILD_DIR/$PLATFORM/$PROJECT_NAME/$BINARY_NAME

  # copy over corresponding chromedriver
  echo "copying over $PLATFORM chromedriver..."
  cp -r $DRIVERS_DIR/$PLATFORM/* $BUILD_DIR/$PLATFORM/$PROJECT_NAME/$DRIVER_DIR/

  # copy over contents of templates dir
  echo "copying over template files..."
  cp -r $TEMPLATES_DIR/* $BUILD_DIR/$PLATFORM/$PROJECT_NAME/

  # create distribution zip file & remove uncompressed project directory
  echo "creating zip file..."
  cd $BUILD_DIR/$PLATFORM
  zip -r -q $PROJECT_NAME.zip $PROJECT_NAME
  rm -rf $PROJECT_NAME
  cd ../..

done
