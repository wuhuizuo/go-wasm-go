package provider

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const httpTestURL = `https://httpbin.org/basic-auth`

func Fibonacci(in int32) int32 {
	if in <= 1 {
		return in
	}
	return Fibonacci(in-1) + Fibonacci(in-2)
}

func HTTPBasicAuth(username, password string) {
	reqURL := fmt.Sprintf("%s/%s/%s", httpTestURL, username, password)

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(req)
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
