package main

import (
	"testing"
	"time"
)

func TestHelloworld(t *testing.T) {
	timestamp := time.Now().Unix()
	t.Log(timestamp)
}
