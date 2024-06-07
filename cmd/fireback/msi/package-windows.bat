# Download the wix from here:
# https://github.com/wixtoolset/wix3/releases/download/wix3141rtm/wix314-binaries.zip

# Build fireback.exe first for arm64

# make arm64 windows
GOARCH=arm64 GOOS=windows go build -ldflags "-s -w" -o ../../..\artifacts/$(project)-server-all/$(project).exe .
..\..\..\thirdparty\wix314-binaries\candle.exe Product.wxs
..\..\..\thirdparty\wix314-binaries\light.exe -out fireback_win_arm64_installer.msi .\Product.wixobj

# make amd64 windows
# Build fireback.exe first for amd64
..\..\..\thirdparty\wix314-binaries\candle.exe Product.wxs
..\..\..\thirdparty\wix314-binaries\light.exe -out fireback_win_amd64_installer.msi .\Product.wixobj
