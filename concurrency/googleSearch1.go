package main

import (
	"fmt"
	"math/rand"
	"time"
)

type result string

type search func(query string) result

func fakeSearch(searchType string) search {
	return func(query string) result {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		return result(fmt.Sprintf("[%s result for %s]\n", searchType, query))
	}
}

var (
	web   = fakeSearch("Web")
	image = fakeSearch("Image")
	video = fakeSearch("Video")
)

func google1(query string) []result {
	var results []result
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
	return results
}

func searchGoogle(version string) {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	var results []result
	switch version {
	case "1.0":
		results = google1("Golang")
	case "2.0":
		results = google2("Golang")
	case "3.0":
		results = google3("Golang")
	case "4.0":
		results = google4("Golang")
	default:
		results = google1("Golang")
	}

	elapsedTime := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsedTime)
}

func google2(query string) []result {
	c := make(chan result)
	go func() { c <- web(query) }()
	go func() { c <- image(query) }()
	go func() { c <- video(query) }()

	var results []result
	for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}
	return results
}

func google3(query string) []result {
	c := make(chan result)
	go func() { c <- web(query) }()
	go func() { c <- image(query) }()
	go func() { c <- video(query) }()

	var results []result
	timeout := time.After(200 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			return results
		}
	}
	return results
}

func firstResult(query string, replicas ...search) result {
	c := make(chan result)
	replicaSearch := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go replicaSearch(i)
	}
	return <-c
}

var (
	web1   = fakeSearch("Web1")
	web2   = fakeSearch("Web2")
	image1 = fakeSearch("Image1")
	image2 = fakeSearch("Image2")
	video1 = fakeSearch("Video1")
	video2 = fakeSearch("Video2")
)

func google4(query string) []result {
	c := make(chan result)
	go func() { c <- firstResult(query, web1, web2) }()
	go func() { c <- firstResult(query, image1, image2) }()
	go func() { c <- firstResult(query, video1, video2) }()

	var results []result
	timeout := time.After(200 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			return results
		}
	}
	return results
}
