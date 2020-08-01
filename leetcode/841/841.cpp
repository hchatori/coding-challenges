#include <iostream>
#include <vector>


class Solution {
public:
	int canVisitAllRoomsRec(std::vector<std::vector<int>>& rooms, std::vector<bool>& roomsUnlocked, int roomNum, int numUnlocked) {
		for (auto& key : rooms[roomNum]) {
			if (roomsUnlocked[key]) {
				continue;
			}

			roomsUnlocked[key] = true;

			numUnlocked++;

			numUnlocked = canVisitAllRoomsRec(rooms, roomsUnlocked, key, numUnlocked);
		}

		return numUnlocked;
	}

    bool canVisitAllRooms(std::vector<std::vector<int>>& rooms) {
		std::vector<bool> roomsUnlocked (rooms.size(), false);
		roomsUnlocked[0] = true; // room 0 starts unlocked

		int numUnlocked = canVisitAllRoomsRec(rooms, roomsUnlocked, 0, 1);

		if (numUnlocked == rooms.size()) {
			return true;
		}

        return false;
    }
};

int main() {
	Solution s;
	std::vector<std::vector<int>> rooms {{1}, {2}, {3}, {}};
	std::cout << s.canVisitAllRooms(rooms) << std::endl;
	return 0;
}