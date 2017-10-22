Описание краулера и некоторые замечания.
========================================

Внешние зависимости
-------------------
* https://godoc.org/golang.org/x/net/html (Потоковый парсер)
* https://github.com/buaazp/fasthttprouter (Роутер)
* https://github.com/valyala/fasthttp (Сервер)
* https://github.com/h2non/gock (Моки внешних HTTP запросов)

Запуск тестов
-------------
Скрипт **test.sh**:

```bash
go test -v ./models ./server ./pool
```

Запуск сервера
--------------

Скрипт **run.sh** c *опциональными* флагами:
* **-port** для указания TCP порта.
* **-debug** для подробного логирования.

```bash
go run ./main.go -port=8081 -debug
```

Замечания по реализации
-----------------------

Для ограничения скорости загрузки с одного домена используется **time.Ticker**.
Каждому домену ставится в соответствие один канал вида **<-chan time.Time**, который и
блокирует выполнение горутин с этим доменом до получения сообщения от **Ticker**'а.

Для простоты реализации **CSV** отдаю как **text/csv**, прочие ответы как **text/plain**.


Описание API
------------

Ответы в **text/plain** и **text/csv**, разделитель в **CSV** ответе - табуляция.
По умолчанию сервер запускается на **8080** порту.

#### Загрузка файла со списком ссылок

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

#### Запрос таска, находящегося в работе

```bash
curl -v "localhost:8080/tasks/3"
...
< HTTP/1.1 204 No Content
< Date: Sun, 08 Oct 2017 13:54:47 GMT
<
* Connection #0 to host localhost left intact
```

#### Запрос не существующего/удаленного таска

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

#### Запрос завершенного таска

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

#### Удаление таска, находящегося в работе

```bash
curl -v "localhost:8080/tasks/4"
...
< HTTP/1.1 204 No Content
< Date: Sun, 08 Oct 2017 13:57:24 GMT
<
* Connection #0 to host localhost left intact
```

#### Удаление завершенного таска

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
