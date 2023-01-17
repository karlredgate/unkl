
package main

import (
    "embed"
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "runtime"

    "github.com/zserge/lorca"
)

//go:embed www
var fs embed.FS

func waitForUI( ui lorca.UI ) {
    channel := make( chan os.Signal )
    signal.Notify( channel, os.Interrupt )
    select {
    case <-channel:
    case <-ui.Done():
    }
}

func createListener() (net.Listener) {
    listener,err := net.Listen( "tcp", "127.0.0.1:0" )
    if err != nil {
        log.Fatal( err )
    }
    go http.Serve( listener, http.FileServer(http.FS(fs)) )
    return listener
}

func createUI() (lorca.UI) {
    args := []string{}
    if runtime.GOOS == "linux" {
        args = append( args, "--class=Lorca" )
    }
    ui,err := lorca.New("", "", 480, 320, args...)
    if err != nil {
            log.Fatal(err)
    }
    return ui
}

func main() {
    fmt.Println("unkl")

    ui := createUI()
    defer ui.Close()
    listener := createListener()

    ui.Load( fmt.Sprintf("http://%s/www", listener.Addr()) )

    waitForUI( ui )
    log.Println( "exiting..." )
}

// vim: autoindent expandtab sw=4
