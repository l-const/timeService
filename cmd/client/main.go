package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var host, port string
	var err error
	if len(os.Args) == 3 {
		host = os.Args[1]
		port = os.Args[2]
	} else {
		err = godotenv.Load("./config/.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		host = os.Getenv("HOST")
		port = os.Getenv("PORT")
	}
	var url string = "http://" + host + ":" + port + "/ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(body)
}
