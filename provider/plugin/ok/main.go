package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var testGlobalVal int32

func ModifyGlobalVal(delta int32) int32 {
	testGlobalVal += delta

	return testGlobalVal
}

func Fibonacci(in int32) int32 {
	if in <= 1 {
		return in
	}
	return Fibonacci(in-1) + Fibonacci(in-2)
}

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

	logrus.Infoln(string(bs))
}

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
