package main

import (
	"context"
	_ "embed"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/deemount/gobpmnLab/repository"
)

// var messageBroker repository.MessageBroker
var bpmnFactory repository.BpmnFactory
var bpmnEngine repository.BpmnEngine
var processInfo *repository.ProcessInfo
var instance repository.ProcessInstance

// init ...
func init() {
	log.Println("step 1 >> main.go: init factory")
	bpmnFactory = repository.NewBpmnFactory()
	log.Println("step 2 >> main.go: init engine")
	bpmnEngine = repository.NewBpmnEngine("gobpmnLab")
	log.Println("step 3 >> main.go: init process instance")
	instance = repository.NewProcessInstance(bpmnEngine)
	log.Println("main.go: Init done")
}

// main ...
func main() {

	start := time.Now()

	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer cpuFile.Close() // error handling omitted

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	logFile, _ := os.Create("logs/build.log")
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("main.go: bpmnFactory.Build")
	file, err := bpmnFactory.Build()
	if err != nil {
		panic(err)
	}

	log.Println("main.go: instance.GetProcessInfo")
	processInfo, err = instance.GetProcessInfo(context.Background(), file)
	if err != nil {
		panic(err)
	}

	log.Println("main.go: instance.Create")
	instanceInfo, err := instance.Create(processInfo.ProcessKey, nil)
	if err != nil {
		panic(err)
	}
	log.Println("main.go: instance.Run")
	if err = instance.Run(instanceInfo); err != nil {
		panic(err)
	} else {
		log.Println("main.go: do your stuff here")
		// Do your stuff here
	}

	log.Println("total time:", time.Since(start))

}
