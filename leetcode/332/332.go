package main

import (
	"fmt"
	"sort"
)

func insertElem(i int, x string, y []string) []string {

	if i < 0 {
		return append([]string{x}, y...)
	}

	y = append(y[:i+1], y[i:]...)
	y[i] = x
	return y
}

func createItinerary(i int, itinerary []string, allTickets map[string][]string) []string {

	for j := 0; j < len(allTickets[itinerary[i-1]]); j++ {

		// Pick one arrival given the departure (represented by itinerary[i-1]).
		// The arrival must not be an empty string and must be valid.
		// Unless the arrival is the last index value of itinerary, an arrival
		// is valid if it as a departure has valid arrivals.

		if itinerary[i] == "" && allTickets[itinerary[i-1]][j] != "" {
			if i == len(itinerary)-1 {
				itinerary[i] = allTickets[itinerary[i-1]][j]
				return itinerary
			}

			if _, ok := allTickets[allTickets[itinerary[i-1]][j]]; ok {

				itinerary[i] = allTickets[itinerary[i-1]][j]
				tmp := itinerary[i]
				allTickets[itinerary[i-1]][j] = ""

				itinerary = createItinerary(i+1, itinerary, allTickets)

				if i != len(itinerary)-1 && itinerary[i+1] == "" {
					itinerary[i] = ""
					allTickets[itinerary[i-1]][j] = tmp
				}
			}
		}
	}

	return itinerary
}

func findItinerary(tickets [][]string) []string {
	itinerary := make([]string, len(tickets)+1)
	itinerary[0] = "JFK"
	allTickets := make(map[string][]string)

	// Create a map where the key is the departure airport and the value is
	// a slice with the arrival airport(s).
	for i := 0; i < len(tickets); i++ {
		currDeparture := tickets[i][0]
		currArrival := tickets[i][1]
		if _, ok := allTickets[currDeparture]; ok {
			allTickets[currDeparture] = append(allTickets[currDeparture], currArrival)
		} else {
			allTickets[currDeparture] = []string{currArrival}
		}
	}

	for _, v := range allTickets {
		sort.Strings(v)
	}

	return createItinerary(1, itinerary, allTickets)
}

func main() {
	// tickets := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	// tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	// tickets := [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}
	// tickets := [][]string{{"EZE", "TIA"}, {"EZE", "AXA"}, {"AUA", "EZE"}, {"EZE", "JFK"}, {"JFK", "ANU"}, {"JFK", "ANU"}, {"AXA", "TIA"}, {"JFK", "AUA"}, {"TIA", "JFK"}, {"ANU", "EZE"}, {"ANU", "EZE"}, {"TIA", "AUA"}}

	tickets := [][]string{{"EZE", "TIA"}, {"EZE", "HBA"}, {"AXA", "TIA"}, {"JFK", "AXA"}, {"ANU", "JFK"}, {"ADL", "ANU"}, {"TIA", "AUA"}, {"ANU", "AUA"}, {"ADL", "EZE"}, {"ADL", "EZE"}, {"EZE", "ADL"}, {"AXA", "EZE"}, {"AUA", "AXA"}, {"JFK", "AXA"}, {"AXA", "AUA"}, {"AUA", "ADL"}, {"ANU", "EZE"}, {"TIA", "ADL"}, {"EZE", "ANU"}, {"AUA", "ANU"}}
	fmt.Println(findItinerary(tickets))
}
