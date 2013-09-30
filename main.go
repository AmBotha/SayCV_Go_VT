// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

import (
	"github.com/lxn/walk"
)

func main() {
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}
	
	if _, err := runDialog(mw); err != nil {
		log.Fatal(err)
	}
	
	// Run the message loop.
	mw.Run()
}