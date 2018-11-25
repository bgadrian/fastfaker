package producer

import (
	"strings"

	"github.com/bgadrian/fastfaker/faker"
	//nolint
)

// GetRandomUserProperties creates a map with predefined user properties
func GetRandomUserProperties() map[string]string {
	//it uses the fastFaker template system
	//for the supported {variable} see https://github.com/bgadrian/fastfaker/blob/master/TEMPLATE_VARIABLES.md
	result := map[string]string{
		//Update the README.md if you change the keys
		KeyUserID:   "uuid",
		"sessionId": "uuid",
		"userAgent": "useragent",
		"country":   "country",
	}

	//get random data
	fastFaker := faker.NewFastFaker()
	for k, v := range result {
		result[k], _ = fastFaker.TemplateCustom(v, "", "")
	}

	//to make the demo more user friendly we only keep the first part of the UUID
	result[KeyUserID] = strings.Split(result[KeyUserID], "-")[0]
	result["sessionId"] = strings.Split(result["sessionId"], "-")[0]

	return result
}

// NewRandomEventGenerator creates a Fake event generator based on a predefined set of patterns
// templateEvents must contains properties, and in their values FastFaker templates.
// ["price" => "###.##", "productCategory" => "{word}]
// each event generated will be based on these patterns, with fake data replacing the templates
func NewRandomEventGenerator(templateEvents []Event) *RandomEventGenerator {
	gen := &RandomEventGenerator{}
	gen.templateEvents = templateEvents
	return gen
}

// RandomEventGenerator creates random fake events
type RandomEventGenerator struct {
	templateEvents []Event
}

// NewEvents creates new fake events
func (g *RandomEventGenerator) NewEvents(count int) []Event {
	//for now we keep our logic simple, with no correlation between the events
	//we just select them at random

	result := make([]Event, 0, count)
	fastFaker := faker.NewFastFaker()

	for len(result) < count {
		patternIndex := fastFaker.Intn(len(g.templateEvents))
		patternEvent := g.templateEvents[patternIndex]
		patternProps := patternEvent.Properties()

		props := make(map[string]string, len(patternProps))
		for key, template := range patternProps {
			props[key] = fastFaker.Template(template)
		}

		result = append(result, NewSimpleEvent(patternEvent.Name(), props))
	}
	return result
}
