// Copyright 2016 The Gini Authors. All rights reserved.  Use of this source
// code is governed by a license that can be found in the License file.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var sigs = make(chan os.Signal, 1)

func init() {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("\nc interrupted\ns UNKNOWN")
				os.Exit(1)
			}
		}
	}()
}
