package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json:"number"`
}

func main() {

	url := "https://api.apify.com/v2/key-value-stores/toDWvRj1JpTXiM8FF/records/LATEST?disableRedirect=true"

	spaceClient := http.Client{
		Timeout: time.Second * 100, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var f interface{}
	jsonErr := json.Unmarshal(body, &f)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	India := f.(map[string]interface{})	
	Regions:=India["regionData"].([]interface{})
	fmt.Println("India Covid Tracker\n")
	fmt.Println(" => Total cases  : ", India["totalCases"])
	fmt.Println(" => Active Cases : ", India["activeCases"])
	fmt.Println(" => New Cases 	  : ", India["activeCasesNew"])
	var c byte
	fmt.Println("\n Press any key to continue : ")
	fmt.Scanf("%c",&c)
	fmt.Println("\nRegion Wise Split\n")
	for _, v := range Regions {
		vv := v.(map[string]interface{})
		fmt.Println(vv["region"])
		for key, val := range vv {
			if key == "region"{
				continue;
			}
			fmt.Println(" => ",key," : ",val)
		}
		fmt.Println("\n Press any key to continue : ")
		fmt.Scanf("%c",&c)
	}

}	
