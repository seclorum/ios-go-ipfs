package main
/*
// adjust LDFLAGS to produce linkable code - note Go starts first, so
// be sure to change AppDelegate main() for chaining ..
#cgo LDFLAGS: -Wl,-U,_iosmain,-U,_PopUpDialogBox
#cgo LDFLAGS: -framework UIKit -framework Foundation -framework CoreGraphics
extern char* PopUpDialogBox(char* msg);
extern void iosmain(int argc, char *argv[]);
*/
import "C"

import (
	"fmt"
	"net/http"
	"os"
	//"code.google.com/p/go.net/context"
	//core "github.com/ipfs/go-ipfs/core"
	//corenet "github.com/ipfs/go-ipfs/core/corenet"
	//fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
)

//export AddUsingGo
func AddUsingGo(a int, b int) int {
	return a + b
}

// basic server for debugging ..
func HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	userInput := C.GoString(C.PopUpDialogBox(C.CString(r.URL.Path[1:])))
	fmt.Fprintf(w, "Hello from Go on an iPhone! Response: %s", userInput)
}

func StartGoServer() {
	fmt.Fprintf(os.Stderr, "Starting net/http/Server ...\n")
	http.HandleFunc("/", HandleHttpRequest)
	http.ListenAndServe(":6060", nil)
}
/*
func ipfs_server() {
	// Basic ipfsnode setup
	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		panic(err)
	}

	nb := core.NewNodeBuilder().Online()
	nb.SetRepo(r)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd, err := nb.Build(ctx)
	if err != nil {
		panic(err)
	}

	list, err := corenet.Listen(nd, "/app/whyrusleeping")
	if err != nil {
		panic(err)
	}
	fmt.Printf("I am peer: %s\n", nd.Identity)

	for {
		con, err := list.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer con.Close()

		fmt.Fprintln(con, "Hello! This is whyrusleepings awesome ipfs service")
		fmt.Printf("Connection from: %s\n", con.Conn().RemotePeer())
	}
}
*/

// Go main entry point
func main() {
	fmt.Fprintf(os.Stderr, "Startup of Golang iosmain ::\n")
	go StartGoServer()
	C.iosmain(0, nil)
}
