
GOROOT_BOOTSTRAP=/usr/local/Cellar/go/1.4.2/libexec/

all:	ipfsios 

goios.ios64-new/README.md:
	git clone https://minux@bitbucket.org/minux/goios.git goios.ios64-new
	cd goios.ios64-new && git checkout -b ios64-new origin/ios64-new 

go-toolchain:	goios.ios64-new/README.md
	cd goios.ios64-new/src && GOROOT_BOOTSTRAP=${GOROOT_BOOTSTRAP} GOARM=7 CGO_ENABLED=1 GOARCH=arm64 CC_FOR_TARGET=`pwd`/../misc/ios/clangwrap.sh CXX_FOR_TARGET=`pwd`/../misc/ios/clangwrap.sh ./make.bash

ipfsios: 	go-toolchain
	cd ipfsios &&  xcodebuild -target ipfsios -arch arm64

clean:
	cd ipfsios && xcodebuild clean && rm -rf build/
