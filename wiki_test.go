// +build linux darwin
// !cgo

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// A build constraint

package main

import (
  "testing"
  "fmt"
  "net/http/httptest"
  "io/ioutil"
  "log"

)

func TestHandler(t *testing.T) {

  req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

  if resp.StatusCode != 201 {
    t.Errorf("Server http not run, got: %d, want: %d.",resp.StatusCode, 200)
  }
}
