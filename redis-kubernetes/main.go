package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
func getQuoteFromAPI() (*QuoteResponse, error) {
	API_URL := "http://quotes.rest/qod.json"
	resp, err := http.Get(API_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("Quote  API returned:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		quoteResp := &QuoteResponse{}
		json.NewDecoder(resp.Body).Decode(quoteResp)
		return quoteResp, nil
	} else {
		return nil, errors.New("Could not get Quote from API")
	}

}

func main() {

	//Create Redis CLient
	var (
		host     = getEnv("REDIS_HOST", "localhost")
		port     = string(getEnv("REDIS_PORT", "6379"))
		password = getEnv("REDIS_PASSWORD", "")
	)

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", indexhandler)
	r.HandleFunc("/qod", qodhandler(client))
	//Instead of writing in this manner
	//http.ListenAndServe(":8088", r)
	//log.Println("Starting server")
	srv := &http.Server{
		Handler: r,
		Addr:    ":8088",
	}
	go func() {
		log.Println("Starting server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	waitForShutdown(srv)

}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello,please hit /qod API to get quote od the day"))
}
func qodhandler(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		d := t.Format("2019-10-28")
		val, err := client.Get(d).Result()
		if err == redis.Nil {
			log.Println("cache miss for date", d)
			quoteResp, err := getQuoteFromAPI()
			if err != nil {
				w.Write([]byte("Sorry!we couldn't get Quote of the day."))
				return
			}
			quote := quoteResp.Contents.Quotes[0].Quote
			client.Set(d, quote, 24*time.Hour)
			w.Write([]byte(quote))
		} else {
			log.Println("Cache hit for date", d)
			w.Write([]byte(val))
		}
	}
}
