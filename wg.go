package main

import (
    "fmt"
    "sync"
)

var x = 0
func increment(wg *sync.WaitGroup, ch chan bool){
    ch <- true
    x = x + 1
    <- ch
    wg.Done()
}
