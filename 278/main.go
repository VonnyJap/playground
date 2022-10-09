package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func main() {

}

func firstBadVersion(n int) int {
	// make a call from middle to n

	start := 1
	end := n
	bad := n

	for start <= end {

		middle := (start + end) / 2
		if isBadVersion(middle) {
			if middle < bad {
				bad = middle
			}
			end = middle - 1
			continue
		}
		start = middle + 1
	}
	return bad
}
