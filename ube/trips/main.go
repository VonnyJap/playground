package main

func main() {}

func minimumTime(time []int, totalTrips int) int64 {

	start := 1
	min, _ := MinMax(time)
	end := min * totalTrips

	for start <= end {
		mid := (start + end) / 2
		current := numTrips(time, mid)
		next := numTrips(time, mid+1)
		if current < totalTrips && next >= totalTrips {
			return int64(mid + 1)
		}
		if start == end {
			return int64(mid)
		}
		if current < totalTrips {
			start = mid
			continue
		}
		end = mid
	}
	return -1
}

func numTrips(s []int, t int) int {
	trips := 0
	for _, el := range s {
		trips += t / el
	}
	return trips
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

// class Solution:
//     def numTrips(self, time, t):
//         trips = 0
//         for el in time:
//             trips += t // el
//         return trips

//     def minimumTime(self, time: List[int], totalTrips: int) -> int:
//         beg = 1
//         end = min(time) * totalTrips
//         while beg <= end:
//             mid = (beg + end) // 2
//             currtrips = self.numTrips(time, mid)
//             nexttrips = self.numTrips(time, mid + 1)
//             if currtrips < totalTrips and nexttrips >= totalTrips:
//                 return mid + 1
//             elif beg == end:
//                 return mid
//             elif currtrips < totalTrips:
//                 beg = mid
//             else:
//                 end = mid
