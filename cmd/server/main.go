package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"timeservice/lib"

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

		fmt.Printf("Loaded from .env host=%v port=%v", host, port)
	}
	http.HandleFunc("/ptlist", mainHandler)
	err = http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tz, t1, t2, period := query["period"][0], query["tz"][0], query["t1"][0], query["t2"][0]
	timeStamps, err := lib.GenerateTimeStamps(t1, t2, tz, period)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		errorResponse := ErrorResponse{
			Status: "status",
			Desc:   err.Error(),
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(timeStamps)
	}

}
