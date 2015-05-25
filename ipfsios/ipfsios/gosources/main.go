package main

/*
// adjust LDFLAGS to produce linkable code - note Go starts first, so
// be sure to change AppDelegate main() for chaining ..
#cgo LDFLAGS: -Wl,-U,_iosmain,-U,_PopUpDialogBox,-U,_getDocumentsDirForIPFS
#cgo LDFLAGS: -framework UIKit -framework Foundation -framework CoreGraphics
extern char* PopUpDialogBox(char* msg);
extern char *getDocumentsDirForIPFS();
extern void iosmain(int argc, char *argv[]);
*/
import "C"

import (
	"code.google.com/p/go.net/context"
	"fmt"
	core "github.com/ipfs/go-ipfs/core"
	corenet "github.com/ipfs/go-ipfs/core/corenet"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
	// "github.com/jbenet/go-ipfs/cmd/ipfs"
	"io"
	"net/http"
	"os"

	"bytes"
	"errors"
	// context "github.com/ipfs/go-ipfs/Godeps/_workspace/src/golang.org/x/net/context"
	assets "github.com/ipfs/go-ipfs/assets"
	// cmds "github.com/ipfs/go-ipfs/commands"
	coreunix "github.com/ipfs/go-ipfs/core/coreunix"
	namesys "github.com/ipfs/go-ipfs/namesys"
	config "github.com/ipfs/go-ipfs/repo/config"
	uio "github.com/ipfs/go-ipfs/unixfs/io"
	u "github.com/ipfs/go-ipfs/util"
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

// func StartGoServer() {
// 	fmt.Fprintf(os.Stderr, "Starting net/http/Server ...\n")
// 	http.HandleFunc("/", HandleHttpRequest)
// 	http.ListenAndServe(":6060", nil)
// }

const nBitsForKeypairDefault = 2048

var errRepoExists = errors.New(`ipfs configuration file already exists!
Reinitializing would overwrite your keys.
(use -f to force overwrite)
`)

func initWithDefaults(out io.Writer, repoRoot string) error {
	err := doInit(out, repoRoot, false, nBitsForKeypairDefault)
	return err
}

func doInit(out io.Writer, repoRoot string, force bool, nBitsForKeypair int) error {
	if _, err := fmt.Fprintf(out, "initializing ipfs node at %s\n", repoRoot); err != nil {
		return err
	}

	if fsrepo.IsInitialized(repoRoot) && !force {
		return errRepoExists
	}

	conf, err := config.Init(out, nBitsForKeypair)
	if err != nil {
		return err
	}

	if fsrepo.IsInitialized(repoRoot) {
		if err := fsrepo.Remove(repoRoot); err != nil {
			return err
		}
	}

	if err := fsrepo.Init(repoRoot, conf); err != nil {
		return err
	}

	if err := addDefaultAssets(out, repoRoot); err != nil {
		return err
	}

	return initializeIpnsKeyspace(repoRoot)
}

func addDefaultAssets(out io.Writer, repoRoot string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := fsrepo.Open(repoRoot)
	if err != nil { // NB: repo is owned by the node
		return err
	}

	nd, err := core.NewIPFSNode(ctx, core.Offline(r))
	if err != nil {
		return err
	}
	defer nd.Close()

	dirb := uio.NewDirectory(nd.DAG)

	// add every file in the assets pkg
	for fname, file := range assets.Init_dir {
		buf := bytes.NewBufferString(file)
		s, err := coreunix.Add(nd, buf)
		if err != nil {
			return err
		}

		k := u.B58KeyDecode(s)
		if err := dirb.AddChild(fname, k); err != nil {
			return err
		}
	}

	dir := dirb.GetNode()
	dkey, err := nd.DAG.Add(dir)
	if err != nil {
		return err
	}

	if err := nd.Pinning.Pin(ctx, dir, true); err != nil {
		return err
	}

	if err := nd.Pinning.Flush(); err != nil {
		return err
	}

	if _, err = fmt.Fprintf(out, "to get started, enter:\n"); err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "\n\tipfs cat /ipfs/%s/readme\n\n", dkey)
	return err
}

func initializeIpnsKeyspace(repoRoot string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := fsrepo.Open(repoRoot)
	if err != nil { // NB: repo is owned by the node
		return err
	}

	nd, err := core.NewIPFSNode(ctx, core.Offline(r))
	if err != nil {
		return err
	}
	defer nd.Close()

	err = nd.SetupOfflineRouting()
	if err != nil {
		return err
	}

	return namesys.InitializeKeyspace(ctx, nd.DAG, nd.Namesys, nd.Pinning, nd.PrivateKey)
}

func ipfs_server() {
	// Basic ipfsnode setup
	docsDir := C.GoString(C.getDocumentsDirForIPFS())
	r, err := fsrepo.Open(docsDir)
	if err != nil {
		rpipe, wpipe := io.Pipe()
		go io.Copy(os.Stdout, rpipe)
		defer wpipe.Close()
		if err := doInit(wpipe, docsDir, false, 1024); err != nil {
			panic(err)
		}
		// }()
		// res.SetOutput(rpipe)

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

// Go main entry point
func main() {
	fmt.Fprintf(os.Stderr, "Startup of Golang iosmain ::\n")
	go ipfs_server()
	C.iosmain(0, nil)
}
