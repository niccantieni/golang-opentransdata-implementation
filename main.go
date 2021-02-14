package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/niccantieni/opentransdata"
	"log"
	"os"
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
	trias, _ := getTriasNow(stationIDBonaduz)
	stop := getNextDeparture(trias)

	txt := stop.Service.PublishedLineName.Text
	time := stop.ThisCall.CallAtStop.ServiceDeparture.TimetabledTime.String
	fmt.Println(txt, time)
}

func getTriasNow(stationID string) (trias opentransdata.Trias, err error) {

	request := opentransdata.TemplateOTDRequestNow()
	request.StopPointRef = stationID

	//Create and send the request
	data, err := opentransdata.CreateRequest(OTDApiKey, request)
	if err != nil {
		return trias, err
	}

	//parse the response
	trias, err = opentransdata.ParseXML(data)
	if err != nil {
		return trias, err
	}

	return trias, err
}

func getNextDeparture(trias opentransdata.Trias) (nextDeparture opentransdata.StopEvent) {
	nextDeparture = trias.ServiceDelivery.DeliveryPayload.StopEventResponse.StopEventResult[0].StopEvent

	return nextDeparture
}
