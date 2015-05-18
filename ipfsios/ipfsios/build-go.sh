#!/bin/bash
echo "Go Build Started .."

set -e

# Important - code the GOROOT here to point to goios' root
export GOROOT=$SRCROOT/../goios.ios64-new/

# Full path to the go-1.4 compiler from goios
GO=$GOROOT/bin/go

# Some tools which Go depends on are in /usr/local/bin - hg, git, etc.
export PATH=$PATH:/usr/local/bin/

# The root of the Go sources in our project
GOSOURCES=$SRCROOT/ipfsios/gosources

# Where Go objects will be put
GG_OBJ=$BUILT_PRODUCTS_DIR/go-obj

# Where cgo objects will be put
GG_CGO_OBJ=$BUILT_PRODUCTS_DIR/cgo-obj

# The destination of the library of our Go code for linking with
ARCHIVE=$BUILT_PRODUCTS_DIR/go-output.a

# Set up GOPATH for a local build/install
export GOPATH=$BUILT_PRODUCTS_DIR/go-src

# Set up GOBIN for any results
export GOBIN=$GOPATH/bin

echo "GO=$GO"
echo "GG_OBJ=$GG_OBJ"
echo "GG_CGO_OBJ=$GG_CGO_OBJ"
echo "ARCHIVE=$ARCHIVE"
echo "GOSOURCES=$GOSOURCES"
echo "GOPATH=$GOPATH"
echo "PATH=$PATH"

mkdir -p $GG_OBJ $GG_CGO_OBJ $GOPATH $GOBIN

cd $GOSOURCES

echo "Generating header _cgo_export.h"
GOARM=7 CGO_ENABLED=1 GOARCH=arm64 $GO tool cgo -objdir $GG_CGO_OBJ main.go
cp $GG_CGO_OBJ/_cgo_export.h $SRCROOT/ipfsios

echo "Getting dependencies"
GOARM=7 CGO_ENABLED=1 GOARCH=arm64 $GO get -v .

echo "Building Go Code"
GOARM=7 CGO_ENABLED=1 GOARCH=arm64 $GO build -i -ldflags '-tmpdir '$GG_OBJ' -linkmode external'

ar rcs $ARCHIVE $GG_OBJ/*.o
echo "Produced Go code as library in $ARCHIVE"
#ar rcs $SRCROOT/go-output.a $GG_OBJ/*.o

cp $ARCHIVE $SRCROOT/ipfsios/
echo "Produced Go code as library in $SRCROOT"

echo "Go Build Completed."