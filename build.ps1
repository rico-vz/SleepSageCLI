# Create the build folder if it doesn't exist
if (-not (Test-Path -Path .\build)) {
    New-Item -ItemType Directory -Path .\build
}

# Specify target platforms
$platforms = @(
    "windows/386",
    "windows/amd64",
    "linux/386",
    "linux/amd64",
    "linux/arm",
    "linux/arm64",
    "darwin/amd64",
    "darwin/arm64"
)

# Loop through platforms and build
foreach ($platform in $platforms) {
    Write-Host "Building for platform: $platform"

    $os, $arch = $platform -split '/'
    $outputName = "build/sleepsage_${os}_${arch}"

    $env:GOOS = $os
    $env:GOARCH = $arch

    if ($os -eq "windows") {
        $outputName += ".exe"
    }

    go build -o $outputName sleepsage.go
}

# Clear environment variables
$env:GOOS = ""
$env:GOARCH = ""
