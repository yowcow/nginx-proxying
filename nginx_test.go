package nginx

import (
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	ServerAddr = "http://127.0.0.1:8888"
)

func newRequest(path string) (*http.Client, *http.Request) {
	client := new(http.Client)

	req, _ := http.NewRequest("GET", ServerAddr+path, nil)
	req.Host = "myhost"

	return client, req
}

func Test_Root(t *testing.T) {
	client, req := newRequest("/")
	resp, err := client.Do(req)

	if err != nil {
		t.Error("expected nil but got", err)
	}

	if resp.StatusCode != 200 {
		t.Error("expected 200 but got", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error("expected nil but got", err)
	}

	expected := "Server Addr: :5001\nRequested URL: /\n"

	if string(body) != expected {
		t.Error("unexpected resp body")
	}
}

func Test_Foo(t *testing.T) {
	client, req := newRequest("/foo")
	resp, err := client.Do(req)

	if err != nil {
		t.Error("expected nil but got", err)
	}

	if resp.StatusCode != 200 {
		t.Error("expected 200 but got", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error("expected nil but got", err)
	}

	expected := "Server Addr: :5001\nRequested URL: /foo\n"

	if string(body) != expected {
		t.Error("unexpected resp body")
	}
}

func Test_Foo_Bar(t *testing.T) {
	client, req := newRequest("/foo/bar")
	resp, err := client.Do(req)

	if err != nil {
		t.Error("expected nil but got", err)
	}

	if resp.StatusCode != 200 {
		t.Error("expected 200 but got", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error("expected nil but got", err)
	}

	expected := "Server Addr: :5002\nRequested URL: /foo/bar\n"

	if string(body) != expected {
		t.Error("unexpected resp body")
	}
}

func Test_Foo_Bar_Buz(t *testing.T) {
	client, req := newRequest("/foo/bar/buz")
	resp, err := client.Do(req)

	if err != nil {
		t.Error("expected nil but got", err)
	}

	if resp.StatusCode != 200 {
		t.Error("expected 200 but got", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error("expected nil but got", err)
	}

	expected := "Server Addr: :5002\nRequested URL: /foo/bar/buz\n"

	if string(body) != expected {
		t.Error("unexpected resp body")
	}
}
