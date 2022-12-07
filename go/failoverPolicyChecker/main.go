package main

import (
	"fmt"
	"os/exec"
	"time"
	"sync"
	"strings"

)


var Wait_goFunc sync.WaitGroup

func check_ImagePullBackOff() {

	defer Wait_goFunc.Done()

	var podName string

	for {	
		
		cmd_checkImagePullBackOff := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep ImagePullBackOff | awk '{print $1}' | head -1")
                output_checkImagePullBackOff, err := cmd_checkImagePullBackOff.Output()
                if err != nil {
                        fmt.Println(err)
                }

		str_checkImagePullBackOff := string(output_checkImagePullBackOff)
		podName = str_checkImagePullBackOff
		podName = strings.ReplaceAll(podName, "\n", "")

		if podName == "" {
                        continue
                } else {
                        fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 ImagePullBackOff 상태로 감지되었습니다.")
                        time.Sleep(time.Second * 10)
                        retCmd := exec.Command("bash", "-c", "kubectl delete pod --grace-period=0 --force --namespace gedge-platform " + podName)
                        _, err := retCmd.Output()
                        if err != nil {
                                fmt.Println(err)
                        }
                        fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + " ImagePullBackOff 상태로 감지된 GEdge 응용 " + podName + "을 제거하였습니다.\n")
                }
	}
}

func check_Complete() {

	defer Wait_goFunc.Done()

	var podName string
	//now := time.Now()
	//secs := now.Unix()

	for {	
		
		fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] Check Complete status....")
		time.Sleep(time.Second * 2)

		cmd_checkComplete := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep Completed | awk '{print $1}' | head -1")
		output_checkComplete, err := cmd_checkComplete.Output()
		if err != nil {
			fmt.Println(err)
		}	
			
		str_checkComplete := string(output_checkComplete)
		podName = str_checkComplete
		podName = strings.ReplaceAll(podName, "\n", "")
		
		if podName == "" {
                        continue
                } else {
                        fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 ImagePullBackOff 상태로 감지되었습니다.")
                        time.Sleep(time.Second * 3)
                        retCmd := exec.Command("bash", "-c", "kubectl delete pod --grace-period=0 --force --namespace gedge-platform " + podName)
                        _, err := retCmd.Output()
                        if err != nil {
                                fmt.Println(err)
                        }
                        fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + " ImagePullBackOff 상태로 감지된 GEdge 응용 " + podName + "을 제거하였습니다.\n")
                }
	}
}

func check_OOMKilled() {

	defer Wait_goFunc.Done()

	var podName string

	for {	
		cmd_checkOOMKilled := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep OOMKilled | awk '{print $1}' | head -1")
		output_checkOOMKilled, err := cmd_checkOOMKilled.Output()
		if err != nil {
			fmt.Println(err)
		}	
			
		str_checkOOMKilled := string(output_checkOOMKilled)
		podName = str_checkOOMKilled
		podName = strings.ReplaceAll(podName, "\n", "")
		
		if podName == "" {
			continue
		} else {
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 OOMKilled 상태로 감지되었습니다.")
			time.Sleep(time.Second * 10)
			retCmd := exec.Command("bash", "-c", "kubectl delete pod --grace-period=0 --force --namespace gedge-platform " + podName)
			_, err := retCmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + " OOMKilled 상태로 감지된 GEdge 응용 " + podName + "을 제거하였습니다.\n")
		}
	}
}

func check_Running() {

	defer Wait_goFunc.Done()

	var podName string

	for {	
		//fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] Check Running status....")
		time.Sleep(time.Second * 2)

		cmd_checkRunning := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep Running | awk '{print $1}' | head -1")
		output_checkRunning, err := cmd_checkRunning.Output()
		if err != nil {
			fmt.Println(err)
		}	
			
		str_checkRunning := string(output_checkRunning)
		podName = str_checkRunning
		podName = strings.ReplaceAll(podName, "\n", "")
		
		if podName == "" {
			continue
		} else {
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 Running 상태로 감지되었습니다.")
			time.Sleep(time.Second * 3)
		}
	}
}


