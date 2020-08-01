package main

import "fmt"

func main() {
	// rooms := [][]int{{1}, {2}, {3}, {}, {4}} // false
	rooms := [][]int{{1}, {2}, {3}, {}} // true
	// rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}} // false
	// rooms := [][]int{{2, 3}, {}, {2}, {1, 3, 1}} // true

	fmt.Println(canVisitAllRooms(rooms))
}

func canVisitAllRooms(rooms [][]int) bool {
	var roomsUnlocked = make([]bool, len(rooms), len(rooms))
	roomsUnlocked[0] = true // room 0 starts unlocked

	numUnlocked := canVisitAllRoomsRec(rooms, roomsUnlocked, 0, 1)

	if numUnlocked == len(rooms) {
		return true
	}
	return false
}

func canVisitAllRoomsRec(rooms [][]int, roomsUnlocked []bool, roomNum int, numUnlocked int) int {

	// fmt.Println("roomNum:", roomNum)

	for _, key := range rooms[roomNum] {
		if roomsUnlocked[key] {
			continue
		}

		roomsUnlocked[key] = true

		// fmt.Println("roomsUnlocked:", roomsUnlocked)

		numUnlocked++

		numUnlocked = canVisitAllRoomsRec(rooms, roomsUnlocked, key, numUnlocked)
	}

	return numUnlocked

	// fmt.Println("numUnlocked:", numUnlocked)
}
