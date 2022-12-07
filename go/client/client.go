package main

import (
	"bytes"
        "encoding/json"
        "net/http"
	"math/rand"
	"fmt"
	"time"
	"flag"
	"os"
)

var RUN_FLAG = flag.String("run_type", "-", "Run Type...")

type SensorData struct {
	DeviceName     string       `json:"deviceName"`
	User           string       `json:"user"`
        Longitude      string       `json:"longitude"`
        Latitude       string       `json:"latitude`
        AirTemperature string       `json:"airTemperature"`
}

func main() {

	flag.Parse()
	runFlag := *RUN_FLAG

	if(runFlag == "-"){
		flag.PrintDefaults()
		os.Exit(0)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i:=0; i<20; i++ {
		tmpLongi := (r1.Float64()*120)+7
		tmpLati := (r1.Float64()*34)+4
		tmpTemp := (r1.Float64()*36)+2

		if runFlag == "type_normal" {
                        tmpJson := SensorData {"RaspberryPi", "sjjeon", fmt.Sprintf("%f", tmpLongi), fmt.Sprintf("%f", tmpLati), fmt.Sprintf("%f", tmpTemp)}
                        pbytes, _ := json.Marshal(tmpJson)
                        buff := bytes.NewBuffer(pbytes)
                        http.Post("http://101.79.4.10:32225/app", "application/json", buff)
		} else if runFlag == "type_oomkilled" {
			tmpJson := SensorData {"RaspberryPi", "sjjeon", fmt.Sprintf("%f", tmpLongi), fmt.Sprintf("%f", tmpLati), fmt.Sprintf("%f", tmpTemp)}
			if i == 10 { tmpJson = SensorData {"RaspberryPi-", "sjjeon", fmt.Sprintf("%f", tmpLongi), fmt.Sprintf("%f", tmpLati), fmt.Sprintf("%f", tmpTemp)} }
			pbytes, _ := json.Marshal(tmpJson)
                        buff := bytes.NewBuffer(pbytes)
                        http.Post("http://101.79.4.10:32225/app", "application/json", buff)
		} else if runFlag == "type_error" {
			tmpJson := SensorData {"RaspberryPi", "sjjeon", fmt.Sprintf("%f", tmpLongi), fmt.Sprintf("%f", tmpLati), fmt.Sprintf("%f", tmpTemp)}
                        if i == 10 { tmpJson = SensorData {"RaspberryPi+", "sjjeon", fmt.Sprintf("%f", tmpLongi), fmt.Sprintf("%f", tmpLati), fmt.Sprintf("%f", tmpTemp)} }
			pbytes, _ := json.Marshal(tmpJson)
                        buff := bytes.NewBuffer(pbytes)
                        http.Post("http://101.79.4.10:32226/app", "application/json", buff)
                } else if runFlag == "type_normal2" {
			tmpJson := SensorData {"RaspberryPi", "sjjeon", fmt.Sprintf("%f", tmpLongi), fmt.Sprintf("%f", tmpLati), fmt.Sprintf("%f", tmpTemp)}
			pbytes, _ := json.Marshal(tmpJson)
                        buff := bytes.NewBuffer(pbytes)
                	http.Post("http://101.79.4.10:32226/app", "application/json", buff)
                }
		//pbytes, _ := json.Marshal(tmpJson)
		//buff := bytes.NewBuffer(pbytes)
		//http.Post("http://101.79.4.10:32225/app", "application/json", buff)

		time.Sleep(1*time.Second)
	}
}
