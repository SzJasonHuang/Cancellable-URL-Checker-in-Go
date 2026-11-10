package main

import (
	"html/template"
	"net/http"
	"strings"
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

var page = template.Must(template.ParseFiles("index.html"))

func main() {
	shops := []shop{
		{"Shop A", 249, 200 * time.Millisecond},
		{"Shop B", 199, 800 * time.Millisecond},
		{"Shop C", 149, 5 * time.Second},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		data := struct {
			Card string
			Results []result
			Best string
		}{
			Card: strings.TrimSpace(r.URL.Query().Get("card"))
		}
	}
}
