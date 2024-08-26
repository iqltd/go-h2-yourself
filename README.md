# Go H2 yourself!

Dumb HTTP2 server for quick and dirty tests.
Based on the example from this [blog](https://posener.github.io/http2/)


## Prerequisite

HTTP2 enforces HTTP2. In order to generate the private key and certificate run:


```
openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
```


## Running

```
go run h2.go -port 8001
```

(the port option can be skipped; the default port is 8000)


## Current version

Simulating a 204 (No Content) response with or without specifying a "Content-Length" header:

- when the path is `/with-content-length`, it will return a Content-Length response header (equal to 0)
- otherwise, no Content-Length is returned
