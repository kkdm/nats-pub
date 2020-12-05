package main

type Opts struct {
    NatsServer string `short:"s" long:"server" description:"nats server url" required:"true"`
    Cluster    string `short:"c" long:"cluster" description:"nats cluster name" required:"true"`
    LogPath    string `short:"l" long:"log-path" description:"log file path" default:"./server.log"`
    Verbose    bool   `short:"v" long:"verbose" description:"verbose log"`
}
