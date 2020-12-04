package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	resp, err := http.Get("http://localhost:5100/")

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("TestGet", string(body))

}

func TestGetWithQueryString(t *testing.T) {

	resp, err := http.Get("http://localhost:5100/with_query_string?id=1000&page=1")

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("TestGetWithQueryString", string(body))

}

func TestPost(t *testing.T) {

	resp, err := http.Post("http://localhost:5100/post",
		"application/x-www-form-urlencoded",
		strings.NewReader("username=admin"))

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("TestPost", string(body))
}

func TestPostForm(t *testing.T) {

	resp, err := http.PostForm("http://localhost:5100/post",
		url.Values{"username": {"admin"}})

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("TestPostForm", string(body))

}

func TestGetPost1(t *testing.T) {

	var r http.Request
	r.ParseForm()
	r.Form.Add("username", "admin")

	queryStr := strings.TrimSpace(r.Form.Encode())

	request, err := http.NewRequest(http.MethodPost,
		"http://localhost:5100/get_post",
		strings.NewReader(queryStr))

	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("TestGetPost1", string(body))

}

func TestPostBinary(t *testing.T) {

	data := bytes.NewReader([]byte("Hello world"))

	request, err := http.NewRequest(http.MethodPost, "http://localhost:5100/post_binary", data)

	if err != nil {
		t.Fatal(err)
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("TestPostBinary", string(body))
}
