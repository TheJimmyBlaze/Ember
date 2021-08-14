
#Calculate build version
$VERSION_PATH = "build_version.semver"
$BuildVersion = Get-Content -Path $VERSION_PATH
Write-Host "Build Version: $BuildVersion"

#Calculate build time
$BuildTime = (Get-Date).DateTime
Write-Host "Build Time: $BuildTime"

#Build Archive used for getting build hash
$COMPRESS_PATH = "build-archive.temp.zip"
$Compress = @{
    Path="."
    CompressionLevel = "Fastest"
    DestinationPath = $COMPRESS_PATH
}
Compress-Archive @Compress
#Calculate build hash
$BuildFileHash = Get-FileHash $COMPRESS_PATH
$BuildHash = $BuildFileHash.Hash
Remove-Item $COMPRESS_PATH
Write-Host "Build Hash: $BuildHash"

#go build
Write-Host "Building..."
go build -ldflags="-X 'github.com/thejimmyblaze/ember/version.BuildVersion=$BuildVersion' -X 'github.com/thejimmyblaze/ember/version.BuildTime=$BuildTime' -X 'github.com/thejimmyblaze/ember/version.BuildHash=$BuildHash'" -o ember-ca.exe

Write-Host "Done"