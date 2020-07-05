package exercise

// MyReader https://tour.go-zh.org/methods/22
type MyReader struct {
	
}

// Reader Reader implementor of MyReader
func (r MyReader) Read(p []byte) (int ,error) {

	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}

	return len(p), nil
}