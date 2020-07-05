package exercise

import "io"

// Rot13Reader https://en.wikipedia.org/wiki/ROT13
type Rot13Reader struct {
	R io.Reader
}

// Read implementor
func (it Rot13Reader) Read(p []byte) (int, error) {

	i, err := it.R.Read(p)

	if err == io.EOF {
		return i, err
	}

	for j := 0; j < i; j++ {
		if p[j] >= 'a' && p[j] <= 'z' {
			p[j] += 13
			if p[j] > 'z' {
				p[j] -= 26
				continue
			}
		}
		if p[j] >= 'A' && p[j] <= 'Z' {
			p[j] += 13
			if p[j] > 'Z' {
				p[j] -= 26
			}
		}
	}

	return i, err
}
