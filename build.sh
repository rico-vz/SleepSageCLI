#!/bin/bash

# Create the build folder if it doesn't exist
if [ ! -d "./build" ]; then
    mkdir -p "./build"
fi

# Specify target platforms
platforms=(
    "windows/386"
    "windows/amd64"
    "linux/386"
    "linux/amd64"
    "linux/arm"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
)

# Loop through platforms and build
for platform in "${platforms[@]}"; do
    echo "Building for platform: $platform"

    IFS='/' read -r os arch <<< "$platform"
    outputName="build/sleepsage_${os}_${arch}"

    if [ "$os" == "windows" ]; then
        outputName+=".exe"
    fi

    GOOS=$os GOARCH=$arch go build -o "$outputName" sleepsage.go
done

# Clear environment variables
export GOOS=""
export GOARCH=""
