package src

import (
    "os"
    "os/signal"
)

var ServerSigChan chan os.Signal

func init() {
    ServerSigChan = make(chan os.Signal)
}

func ShutDownServer(err error) {
    ServerSigChan<-os.Interrupt
}

func SeverNotify() {
    signal.Notify(ServerSigChan, os.Interrupt)
    <-ServerSigChan
}
