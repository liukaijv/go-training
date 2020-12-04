package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.Method, request.URL.Path)
		writer.Write([]byte("Hello world"))
	})

	http.HandleFunc("/with_query_string", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.Method, request.URL.Path)
		query := request.URL.Query()

		queryMap := make(map[string]interface{})

		for k, v := range query {
			queryMap[k] = v[0]
		}

		content, err := json.Marshal(queryMap)

		if err != nil {
			log.Fatal(err)
		}

		writer.Write(content)
	})

	http.HandleFunc("/post", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.Method, request.URL.Path)
		request.ParseForm()

		formMap := make(map[string]interface{})

		for k, v := range request.Form {
			formMap[k] = v[0]
		}

		content, err := json.Marshal(formMap)

		if err != nil {
			log.Fatal(err)
		}

		writer.Write(content)
	})

	http.HandleFunc("/get_post", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.Method, request.URL.Path)
		request.ParseForm()

		formMap := make(map[string]interface{})

		for k, v := range request.Form {
			formMap[k] = v[0]
		}

		content, err := json.Marshal(formMap)

		if err != nil {
			log.Fatal(err)
		}

		writer.Write(content)

	})

	http.HandleFunc("/post_binary", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.Method, request.URL.Path)

		body, err := ioutil.ReadAll(request.Body)

		if err != nil {
			log.Fatal(err)
		}

		defer request.Body.Close()

		writer.Write(body)

	})

	http.ListenAndServe(":5100", nil)
}
