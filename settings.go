package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

// Settings represents the application settings.
type Settings struct {
	Version    string     `json:"version"`
	Throttling Throttling `json:"throttling"`
	Links      []Link     `json:"links"`
}

// Throttling represents the throttling settings.
type Throttling struct {
	Today time.Time `json:"today"`
	Calls int       `json:"calls"`
}

// Link represents a URI we can open using the '-o' parameter.
type Link struct {
	ID  string `json:"id"`
	URI string `json:"uri"`
}

// LoadSettings loads the settings from the settings file, or returns
// the default settings if no settings file exists.
func LoadSettings() (Settings, error) {

	var s Settings

	exists, err := exists(getSettingsPath())
	if err != nil {
		return s, err
	}
	if !exists {
		return createDefaultSettings(), nil
	}

	raw, err := ioutil.ReadFile(getSettingsPath())
	if err != nil {
		return s, err
	}

	json.Unmarshal(raw, &s)
	return s, err
}

func getSettingsPath() string {
	return os.Getenv("HOME") + "/.google-it.json"
}

func createDefaultSettings() Settings {
	return Settings{
		Version: "1.0",
		Throttling: Throttling{
			Today: time.Now(),
			Calls: 0,
		},
		Links: []Link{},
	}
}

// SaveSettings saves the application settings to file.
func SaveSettings(settings Settings) error {
	settingsJSON, _ := json.Marshal(settings)
	err := ioutil.WriteFile(getSettingsPath(), settingsJSON, 0644)
	return err
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
