package main

import (
  "fmt"
  "os"
)

func main() {
  eventType := os.Getenv("GITHUB_EVENT_NAME")
	var eventData []byte
	if path := os.Getenv("GITHUB_EVENT_PATH"); path != "" {
		eventData, err = os.ReadFile(path)
		if err != nil {
			fail(fmt.Errorf("failed to read event data: %w", err))
		}
	} else {
		fail(fmt.Errorf("event data is required"))
	}
  fmt.Println(eventType, string(eventData))
}

func fail(err error) {
  fmt.Fprintf(os.Stderr, "Action failed: %s\n", err)
  os.Exit(1)
}
