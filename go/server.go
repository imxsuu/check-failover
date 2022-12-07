// main.go
package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "time"
    "syscall"
)

type SensorData struct {
    DeviceName     string       `json:"deviceName"`
    User           string       `json:"user"`
    Longitude      string       `json:"longitude"`
    Latitude       string       `json:"latitude`
    AirTemperature string       `json:"airTemperature"`
}

type appHandler struct{}

func (f *appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        var sensorData SensorData
        json.NewDecoder(r.Body).Decode(&sensorData)
        data, _ := json.Marshal(sensorData)
        now := time.Now()
        customTime := now.Format("2006-01-02 15:04:05")

        if sensorData.DeviceName == "RaspberryPi" {
                w.Header().Add("content-type", "application/json")
                w.WriteHeader(http.StatusCreated)
                fmt.Println("[RECV] [" + customTime + "]\n" + string(data) + "\n")
                fmt.Fprint(w, string(data))
        } else if sensorData.DeviceName == "RaspberryPi-" {
                w.Header().Add("content-type", "application/json")
                w.WriteHeader(http.StatusCreated)
                fmt.Println("[RECV] [" + customTime + "]\n" + string(data) + "\n")
                fmt.Fprint(w, string(data))
                size := 1024 * 1024 * 1024
                for j := 0; j < 100; j += 1 {
                        a := make([]int, size)
                        for i := 0; i < size; i += 1 {
                                a[i] = i
                        }
                        a = nil
                }
        } else if sensorData.DeviceName == "RaspberryPi+" {
                w.Header().Add("content-type", "application/json")
                w.WriteHeader(http.StatusCreated)
                fmt.Println("[RECV] [" + customTime + "]\n" + string(data) + "\n")
                fmt.Fprint(w, string(data))
		parentpid := syscall.Getppid()
                syscall.Kill(parentpid, syscall.SIGTERM)
        }
}

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprint(w, "Recived\n")
        })
        http.Handle("/app", &appHandler{})
        http.ListenAndServe(":9010", nil)
}
