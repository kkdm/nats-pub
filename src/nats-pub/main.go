package main

import (
    "os"
    "log"
    "net/http"
    "github.com/google/logger"
    "github.com/gorilla/mux"
    "github.com/jessevdk/go-flags"
)

var opts Opts
var parser = flags.NewParser(&opts, flags.Default)

func getRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/publish", publishHandler).
        Methods("POST").
        Headers("Content-Type", "application/json")

    return router
}

func main() {
    if _, err := parser.Parse(); err != nil {
        os.Exit(1)
    }

    logfile, _ := os.OpenFile(opts.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    defer logfile.Close()

    logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    defer logger.Init("Logger", opts.Verbose, true, logfile).Close()

    logger.Infof("server params: server: %s, cluster: %s, log-path: %s, verbose: %t",
        opts.NatsServer, opts.Cluster, opts.LogPath, opts.Verbose)

    router := getRouter()
    logger.Infof("starting server: :8080")
    logger.Fatal(http.ListenAndServe(":8080", router))
}
