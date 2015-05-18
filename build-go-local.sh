
#!/bin/bash
echo "Go Build Started .."

export GO=/Users/j/Documents/ipfs/ipfs-on-ios/ipfsios/../goios.ios64-new//bin/go
export GG_OBJ=/Users/j/Library/Developer/Xcode/DerivedData/ipfsios-aklfckqgczeogudriwvjkojnmrsg/Build/Products/Debug-iphoneos/go-obj
export GG_CGO_OBJ=/Users/j/Library/Developer/Xcode/DerivedData/ipfsios-aklfckqgczeogudriwvjkojnmrsg/Build/Products/Debug-iphoneos/cgo-obj
export ARCHIVE=/Users/j/Library/Developer/Xcode/DerivedData/ipfsios-aklfckqgczeogudriwvjkojnmrsg/Build/Products/Debug-iphoneos/go-output.a
export GOSOURCES=/Users/j/Documents/ipfs/ipfs-on-ios/ipfsios/ipfsios/gosources
export GOPATH=/Users/j/Library/Developer/Xcode/DerivedData/ipfsios-aklfckqgczeogudriwvjkojnmrsg/Build/Products/Debug-iphoneos/go-src
export PATH=$PATH:/usr/local/bin

#echo "Generating header _cgo_export.h"
#GOARM=7 CGO_ENABLED=1 GOARCH=arm64 $GO tool cgo -objdir $GG_CGO_OBJ main.go
#cp $GG_CGO_OBJ/_cgo_export.h $SRCROOT/ipfsios

#echo "Getting dependencies"
#GOARM=7 CGO_ENABLED=1 GOARCH=arm64 $GO get -v .

#echo "Building Go Code"
GOARM=7 CGO_ENABLED=1 GOARCH=arm64 $GO build -i -ldflags '-tmpdir '$GG_OBJ' -linkmode external' .
echo $GG_OBJ
ls $GG_OBJ

#ar rcs $ARCHIVE $GG_OBJ/*.o
#echo "Produced Go code as library in $ARCHIVE"
#ar rcs $SRCROOT/go-output.a $GG_OBJ/*.o

#cp $ARCHIVE $SRCROOT/ipfsios/
#echo "Produced Go code as library in $SRCROOT"

#echo "Go Build Completed."
