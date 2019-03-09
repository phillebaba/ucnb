package main

import (
	"log"
	"bytes"
	"errors"
	"encoding/json"
	"net/http"
	"net/url"
)

func parseOutputPlugin(pluginString string) (OutputPlugin, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(pluginString), &data)
	if err != nil {
		return nil, err
	}

	plugin_type := data["type"].(string)
	switch plugin_type {
		case "ifttt":
			value := data["value"].(map[string]interface{})
			output := Ifttt{eventName: value["event_name"].(string), apiKey: value["api_key"].(string)}
			return output, nil
		case "http":
			value := data["value"].(map[string]interface{})
			output := Http{endpoint: value["endpoint"].(string)}
			return output, nil
		default:
			return nil, errors.New("Output plugin type not recognized")
	}
}

type OutputPlugin interface {
	Send(message string) error
}

type Ifttt struct {
	eventName string
	apiKey string
}

func (i Ifttt) Send(message string) error {
	log.Println("Triggering IFTTT Webhook")

	url := "https://maker.ifttt.com/trigger/" + i.eventName + "/with/key/" + i.apiKey
	values := map[string]string{}
	values["value1"] = message

	body, err := json.Marshal(values)

	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewReader(body))

	if err != nil {
		return err
	}

	return nil
}

type Http struct {
	endpoint string
}

func (w Http) Send(message string) error {
	log.Println("Sending Post Request")

	values := map[string]string{}
	values["message"] = message

	_, err := http.PostForm(w.endpoint, url.Values{"message": {message}})

	if err != nil {
		return err
	}

	return nil
}
