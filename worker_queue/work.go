package main

import (
	"time"
)

// struct to holds work request
type WorkRequest struct {
	Name  string
	Delay time.Duration
}
