package application

import (
	"context"
	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"log"
	"time"
)

func StressTestWorker(client worker.JobClient, job entities.Job) {
	//get Task variable and headers
	jobKey := job.GetKey()
	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		failJob(client, job)
		return
	}

	//Logic
	time.Sleep(100 * time.Millisecond)
	//set Task variable

	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}
	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}
	log.Println("Successfully completed job")
}
