package exercise

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"golang.org/x/tour/tree"
	"golang.org/x/tour/wc"
)

// RunExercise 运行练习
func RunExercise() {
	wc.Test(WordCount)

	f := Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}

	fmt.Println()

	v := Vertex{X: 3, Y: 4}
	v.Scale(10)
	fmt.Printf("abs of %v, %v is %v \n", v.X, v.Y, v.Abs())

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, &ip)
	}

	b := Bird{Name: "Bird"}
	b.Fly()

	e := Eagle{Name: "Eagle"}
	e.Fly()

	fmt.Printf("Sqrt of %v is %v", 9, Sqrt(9))

	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
	reader.Validate(MyReader{})

	reader := strings.NewReader("你好，中国！")
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, buffer)
		fmt.Printf("b[:n] = %q\n", buffer[:n])
		if err == io.EOF {
			break
		}
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{R: s}
	io.Copy(os.Stdout, &r)

	m := Image{Width: 100, Height: 100}
	pic.ShowImage(&m)

	t1 := tree.New(1)
	t2 := tree.New(1)

	fmt.Println(Same(t1, t2))

	// fetcher 是填充后的 fakeFetcher。
	var fetcher = FakeFetcher{
		"https://golang.org/": &FakeResult{
			Body: "The Go Programming Language",
			Urls: []string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &FakeResult{
			Body: "Packages",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": &FakeResult{
			Body: "Package fmt",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": &FakeResult{
			Body: "Package os",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	}

	Crawl("https://golang.org/", 4, fetcher)

	time.Sleep(500 * time.Millisecond)
}
