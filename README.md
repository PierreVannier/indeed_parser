# indeed_parser

This is a small crawler written in Golang as an exercise. The idea is to crawl urls taken from a received POST json payload looking like the following. (Indeed website, fetch the job title, the location and the company).

```http
POST /get_jobs HTTP/1.0
Content-Type: application/json

{
  "urls": [
    "http://www.indeed.com/viewjob?jk=61c98a0aa32a191b",
    "http://www.indeed.com/viewjob?jk=27342900632b9796"
  ]
}

```

And return this kind of response.

```http
HTTP/1.1 200 OK
Date: Wed, 17 Feb 2016 01:45:49 GMT
Content-Type: application/json

[
  {
    "title": "Software Engineer job",
    "location": "San Francisco, CA",
    "company": "Braintree",
    "url": "http://www.indeed.com/viewjob?jk=27342900632b9796"
  },
  {
    "title": "Software Engineer job",
    "location": "Tigard, OR",
    "company": "Zevez Corporation",
    "url": "http://www.indeed.com/viewjob?jk=61c98a0aa32a191b"
  }
]
```
Instead of parsing the HTML response, I based my fetch on the [opengraph](http://ogp.me/) og:title. 

## Install
```
got get github.com/dyatlov/go-opengraph/opengraph
go build parser.go && sudo ./parser
```
I created a json file for ease of testing called `urls.json`. 
It contains some URLs to fetch.

You can launch a curl (from another terminal):
```
curl -H "Content-type:application/json" --data @urls.json http://localhost/get_jobs
```
This should send you back the correct response.

## Side note
There is definitely room for improvements. Error management, timeout of server...etc.

I spent 4 hours on this program; noticeably, first time I used goroutines and channels, concurrency. As a matter of fact I had to work a bit on this side of the program :-)

Need more information, feel free to ping me at pierre.vannier at gmail.



