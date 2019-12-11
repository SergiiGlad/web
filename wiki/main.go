// +build linux darwin
// !cgo

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// A build constraint

package main

import (
  "fmt"
	"net/http"
  "log"

)

func handler(w http.ResponseWriter, r *http.Request) {
    name := "Golang"
    if len(r.URL.Path[1:]) != 0 {
      name = r.URL.Path[1:]
      }
    fmt.Fprintf(w, "Hi there, %s!", name )
}

func main() {

  http.HandleFunc("/", handler)

  log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))

}
