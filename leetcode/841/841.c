#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_NUM_OF_ROOMS 1000

bool canVisitAllRooms(int **rooms, int roomsSize, int *roomsColSize)
{
	// Create a roomsUnlocked array, initialized with 'false'.
	bool roomsUnlocked[MAX_NUM_OF_ROOMS] = {true};

	// Create a counter that tracks the number of unlocked rooms. Start with 1
	// because rooms[0] is always unlocked.
	int numRoomsUnlocked = 1;

	// Create 'queue', a block of memory of int pointers, initialized with 0.
	int queue[MAX_NUM_OF_ROOMS] = {0};
	int front = 0;
	int end = 0;
	int numsInQueue = 1;

	// For each number in the queue, for each key in rooms[i], check and update.
	while (numsInQueue != 0)
	{
		int room = queue[front];
		for (int i = 0; i < roomsColSize[room]; i++)
		{
			int key = rooms[room][i];
			if (!roomsUnlocked[key])
			{
				end++;
				queue[end] = key;
				numsInQueue++;
				roomsUnlocked[key] = true;
				numRoomsUnlocked++;
			}
		}
		front++;
		numsInQueue--;
	}

	// Once the queue is empty, if numRoomsUnlocked == roomsSize, return true.
	if (numRoomsUnlocked == roomsSize)
	{
		return true;
	}

	return false;
}

int main()
{
	int roomsSize = 4;

	int roomsColSize[] = {2, 3, 1, 1};

	// First, allocate memory for rooms, then allocate memory for each room in
	// rooms.
	int **rooms = (int **)malloc(roomsSize * sizeof(int *));
	for (int i = 0; i < roomsSize; i++)
	{
		rooms[i] = (int *)malloc(roomsColSize[i] * sizeof(int));
	}
	rooms[0][0] = 1;
	rooms[0][1] = 3;
	rooms[1][0] = 3;
	rooms[1][1] = 0;
	rooms[1][2] = 1;
	rooms[2][0] = 2;
	rooms[3][0] = 0;

	canVisitAllRooms(rooms, roomsSize, roomsColSize);

	free(roomsColSize);

	for (int i = 0; i < roomsSize; i++)
	{
		free(rooms[i]);
	}
	free(rooms);

	return 0;
}