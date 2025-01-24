package solutions

import (
	"encoding/json"
	"log"
	"math"
	"strings"
)

func Solution7() {

}

type Queue struct {
	items []any
}

func (q *Queue) Enqueue(value ...any) {
	if len(q.items) == 0 {
		q.items = []any{}
	}
	q.items = append(q.items, value...)
}

func (q *Queue) Dequeue() any {
	if len(q.items) == 0 {
		panic("Queue is empty")
	}
	value := q.items[0]
	q.items = q.items[1:] // Avoids reallocation by slicing
	return value
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

type Node struct {
	Source      string
	Destination string
	Price       float64
}

type Journey struct {
	Path       Node
	Step       int
	TotalPrice float64
	History    []string
}

func dump(v any) {
	raw, _ := json.MarshalIndent(v, "", "   ")
	log.Println(string(raw))
}

func main() {
	k := 3
	src := "JFK"
	dst := "LAX"

	// key city_dest|step
	currCheapest := float64(0)

	paths := []Node{
		Node{Source: "ATL", Destination: "SFO", Price: 400},
		Node{Source: "ORD", Destination: "LAX", Price: 200},
		Node{Source: "LAX", Destination: "DFW", Price: 80},
		Node{Source: "JFK", Destination: "HKG", Price: 800},
		Node{Source: "ATL", Destination: "ORD", Price: 90},
		Node{Source: "JFK", Destination: "LAX", Price: 500},
		Node{Source: "JFK", Destination: "DFW", Price: 850},
		Node{Source: "JFK", Destination: "ATL", Price: 150},
	}

	flight := map[string][]Node{}

	for _, n := range paths {
		if len(flight[n.Source]) == 0 {
			flight[n.Source] = []Node{}
		}

		flight[n.Source] = append(flight[n.Source], n)
	}

	queue := Queue{}

	// get the first path

	firstPath := flight[src]

	for _, v := range firstPath {
		queue.Enqueue(Journey{
			Path:       v,
			Step:       1,
			TotalPrice: v.Price,
			History:    []string{v.Source},
		})
	}

	for !queue.IsEmpty() {
		raw := queue.Dequeue()
		journey := raw.(Journey)
		step := journey.Step
		price := journey.TotalPrice

		if price == 0 { // initial price
			price = journey.Path.Price // taken from the flight
		}

		if journey.Path.Destination == dst && k <= journey.Step+1 {
			journey.History = append(journey.History, journey.Path.Destination)
			if currCheapest == 0 {
				currCheapest = price
			} else {
				currCheapest = math.Min(currCheapest, price)
			}
			log.Println(strings.Join(journey.History, "->"))
			log.Println("Cheapest Price:", price)
			log.Println("Total Steps: ", journey.Step)
			continue
		}

		nextPath := flight[journey.Path.Destination]
		for _, nPath := range nextPath {
			if currCheapest != 0 && price+nPath.Price >= currCheapest {
				log.Println("Dropped: ", strings.Join(journey.History, "->"), "price: ", price+nPath.Price)
				continue
			}

			queue.Enqueue(Journey{
				Path:       nPath,
				Step:       step + 1,
				TotalPrice: price + nPath.Price,
				History:    append(journey.History, nPath.Source),
			})
		}
	}
}
