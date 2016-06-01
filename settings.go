package main

import (
    "encoding/json"
    "io/ioutil"
    "time"
    "os"
)


type Settings struct {
  Version string `json:"version"`
  Throttling Throttling `json:"throttling"`
  Links []Link `json:"links"`
}

type Throttling struct {
  Today time.Time `json:"today"`
  Calls int `json:"calls"`
}

type Link struct {
  Id string `json:"id"`
  Uri string `json:"uri"`
}

func LoadSettings() (Settings, error) {

  var s Settings

  exists, err := exists(GetSettingsPath())
  if err != nil {
    return s, err
  }
  if !exists {
    return CreateDefaultSettings(), nil
  }

  raw, err := ioutil.ReadFile(GetSettingsPath())
  if err != nil {
    return s, err
  }

  json.Unmarshal(raw, &s)
  return s, err
}

func GetSettingsPath() string {
  return os.Getenv("HOME") + "/.google-it.json"
}

func CreateDefaultSettings() Settings {
  return Settings{
        Version: "1.0",
        Throttling: Throttling{
            Today: time.Now(),
            Calls: 0,
        },
        Links: []Link{},
    }
}

func SaveSettings(settings Settings) error {
    settingsJson, _ := json.Marshal(settings)
    err := ioutil.WriteFile(GetSettingsPath(), settingsJson, 0644)
    return err
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}