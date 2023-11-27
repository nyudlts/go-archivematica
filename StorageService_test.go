package go_am

import (
	"testing"
)

func TestStorageServiceEndpoints(t *testing.T) {

	ssClient, err := NewAMClient(config, 20)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ssClient)

	var pipelineUUID string
	t.Run("get pipelines", func(t *testing.T) {
		pipelines, err := ssClient.GetPipelines()
		if err != nil {
			panic(err)
		}
		if len(pipelines.Objects) > 0 {
			pipelineUUID = pipelines.Objects[0].UUID.String()
		} else {
			t.Error("No piplines returned")
		}
	})

	t.Run("get a pipeline", func(t *testing.T) {
		pipeline, err := ssClient.GetPipeline(pipelineUUID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", pipeline)
	})

	var spaceUUID string
	t.Run("get spaces", func(t *testing.T) {
		spaces, err := ssClient.GetSpaces()
		if err != nil {
			t.Error(err)
		}
		if len(spaces.Objects) > 0 {
			spaceUUID = spaces.Objects[0].UUID.String()
		} else {
			t.Error("No spaces returned")
		}

		t.Logf("%v\n", spaceUUID)
	})

	t.Run("get space", func(t *testing.T) {
		space, err := ssClient.GetSpace(spaceUUID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", space)
	})

	var locationUUID string
	t.Run("get locations", func(t *testing.T) {
		locations, err := ssClient.GetLocations()
		if err != nil {
			t.Error(err)
		}
		if len(locations.Objects) > 0 {
			locationUUID = locations.Objects[0].UUID.String()
			t.Log(locations.Objects[0].Description)
		} else {
			t.Error("No locations returned")
		}
	})
	t.Run("get location", func(t *testing.T) {
		location, err := ssClient.GetLocation(locationUUID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", location)
	})

	t.Run("browse a location", func(t *testing.T) {
		locationBrowser, err := ssClient.BrowseLocation("ae15ce6f-40e7-4ed7-b8ce-11c65a7643e9")
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", locationBrowser)
	})

	t.Run("get a location by name", func(t *testing.T) {
		locationName := "staging ingest"
		location, err := ssClient.GetLocationByName(locationName)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", location)
	})
}
