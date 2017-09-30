package dht

import "github.com/sirupsen/logrus"

// distEntry is used to hold the distance between nodes
type distEntry struct {
	id   string
	dist []int
}

// Xor gets to byte arrays and returns and array of integers with the xor
// for between the two equivalent bytes
func xor(a, b []byte) []int {
	compA := []byte{}
	compB := []byte{}
	res := []int{}

	lenA := len(a)
	lenB := len(b)

	// Make both byte arrays have the same size
	if lenA > lenB {
		compA = a
		compB = make([]byte, lenA)
		// Need to leave leftmost bytes empty in order compare
		// the equivalent bytes
		copy(compB[lenA-lenB:], b)
	} else {
		compB = b
		compA = make([]byte, lenB)
		copy(compA[lenB-lenA:], a)
	}

	for i := range compA {
		res = append(res, int(compA[i]^compB[i]))
	}

	return res
}

// lessIntArr compares two int array return true if a less than b
func lessIntArr(a, b []int) bool {
	for i := range a {
		if a[i] > b[i] {
			return false
		}
		if a[i] < b[i] {
			return true
		}
	}

	return true
}

func comparePeers(a, b, targetPeer string) string {
	logrus.Infof("A: %s B: %s TARGET: %s", a, b, targetPeer)
	distA := xor([]byte(a), []byte(targetPeer))
	distB := xor([]byte(b), []byte(targetPeer))
	if lessIntArr(distA, distB) {
		return a
	}
	return b
}