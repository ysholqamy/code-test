package event

import "fmt"

//eventsData describes the events that were captured for a single session
type eventsData struct {
	websiteURL         string
	resizeFrom         dimension
	resizeTo           dimension
	copyAndPaste       map[string]bool
	formCompletionTime int
}

// newEventsData takes a websiteURL and creates a new EventsData.
func newEventsData(url string) *eventsData {
	return &eventsData{websiteURL: url, copyAndPaste: make(map[string]bool, 0)}
}

type dimension struct {
	width  string
	height string
}

//newDimension takes dimension as map. i.e. map{"width":"10","height":"10"} and creates a new Dimension
func newDimension(values map[string]interface{}) dimension {
	return dimension{width: values["width"].(string), height: values["height"].(string)}
}

// populate takes an event raw data and populates EventsData according to eventType.
func (ev *eventsData) populate(event map[string]interface{}, sid string) Result {
	//TODO: guard against zero values returned from map.
	eventType := event["eventType"].(string)
	var message string
	switch eventType {

	case "timeTaken":
		ev.formCompletionTime = int(event["time"].(float64))
		message = fmt.Sprintf("Done capturing events for session %s, captured %+v", sid, ev)

	case "copyAndPaste":
		formId := event["formId"].(string)
		pasted := event["pasted"].(bool)

		ev.copyAndPaste[formId] = pasted
		message = fmt.Sprintf("Captured copyAndPaste event for session %s. state: %+v", sid, ev)

	case "resize":
		ev.resizeFrom = newDimension(event["resizeFrom"].(map[string]interface{}))
		ev.resizeTo = newDimension(event["resizeTo"].(map[string]interface{}))
		message = fmt.Sprintf("Captured resize event for session %s. state: %+v", sid, ev)

	default:
		err := fmt.Errorf("Missing eventType %+v while capturing for sessionId %s", event, sid)
		return Result{"", err}
	}
	return Result{message, nil}
}
