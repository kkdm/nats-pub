package main

import (
    "os"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/google/logger"
    "github.com/gorilla/mux"
    "github.com/jessevdk/go-flags"
    "github.com/nats-io/stan.go"
	nats "github.com/nats-io/nats.go"
)

type Item struct {
    Subject      string `json:"subject"`
    Message      string `json:"message"`
}

type Result struct {
    Message string `json:"message"`
    Ok      bool   `json:"ok"`
}

type Opts struct {
    NatsServer string `short:"s" long:"server" description:"nats server url" required:"true"`
    Cluster    string `short:"c" long:"cluster" description:"nats cluster name" required:"true"`
    LogPath    string `short:"l" long:"log-path" description:"log file path" default:"./server.log"`
    Verbose    bool   `short:"v" long:"verbose" description:"verbose log"`
}

var opts Opts
var parser = flags.NewParser(&opts, flags.Default)

func responseJson(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

func publish(w http.ResponseWriter, r *http.Request) {
    var item Item

    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &item)

    nc, err := nats.Connect(opts.NatsServer)
    if err != nil {
        logger.Errorf("%v", err)

        responseJson(w, http.StatusInternalServerError,
            Result{
                Message: fmt.Sprintf("%v", err),
                Ok: false,
            })

        return
    }
    defer nc.Close()

    sc, err := stan.Connect(opts.Cluster, "client-1", stan.NatsConn(nc))
    if err != nil {
        logger.Errorf("connection failed: %v", err)
        responseJson(w, http.StatusInternalServerError,
            Result{
                Message: fmt.Sprintf("%v", err),
                Ok: false,
            })

        return
    }

    err = sc.Publish(item.Subject, []byte(item.Message))
    if err != nil {
        logger.Errorf("publish failed: %v", err)

        responseJson(w, http.StatusInternalServerError,
            Result{
                Message: fmt.Sprintf("%v", err),
                Ok: false,
            })

        return
    }

    logger.Infof("published: [%s]: '%s'", item.Subject, item.Message)

    responseJson(w, http.StatusOK,
        Result{
            Message: "published",
            Ok: true,
        })
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/publish", publish).
        Methods("POST").
        Headers("Content-Type", "application/json")
    logger.Infof("starting server: :8080")
    logger.Fatal(http.ListenAndServe(":8080", myRouter))
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

    handleRequests()
}
