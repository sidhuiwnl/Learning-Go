package main

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type client struct {
	tokens int
	lastRefill time.Time
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

const (
	maxTokens = 5
	refillTimeSec = 10
)

func RateLimiter(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		mu.Lock()

		//returns the c is the value stored at that key and exisits says true/false
		c, exists := clients[ip]

		if !exists{
			clients[ip] = &client{
				tokens:     maxTokens,
				lastRefill: time.Now(),
			}

			c = clients[ip]
		}

		if time.Since(c.lastRefill) > time.Second * refillTimeSec{
			c.tokens = maxTokens
			c.lastRefill = time.Now()
		}

		if c.tokens < 0 {
			mu.Unlock()
			http.Error(w, "Too many requests", 429)
			return
		}

		c.tokens--
		mu.Unlock()

		next(w, r)

	}
}