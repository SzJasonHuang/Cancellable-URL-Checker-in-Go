package main

import (
	"time"
)

type shop struct {
	name  string
	cents int
	delay time.Duration
}

type result struct {
	shop  string
	price string
}
