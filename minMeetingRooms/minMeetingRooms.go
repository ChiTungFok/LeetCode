package minMeetingRooms

import "sort"

func minMeetingRooms(intervals [][]int) int {
	var rooms [][][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	for _, interval := range intervals {
		if len(rooms) > 1 {
			sort.Slice(rooms, func(i, j int) bool {
				return rooms[i][len(rooms[i])-1][1] > rooms[j][len(rooms[j])-1][1]
			})
		}
		var hasAdded bool
		for i := range rooms {
			if len(rooms[i]) == 0 {
				continue
			}
			if rooms[i][len(rooms[i])-1][1] <= interval[0] {
				rooms[i] = append(rooms[i], interval)
				hasAdded = true
				break
			}
		}
		if !hasAdded {
			var room [][]int
			room = append(room, interval)
			rooms = append(rooms, room)
		}
	}
	return len(rooms)
}
