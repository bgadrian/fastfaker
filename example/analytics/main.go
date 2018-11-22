package main

import (
	"fmt"
	"time"

	"github.com/bgadrian/fastfaker/example/analytics/producer"
)

func main() {
	/**
		This is a smaller version of the Kafka-analytics generator https://github.com/bgadrian/kafka-playground

		It creates a random amount of online users, and each of them are creating random events,
	simulating the activity on a E-commerce website.
	New events are easy to add, see: buildDispatcher

		By using FastFaker we can generate a large amount of fake events to test our analytics pipelines.
	*/

	collector := make(chan producer.Event)
	disp := buildDispatcher(25, collector, 15)
	disp.Start(time.Second * 2)

	for ev := range collector {
		props := ev.Properties()
		userId := props[producer.KeyUserID]
		delete(props, producer.KeyUserID)
		delete(props, "event")
		fmt.Printf("user %s: [%s]='%s'\n", userId, ev.Name(), props)
	}
}

func buildDispatcher(peak int, collector chan producer.Event, eventsPerMinPerUser int) *producer.Dispatcher {
	eventTemplates := []producer.Event{
		//for the supported {variable} see https://github.com/bgadrian/fastfaker/blob/master/TEMPLATE_VARIABLES.md
		producer.NewSimpleEvent("pageView", map[string]string{"page": "/{bs}/#/"}),
		producer.NewSimpleEvent("click", map[string]string{"element": "{hackerabbreviation}"}),
		producer.NewSimpleEvent("buy", map[string]string{
			"name":   "{carmaker}",
			"color":  " {safecolor}",
			"price":  "{uint8}.#",
			"amount": "#",
		}),
	}

	gen := producer.NewRandomEventGenerator(eventTemplates)

	queen := func() producer.User {
		//each user will generate a batch of events each random( minMs and maxMs)
		minMs := 30000
		maxMs := 140000
		user := producer.NewSimpleUser(eventsPerMinPerUser, gen.NewEvents, collector, minMs, maxMs)

		return user
	}

	return producer.NewDispatcher(queen)
}
