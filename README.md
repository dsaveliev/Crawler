Simple web crawler
========================================

External dependencies
-------------------
* https://godoc.org/golang.org/x/net/html
* https://github.com/buaazp/fasthttprouter
* https://github.com/valyala/fasthttp
* https://github.com/h2non/gock

Running the tests
-------------
Shell script **test.sh**:

```bash
go test -v ./models ./server ./pool
```

Running the server
--------------

Shell script **run.sh** with *optional* flags:
* **-port** TCP port.
* **-debug** verbose logging.

```bash
go run ./main.go -port=8081 -debug
```

Implementation notes
-----------------------

**time.Ticker** was used to implement throttling per domain.
Each domain has a corresponding **<-chan time.Time** channel, and this channel is used
as a semaphore to control the download speed amongst all goroutines for this domain.

For the sake of simplicity I retrun **CSV** as **text/csv**, others responses as **text/plain**.
Default server port is **8080**.

API description
------------

#### Upload file with link list

```bash
$ curl -X POST --data-binary "@urls.txt" localhost:8080/tasks
...
< HTTP/1.1 201 Created
< Date: Sun, 08 Oct 2017 13:52:02 GMT
< Content-Length: 1
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
2⏎
```

#### Get working task

```bash
curl -v "localhost:8080/tasks/3"
...
< HTTP/1.1 204 No Content
< Date: Sun, 08 Oct 2017 13:54:47 GMT
<
* Connection #0 to host localhost left intact
```

#### Get not existing/deleted task

```bash
 curl -v "localhost:8080/tasks/1"
 ...
< HTTP/1.1 404 Not Found
< Date: Sun, 08 Oct 2017 13:59:42 GMT
< Content-Length: 15
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Task not found.⏎
```

#### Get completed task

```bash
curl -v "localhost:8080/tasks/1"
...
< HTTP/1.1 200 OK
< Content-Type: text/csv; charset=UTF-8
< Date: Sun, 08 Oct 2017 13:55:45 GMT
< Transfer-Encoding: chunked
<
StatusCode	URL	Title	Description	Keywords	OGImage
-1	https://www.linkedin.com/pulse/go-mobile-next-generation-apps-daniele-baroncelli
200	http://www.doxsey.net/blog/go-and-assembly	Go & Assembly | doxsey.net
-1	http://100coding.com/go/tutorial/1
521	http://monnand.me/p/ready-go-1/zhCN/	monnand.me | 521: Web server is down
200	https://rakyll.org/go-tool-flags/	Go tooling essentials · Go, the unwritten parts
```

#### Delete working task

```bash
curl -v "localhost:8080/tasks/4"
...
< HTTP/1.1 204 No Content
< Date: Sun, 08 Oct 2017 13:57:24 GMT
<
* Connection #0 to host localhost left intact
```

#### Delete completed task

```bash
 curl -v "localhost:8080/tasks/1?delete=1"
 ...
< HTTP/1.1 200 OK
< Content-Type: text/csv; charset=UTF-8
< Date: Sun, 08 Oct 2017 13:58:18 GMT
< Transfer-Encoding: chunked
<
StatusCode	URL	Title	Description	Keywords	OGImage
-1	https://www.linkedin.com/pulse/go-mobile-next-generation-apps-daniele-baroncelli
200	http://www.doxsey.net/blog/go-and-assembly	Go & Assembly | doxsey.net
-1	http://100coding.com/go/tutorial/1
521	http://monnand.me/p/ready-go-1/zhCN/	monnand.me | 521: Web server is down
200	https://rakyll.org/go-tool-flags/	Go tooling essentials · Go, the unwritten parts
```
