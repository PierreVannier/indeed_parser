# indeed_parser

This is a small crawler written in Golang. The idea is to crawl urls taken from a received POST json payload.
```
POST /get_jobs HTTP/1.0
Content-Type: application/json

{

“urls”: [
“http://www.indeed.com/viewjob?jk=61c98a0aa32a191b”,
“http://www.indeed.com/viewjob?jk=27342900632b9796”,
…
]
}

```


// - Compile and launch the server (port 8080 by default)
// - Create a file (url.json) with a content similar to the following json :
// {
// 	"urls": [
// 	"http://www.indeed.com/viewjob?jk=61c98a0aa32a191b",
// 	"http://www.indeed.com/viewjob?jk=27342900632b9796",
// 	"http://www.indeed.com/viewjob?jk=65bc65cb685bbf3e"
// 	]
// }
// - Test and call the server with a call similar to :
// curl -H "Content-type:application/json" --data @urls.json http://localhost:8080