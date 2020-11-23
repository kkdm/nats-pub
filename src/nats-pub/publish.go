package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/google/logger"
    "github.com/nats-io/stan.go"
	nats "github.com/nats-io/nats.go"
)

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
