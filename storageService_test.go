package go_archivematica

import (
	"testing"
)

func TestStorageServiceEndpoints(t *testing.T) {

	ssClient, err := NewAMClient(config, 20)
	if err != nil {
		t.Fatal(err)
	}

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
		_, err := ssClient.GetPipeline(pipelineUUID)
		if err != nil {
			t.Error(err)
		}
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

	})

	t.Run("get space", func(t *testing.T) {
		_, err := ssClient.GetSpace(spaceUUID)
		if err != nil {
			t.Error(err)
		}
	})

	var locationRef Location

	t.Run("get locations", func(t *testing.T) {
		locations, err := ssClient.GetLocations()
		if err != nil {
			t.Error(err)
		}
		if len(locations.Objects) > 0 {
			locationRef = locations.Objects[0]
		} else {
			t.Error("No locations returned")
		}
	})

	t.Run("get location", func(t *testing.T) {
		_, err := ssClient.GetLocation(locationRef.UUID.String())
		if err != nil {
			t.Error(err)
		}

	})

	t.Run("browse a location", func(t *testing.T) {
		_, err := ssClient.BrowseLocation(locationRef.UUID.String())
		if err != nil {
			t.Error(err)
		}

	})

	t.Run("get a location by name", func(t *testing.T) {
		_, err := ssClient.GetLocationByName(locationRef.Description)
		if err != nil {
			t.Error(err)
		}
	})

}
