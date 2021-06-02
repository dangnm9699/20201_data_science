package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"goex/model"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	from := flag.String("from", "VND", "SRC Currency")
	to := flag.String("to", "USD", "DST Currency")
	flag.Parse()
	API := fmt.Sprintf("https://www.xe.com/api/stats.php?fromCurrency=%s&toCurrency=%s", *from, *to)
	resp, err := http.Get(API)
	if err != nil {
		log.Fatal("1", err)
	}
	nbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("2", err)
	}
	var stat model.Stat
	err = json.Unmarshal(nbytes, &stat)
	if err != nil {
		log.Fatal("3", err)
	}
	fmt.Printf("From %s To %s: %f\n", *from, *to, stat.Payload.Last30Days.Average)
}
