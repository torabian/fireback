#!/bin/bash

# Define variables
DIRNAME="../../artifacts/fireback-server-all/"
PACKAGE_NAME="fireback"
PACKAGE_VERSION="1.2.2"
DESCRIPTION="Fireback ultimate golang framework"
MAINTAINER="Ali Torabi <ali-torabian@outlook.com>"
BIN_PATH_AMD="../../artifacts/fireback-server-all/fireback" # Change this to the actual path of your amd64 binary
BIN_PATH_ARM="../../artifacts/fireback-server-all/fireback"   # Change this to the actual path of your arm binary

get_version() {
    local BIN_PATH=$1
    VERSION=$($BIN_PATH version | head -n 1 )
    echo $VERSION
}

# Create a function to build the package for a specific architecture
build_package() {
    local ARCH=$1
    local BIN_PATH=$2
    local VERSION=$(get_version "go run . version")

    echo $VERSION

    rm -rf ${DIRNAME}${PACKAGE_NAME}-${ARCH}
    # Create the directory structure
    mkdir -p ${DIRNAME}${PACKAGE_NAME}-${ARCH}/{DEBIAN,usr/local/bin}

    # Create the control file
    cat <<EOF > ${DIRNAME}${PACKAGE_NAME}-${ARCH}/DEBIAN/control
Package: ${PACKAGE_NAME}
Version: ${VERSION}
Section: base
Priority: optional
Architecture: ${ARCH}
Essential: no
Installed-Size: $(du -ks ${BIN_PATH} | cut -f1)
Maintainer: ${MAINTAINER}
Description: ${DESCRIPTION}
EOF

    make linux-${ARCH}

    # Copy the binary
    cp ${BIN_PATH} ${DIRNAME}${PACKAGE_NAME}-${ARCH}/usr/local/bin/${PACKAGE_NAME}
    chmod 755 ${DIRNAME}${PACKAGE_NAME}-${ARCH}/usr/local/bin/${PACKAGE_NAME}

    # Build the package
    dpkg-deb --build ${DIRNAME}${PACKAGE_NAME}-${ARCH}

    echo "Package ${DIRNAME}${PACKAGE_NAME}-${ARCH}.deb created successfully."
}

# Build package for amd64
build_package "amd64" ${BIN_PATH_AMD}

# Build package for arm
build_package "arm64" ${BIN_PATH_ARM}
