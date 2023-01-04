package main

func main() {

	// [['abb', 'abc'], ['xyz', 'xxz'], ['eee', 'eee'], ['aaaaac', 'aabaag']]

}

// 1. build the map by iterating over -> [b->c, y->x, a->b, c->g] -> this looks like adjacency list
// 2. once the map is built -> now need to create a sequence
// 3. to create a sequence -> we can do dfs
// 4. and with this we want to create dfs with memoization, storing the intermediate value

func findLongestMutation(input [][]string) string {

}

func buildMap(input [][]string) map[string]string {

	result := map[string]string{}

	for _, mutation := range input {
		start := mutation[0]
		end := mutation[1]
		for j, char := range start {
			// when it is not mutated
			if start[j] == end[j] {
				continue
			}
			if _, ok := result[string(char)]; ok {
				continue
			}
			result[string(char)] = string(end[j])
		}
	}
	return result
}
