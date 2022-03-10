package native

import "fmt"

// go wasm 中不能调用的:
// import "fmt"

func Fibonacci(in int32) int32 {
	if in <= 1 {
		return in
	}
	return Fibonacci(in-1) + Fibonacci(in-2)
}

func RequestHTTP() int32 {
	return 0
	// fmt.Println("RequestHTTP")
}

// current not support in WASI.
/*
func RequestHTTP() {
	httpTestURL := `https://www.baidu.com`

	resp, err := http.Get(httpTestURL)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
*/

func FileIO() int32 {
	return 0
	// fmt.Println("FileIO")
}

// current not support in WASI.
/*
func FileIO() error {
	dir, err := ioutil.TempDir("", "test-*")
	if err != nil {
		return err
	}
	defer os.Remove(dir)

	fmt.Println(dir)

	f, err := ioutil.TempFile(dir, "tmp-*")
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	fmt.Println(f.Name())

	writeN, err := f.WriteAt([]byte("test-content"), 0)
	if err != nil {
		return err
	}

	readDest := make([]byte, writeN)
	readN, err := f.ReadAt(readDest, 0)
	if err != nil {
		return err
	}
	if writeN != readN {
		return fmt.Errorf("read length is %d, written length is %d", readN, writeN)
	}

	return nil
}
*/

func MultiThreads(num int32) int32 {
	return 0
	// fmt.Println("MultiThreads")
}

// current not support in WASI.
/*
func MultiThreads(num int32) {
	g := new(errgroup.Group)

	for i := int32(0); i < num; i++ {
		g.Go(func() error {
			Fibonacci(30)
			return nil
		})
	}

	g.Wait()
}
*/

// BytesTest test for byte slice args and returns.
func BytesTest(in []byte) []byte {
	println("-n-|", fmt.Sprint(in), "|---")
	return append(in, '-', '-', '-')
}

// String test for string args and returns.
func StringTest(in string) string {
	println("-n-|", in, "|---")
	return in + "---"
}

// InterfaceTest test for interface{} args and returns.
func InterfaceTest(in interface{}) interface{} {
	println(in)
	return fmt.Sprintf("%v -- %s", in, "---")
}

// ErrTest test for err args and returns.
func ErrTest(in error) error {
	println(in)
	return fmt.Errorf("%v---", in)
}
