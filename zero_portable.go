// +build !amd64

package zero

func Zero(a []byte) bool {
	for _, c := range a {
		if c != 0 {
			return false
		}
	}
	return true
}
