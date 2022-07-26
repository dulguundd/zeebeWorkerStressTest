package application

import (
	"context"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"log"
)

var readyClose = make(chan struct{})

var JobTypeString = [...]string{
	"test-worker-1",
	"test-worker-2",
	"test-worker-3",
	"test-worker-4",
	"test-worker-5",
}

func Start() {
	environment := getEnvironment()
	zbClient, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         environment.zeebeConfig.zeebeAddress,
		UsePlaintextConnection: true,
	})
	if err != nil {
		panic(err)
	}

	jobWorkerType := zbClient.NewJobWorker().JobType(environment.serviceConfig.workerType)

	switch environment.serviceConfig.workerType {
	case JobTypeString[0]: //if workerType == "payment-service"
		jobWorker := jobWorkerType.Handler(StressTestWorker).Open()
		closeJobWorker(jobWorker, environment.serviceConfig.workerType)
	case JobTypeString[1]: //if workerType == "inventory-service"
		jobWorker := jobWorkerType.Handler(StressTestWorker).Open()
		closeJobWorker(jobWorker, environment.serviceConfig.workerType)
	case JobTypeString[2]:
		jobWorker := jobWorkerType.Handler(StressTestWorker).Open()
		closeJobWorker(jobWorker, environment.serviceConfig.workerType)
	case JobTypeString[3]:
		jobWorker := jobWorkerType.Handler(StressTestWorker).Open()
		closeJobWorker(jobWorker, environment.serviceConfig.workerType)
	case JobTypeString[4]:
		jobWorker := jobWorkerType.Handler(StressTestWorker).Open()
		closeJobWorker(jobWorker, environment.serviceConfig.workerType)
	default:
		fmt.Printf("No matched worker type! \n")
		return
	}
}

func closeJobWorker(jobWorker worker.JobWorker, workerType string) {
	fmt.Printf("Starting worker type: " + workerType + "\n")

	<-readyClose
	jobWorker.Close()
	jobWorker.AwaitClose()

}

func failJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())

	ctx := context.Background()
	_, err := client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send(ctx)
	if err != nil {
		panic(err)
	}
}
