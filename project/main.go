package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type shop struct {
	name  string
	cents int
	delay time.Duration
}

type result struct {
	Shop  string
	Price string
}

var page = template.Must(template.ParseFiles("index.html"))

func main() {
	shops := []shop{
		{"Shop A", 249, 200 * time.Millisecond},
		{"Shop B", 199, 800 * time.Millisecond},
		{"Shop C", 149, 5 * time.Second},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Card    string
			Results []result
			Best    string
		}{Card: strings.TrimSpace(r.URL.Query().Get("card"))}

		if data.Card != "" {
			ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
			defer cancel()

			data.Results = make([]result, len(shops))
			var wg sync.WaitGroup

			for i, s := range shops {

				wg.Go(func() {
					data.Results[i] = checkShop(ctx, s)
				})
			}
			wg.Wait()
			bestCents := 0
			for i, res := range data.Results {
				if res.Price == "Timed out" || res.Price == "Canceled" {
					continue
				}
				if bestCents == 0 || shops[i].cents < bestCents {
					bestCents = shops[i].cents
					data.Best = res.Price + " at " + res.Shop
				}
			}
		}
		if err := page.Execute(w, data); err != nil {
			log.Print(err)
		}
	})
	log.Println("Open http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