func check_Pending() {

	defer Wait_goFunc.Done()

	var podName string

	for {	
		
		//fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] Check Pending status....")
		time.Sleep(time.Second * 2)

		cmd_checkPending := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep Pending | awk '{print $1}' | head -1")
		output_checkPending, err := cmd_checkPending.Output()
		if err != nil {
			fmt.Println(err)
		}	
			
		str_checPending := string(output_checkPending)
		podName = str_checPending
		podName = strings.ReplaceAll(podName, "\n", "")
		
		if podName == "" {
			continue
		} else {
			fmt.Println(podName)
			time.Sleep(time.Second * 2)			
		}
	}
}


func check_Terminating() {

	defer Wait_goFunc.Done()

	var podName string

	for {	
		
		//fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] Check Pending status....")
		time.Sleep(time.Second * 2)

		cmd_checkTerminating := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep Terminating | awk '{print $1}' | head -1")
		output_checkTerminating, err := cmd_checkTerminating.Output()
		if err != nil {
			fmt.Println(err)
		}	
			
		str_checTerminating := string(output_checkTerminating)
		podName = str_checTerminating
		podName = strings.ReplaceAll(podName, "\n", "")
		
		if podName == "" {
			continue
		} else {
			fmt.Println(podName)
			time.Sleep(time.Second * 2)			
		}
	}
}


func check_Complete_logBackup() {

	defer Wait_goFunc.Done()

	var podName string

	for {	
		
		fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] Check Log backup process status....")
		time.Sleep(time.Second * 5)

		cmd_checkComplete_logBackup := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep Completed | awk '{print $1}' | head -1")
		output_checkComplete_logBackup, err := cmd_checkComplete_logBackup.Output()
		if err != nil {
			fmt.Println(err)
		}	

		str_checkComplete_logBackup := string(output_checkComplete_logBackup)
		podName = str_checkComplete_logBackup
		podName = strings.ReplaceAll(podName, "\n", "")

		if podName == "" {
			continue
		} else {
			fmt.Println(podName)
			time.Sleep(time.Second * 2)			
		}
	}
}


func check_Error() {

	defer Wait_goFunc.Done()

	var podName string

	for {	
		
		//fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] Check Error status....")
		time.Sleep(time.Second * 2)

		cmd_checkError := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep Error | awk '{print $1}' | head -1")
		output_checkError, err := cmd_checkError.Output()
		if err != nil {
			fmt.Println(err)
		}	

		str_checkError := string(output_checkError)
		podName = str_checkError
		podName = strings.ReplaceAll(podName, "\n", "")
                
		if podName == "" {
                        continue
                } else {
                        fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 Error 상태로 감지되었습니다.")
                        time.Sleep(time.Second * 10)
                        retCmd := exec.Command("bash", "-c", "kubectl delete pod --grace-period=0 --force --namespace gedge-platform " + podName)
                        _, err := retCmd.Output()
                        if err != nil {
                                fmt.Println(err)
                        }
                        fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + " Error 상태로 감지된 GEdge 응용 " + podName + "을 제거하였습니다.\n")
                }
	}
}


func check_gpuAllocation() {
		
	defer Wait_goFunc.Done()

	for {

		//fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] Check GPU allocating error....")
		time.Sleep(time.Second * 5)

		cmd_checkAllocation := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep UnexpectedAdmissionError | awk '{print $1}' | head -1")
		output_checkAllocation, err := cmd_checkAllocation.Output()
		if err != nil {
			fmt.Println(err)
		}	

		str_checkAllocation := string(output_checkAllocation)
		podName := str_checkAllocation
		podName = strings.ReplaceAll(podName, "\n", "")

		if podName == "" {
			continue
		} else {
			fmt.Println(podName)
			time.Sleep(time.Second * 2)			
		}
	}
}


func main(){

	Wait_goFunc.Add(1)
	
	go check_ImagePullBackOff()
	
	go check_OOMKilled()
	//go check_Running()
	//go check_Complete()
	//go check_Pending()
	//go check_Terminating()
	//go check_gpuAllocation()	

	go check_Error()
	//go check_Complete_logBackup()

	Wait_goFunc.Wait()
}
