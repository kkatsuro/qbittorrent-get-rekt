package main

import (
    "fmt"
    "log"
    "strings"
    "time"

    "github.com/BurntSushi/xgbutil"
    "github.com/BurntSushi/xgbutil/ewmh"
    "github.com/BurntSushi/xgbutil/icccm"
)

// Improvement idea: don't work on list of windows in a loop,
// loop on it at start and wait for window appearance event instead.

func main() {
    X, err := xgbutil.NewConn()
    if err != nil {
        log.Fatal(err)
    }

    window_got_rekt := false
    for !window_got_rekt {
        clientids, err := ewmh.ClientListGet(X)
        if err != nil {
            log.Fatal(err)
        }

        for _, clientid := range clientids {
            name, err := ewmh.WmNameGet(X, clientid)
            if err != nil || len(name) == 0 {
                name, err = icccm.WmNameGet(X, clientid)
                if err != nil || len(name) == 0 {
                    name = "N/A"
                }
            }

            if !strings.Contains(name, "qBittorrent") {
                continue
            }
            fmt.Printf("closing '%s' window\n", name)
            err = ewmh.CloseWindow(X, clientid)
            window_got_rekt = true
        }
        time.Sleep(10 * time.Millisecond)
    }
}
