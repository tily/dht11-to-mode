package main

import (
	"github.com/d2r2/go-dht"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tily/modeclient"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	interval, _ := strconv.Atoi(os.Getenv("INTERVAL"))

	endpoint := os.Getenv("MODE_ENDPOINT")
	deviceId, _ := strconv.Atoi(os.Getenv("MODE_DEVICE_ID"))
	deviceAPIKey := os.Getenv("MODE_DEVICE_API_KEY")

	device := modeclient.NewDevice(endpoint, deviceAPIKey, deviceId)

	sEvent := modeclient.Event{
		EventType: "dht11-start",
		EventData: map[string]int{"value": 1},
	}
	trigger(device, sEvent)

	crawl(device, interval)
}

func crawl(device modeclient.Device, interval int) {
	for true {
		temperature, humidity, retried, _ := dht.ReadDHTxxWithRetry(dht.DHT11, 4, true, 10)

		tEvent := modeclient.Event{
			EventType: "dht11-temperature",
			EventData: map[string]float32{"value": temperature},
		}
		hEvent := modeclient.Event{
			EventType: "dht11-humidity",
			EventData: map[string]float32{"value": humidity},
		}
		rEvent := modeclient.Event{
			EventType: "dht11-retried",
			EventData: map[string]int{"value": retried},
		}

		trigger(device, tEvent)
		trigger(device, hEvent)
		trigger(device, rEvent)

		time.Sleep(time.Second * time.Duration(interval))
	}
}

func trigger(device modeclient.Device, event modeclient.Event) {
	log.Printf("[dht11] Triggering event: %+v\n", event)
	_, err := device.TriggerEvent(event)
	if err != nil {
		log.Printf("[dht11] Error happened: %s", err)
	}
}
