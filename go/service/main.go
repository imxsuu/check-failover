package main

import (
        "flag"
        "os"
	"os/exec"
	"log"
)

var FLAG = flag.String("service_name", "-", "Service Name...")

func main() {

        flag.Parse()
        Flag := *FLAG

        if(Flag == "-"){
                flag.PrintDefaults()
                os.Exit(0)
        }

	if Flag == "data" {
		cmd := exec.Command("bash", "-c", "kubectl create -f ../kube/podImagePullBackOff.yaml")
		if err := cmd.Run(); err != nil {
                        log.Fatal(err)
                }
	}
}
