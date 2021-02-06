package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/niccantieni/opentransdata"
	"log"
	"os"
	"time"
)

var (
	OTDApiKey        = goDotEnvVariable("APIKEY_TRANSDATA")
	stationIDBonaduz = "8509184"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	//Load Location in Zürich, everything else does not make much sense
	zurich, err := time.LoadLocation("Europe/Zurich")
	if err != nil {
		fmt.Println(err)
	}

	//current time in Zurich
	now := time.Now().In(zurich)

	//format timestamp as timestamp correctly; OTD interprets this as localtime (somehow, but not really?!?)
	// whatever, it works like this to get the current (as in right now, instant) events.
	depArrTime := now.Format(opentransdata.ShortRFC3339)

	//create request model
	req := opentransdata.NewOTDRequest("", stationIDBonaduz, depArrTime, "5",
		"departure", true, true, true)

	//Create and send the request
	resp, data, err := opentransdata.CreateRequest(OTDApiKey, req)
	if err != nil {
		fmt.Println(resp, err)
	}
	//fmt.Println(string(data))

	//parse the response
	struc, err := opentransdata.ParseXML(data)
	if err != nil {
		fmt.Println(err)
	}

	check := struc.ServiceDelivery.DeliveryPayload.StopEventResponse.StopEventResult[0].StopEvent.ThisCall.CallAtStop
	fmt.Println(check)
}
