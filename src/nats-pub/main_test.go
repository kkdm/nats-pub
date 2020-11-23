package main

import (
    "testing"
    "github.com/gorilla/mux"

    "github.com/stretchr/testify/assert"
)

func TestGetRouter(t *testing.T) {
    r := getRouter()

    expect := mux.NewRouter().StrictSlash(true)
    expect.HandleFunc("/publish", publishHandler).
        Methods("POST").
        Headers("Content-Type", "application/json")
    assert.ObjectsAreEqual(expect, r)
}
