package main

import (
	"encoding/json"
	"fmt"
	"github.com/dyatlov/go-opengraph/opengraph"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type MyError struct {
	At     time.Time
	Reason string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.At, e.Reason)
}

type URLlist struct {
	URLs []string `json:"urls"`
}

type Job struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Company  string `json:"company"`
	URL      string `json:"url"`
}

func getOpenGraph(url string) (*Job, error) {
	resp, err := http.Get(url)

	if err != nil {
		return &Job{}, MyError{time.Now(), "Problem with network."}
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Job{}, MyError{time.Now(), "Problem fetching body response."}
	}

	og := opengraph.NewOpenGraph()
	err = og.ProcessHTML(strings.NewReader(string(body)))

	if err != nil {
		return &Job{}, MyError{time.Now(), "Problem fetching og:title property."}
	}

	if og.Title != "" {
		job := strings.Split(og.Title, " - ")
		fmt.Println("Analyzing url : ", url)
		if len(job) < 3 {
			return &Job{}, MyError{time.Now(), fmt.Sprintf("Problem with title information %s", url)}
		}
		return &Job{job[0], job[2], job[1], url}, nil
	} else {
		return &Job{}, MyError{time.Now(), fmt.Sprintf("Problem fetching title information %s", url)}
	}
	return &Job{}, MyError{time.Now(), "Unknown error."}
}

func writeJson(jobs []Job, rw http.ResponseWriter) {
	js, err := json.Marshal(jobs)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func postRequest(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var t URLlist
	if err := decoder.Decode(&t); err != nil {
		panic(err)
	}

	c := make(chan Job)
	jobs := make([]Job, 0)

	for _, url := range t.URLs {
		go func(url string) {
			job, err := getOpenGraph(url)
			if err != nil {
				fmt.Println(err)
			}
			c <- *job
		}(url)
	}

	for i := 0; i < len(t.URLs); i++ {
		jobs = append(jobs, <-c)
	}
	writeJson(jobs, rw)
}

func main() {
	fmt.Println("Server is running...")
	http.HandleFunc("/get_jobs", postRequest)
	http.ListenAndServe(":80", nil)
}
