package go_archivematica

import "testing"

func TestLocations(t *testing.T) {

	var err error
	client, err = NewAMClient(config, environment, 20)
	if err != nil {
		t.Error(err)
	}

	var locationUUID string
	t.Run("Test Get Locations", func(t *testing.T) {
		locations, err := client.GetLocations()
		if err != nil {
			t.Errorf("GetLocations() error: %v", err)
		}
		locationUUID = locations.Objects[0].UUID.String()
		t.Log(locations)
	})

	t.Run("GetLocation", func(t *testing.T) {
		location, err := client.GetLocation(locationUUID)
		if err != nil {
			t.Errorf("GetLocation() error: %v", err)
		}
		t.Log(location)
	})

	t.Run("Browse Location", func(t *testing.T) {
		browser, err := client.BrowseLocation(locationUUID)
		if err != nil {
			t.Errorf("BrowseLocation() error: %v", err)
		}

		t.Log(browser)
	})
}
