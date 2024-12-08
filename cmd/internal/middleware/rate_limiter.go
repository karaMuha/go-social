package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func RateLimiter(next http.Handler) http.Handler {
	var mutex sync.Mutex
	clients := make(map[string]*client)

	// loop through the client ips every minute and clean up those who haven't send a request for atleast three minutes
	go func() {
		for {
			time.Sleep(time.Minute)
			mutex.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mutex.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mutex.Lock()

		if _, found := clients[ip]; !found {
			clients[ip] = &client{
				limiter: rate.NewLimiter(rate.Limit(3), 30),
			}
		}

		clients[ip].lastSeen = time.Now()

		mutex.Unlock()

		if !clients[ip].limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
