package go_archivematica

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestStorageServicePipelines(t *testing.T) {

	var pipelineUUID uuid.UUID

	t.Run("TestGetPipelines", func(t *testing.T) {
		pipelines, err := client.GetPipelines()
		if err != nil {
			t.Errorf("Error getting pipelines: %v", err)
		}

		for _, pipelineObj := range pipelines.Objects {
			fmt.Printf("%+v\n", pipelineObj)
		}

		pipelineUUID = pipelines.Objects[0].UUID
	})

	t.Run("TestGetPipeline", func(t *testing.T) {
		pipeline, err := client.GetPipeline(pipelineUUID.String())
		if err != nil {
			t.Errorf("Error getting pipeline: %v", err)
		}

		fmt.Printf("%+v\n", pipeline)
	})

	//Creating a pipline is probably not something we want to do in a test, but here is the code if we wanted to do it.
	/*
		t.Skip("TestCreatePipeline", func(t *testing.T) {
			createPipelineRequest := CreatePipelineRequest{
				Description:            "Test Pipeline",
				CreateDefaultLocations: true,
			}

			err := client.CreatePipeline(createPipelineRequest)
			if err != nil {
				t.Errorf("Error creating pipeline: %v", err)
			}
		})
	*/
}
