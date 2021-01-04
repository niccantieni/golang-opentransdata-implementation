package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/niccantieni/opentransdata"
	"log"
	"os"
	"strings"
	"time"
)

var (
	OTDApiKey = goDotEnvVariable("APIKEY_TRANSDATA")
	stationID = "8509184"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	//set time for request (point of time of interest)
	now := time.Now()
	//Format time as 2021-04-13T22:10:56Z, in OTD interpreted as Zulu (UTC)
	depArrTime := strings.Split(now.Format(time.RFC3339), "+")[0]

	//create request model
	req := opentransdata.NewOTDRequest("", stationID, depArrTime)

	//set parameters
	req.Parameters.IncludeRealtimeData = true
	req.Parameters.IncludeOnwardCalls = true
	req.Parameters.IncludePreviousCalls = true
	req.Parameters.NumberOfResults = "5"

	//Create and send the Request
	resp, data, err := opentransdata.CreateRequest(OTDApiKey, req)
	if err != nil {
		fmt.Println(resp, err)
	}
	fmt.Println(string(data))

	struc, err := opentransdata.ParseXML(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(struc)

}
