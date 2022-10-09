package main

import "sort"

func main() {

}
func rankTeams(votes []string) string {

	if len(votes) == 0 {
		return votes[0]
	}

	teamRanks := make(map[byte][]int)
	for i := 0; i < len(votes); i++ {
		for j := 0; j < len(votes[i]); j++ {
			team := votes[i][j]
			if teamRanks[team] == nil {
				teamRanks[team] = make([]int, len(votes[0]))
			}
			teamRanks[team][j]++
		}
	}
	// make them an array so it can be sorted
	type pair struct {
		team  byte
		ranks []int
	}
	var teams []pair
	for team, ranks := range teamRanks {
		teams = append(teams, pair{team, ranks})
	}
	sort.Slice(teams, func(i, j int) bool {
		for k := 0; k < len(votes[0]); k++ {
			if teams[i].ranks[k] > teams[j].ranks[k] {
				return true
			} else if teams[i].ranks[k] < teams[j].ranks[k] {
				return false
			}
		}
		// The rank is the same
		return teams[i].team < teams[j].team
	})
	res := make([]byte, len(teams))
	for i, p := range teams {
		res[i] = p.team
	}
	return string(res)
}
