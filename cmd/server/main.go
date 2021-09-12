package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Desc   string `json:"desc"`
}

type TimeStampResponse []string

func main() {
	var err error
	var host, port string
	if len(os.Args) == 3 {
		host = os.Args[1]
		port = os.Args[2]
	} else {
		host, errH := os.LookupEnv("HOST")
		port, errP := os.LookupEnv("PORT")
		if !errH || !errP {
			err = godotenv.Load("./config/.env")
			if err != nil {
				log.Fatal("Error loading .env file")
			}
		}
		// host = os.Getenv("HOST")
		// port = os.Getenv("PORT")
		fmt.Printf("Loaded from .env host=%v port=%v", host, port)
	}

	http.HandleFunc("/ptlist", mainHandler)

	err = http.ListenAndServe(host+":"+port, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(200)
	// fmt.Fprintf(w, "Hello there from kostas %s\n", r.URL.Query()["t1"])

	// query := r.URL.Query()
	// period := query["period"][0]
	// tz := query["tz"][0]
	// t1 := query["t1"][0]
	// t2 := query["t2"][0]

	// fmt.Printf("%s %s %s %s\n", period, tz, t1, t2)
	// loc := setLocation(tz)
	// fmt.Printf(loc.String())

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(404)
	// errR := ErrorResponse{
	// Status: "error",
	// Desc:   "Unsupported period",
	// }

	// json.NewEncoder(w).Encode(errR)
	w.WriteHeader(200)
	times := TimeStampResponse{
		"2344555555",
		"3434343434",
		"2323232323",
		"23232323232",
	}

	json.NewEncoder(w).Encode(times)

}
