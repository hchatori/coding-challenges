package main

import "fmt"

type UndergroundSystem struct {
	in  map[int]Tuple
	out map[string][2]int
}

type Tuple struct {
	stationName string
	time        int
}

func main() {
	us := Constructor()
	us.CheckIn(10, "Leyton", 3)
	us.CheckOut(10, "Paradise", 8)
	fmt.Println(us.GetAverageTime("Leyton", "Paradise"))
	us.CheckIn(5, "Leyton", 10)
	us.CheckOut(5, "Paradise", 16)
	fmt.Println(us.GetAverageTime("Leyton", "Paradise"))
	us.CheckIn(2, "Leyton", 21)
	us.CheckOut(2, "Paradise", 30)
	fmt.Println(us.GetAverageTime("Leyton", "Paradise"))
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		in:  map[int]Tuple{},
		out: map[string][2]int{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.in[id] = Tuple{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	departureStation := this.in[id]
	stationToStationName := departureStation.stationName + stationName
	travelTime := t - departureStation.time
	var sumCounts [2]int

	if _, ok := this.out[stationToStationName]; !ok {
		sumCounts = [2]int{travelTime, 1}
	} else {
		sumCounts = this.out[stationToStationName]
		sumCounts[0] += travelTime
		sumCounts[1]++
	}

	this.out[stationToStationName] = sumCounts
	delete(this.in, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	stationToStationName := startStation + endStation
	sumCounts := this.out[stationToStationName]

	return float64(sumCounts[0]) / float64(sumCounts[1])
}
