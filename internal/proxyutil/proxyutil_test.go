package proxyutil

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func Example_redirect() {

	endHdl := func(rw http.ResponseWriter, req *http.Request) {
		b := req.Body
		d, err := io.ReadAll(b)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(d))
		rw.Write([]byte("Hello World"))
		rw.WriteHeader(http.StatusOK)
	}

	endServer := httptest.NewServer(http.HandlerFunc(endHdl))

	u, err := url.Parse(endServer.URL)
	if err != nil {
		fmt.Println(err)
	}

	r := Redirector{
		targetURL: u,
	}

	proxyServer := httptest.NewServer(r)

	reader := bytes.NewReader([]byte("abc"))

	req, err := http.NewRequest(http.MethodPost, proxyServer.URL, reader)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	fmt.Println(resp.StatusCode)

	// Output:
	// abc
	// Hello World
	// 200

}
