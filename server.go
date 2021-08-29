package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/products/{fetchCountPercentage}", products).Methods("GET")
	fmt.Println("Hello every one")
	fmt.Println("Serving @ 127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the root page"))
}

func products(w http.ResponseWriter, r *http.Request) {

	fetchCountPercentage, errInput := strconv.ParseFloat(mux.Vars(r)["fetchCountPercentage"], 64)

	fetchCount := 0

	if errInput != nil {
		fmt.Println(errInput.Error())
	} else {
		fetchCount = int(float64(len(productList)) * fetchCountPercentage / 100)
		if fetchCount > len(productList) {
			fetchCount = len(productList)
		}
	}
	//write to response
	jsonList, err := json.Marshal(productList[0:fetchCount])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(jsonList)
	}

}

type product struct {
	Name  string
	Price float64
	Coun  int
}

var productList = []product{

	{"p1", 25.0, 30},
	{"p2", 20.0, 10},
	{"p3", 250.0, 320},
	{"p4", 256.0, 730},
	{"p5", 24.0, 340},
	{"p6", 10.0, 300},
	{"p7", 100.0, 230},
	{"p8", 2543.0, 120},
	{"p9", 255.0, 10},
	{"p10", 175.0, 20},
	{"p5", 24.0, 340},
	{"p6", 10.0, 300},
	{"p7", 100.0, 230},
	{"p8", 2543.0, 120},
	{"p9", 255.0, 10},
	{"p10", 175.0, 20},
}
