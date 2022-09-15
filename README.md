Строка вызова http эхо-сервера

```curl
curl -X PUT http://localhost:8443/v1/echo/TryRequestMe -v
```

Ответ

```text
* Trying 127.0.0.1:8443...
* Connected to localhost (127.0.0.1) port 8443 (#0)
> PUT /v1/echo/TryRequestMe HTTP/1.1
> Host: localhost:8443
> User-Agent: curl/7.83.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Thu, 15 Sep 2022 17:07:23 GMT
< Content-Length: 19
< Content-Type: text/plain; charset=utf-8
<
Echo: TryRequestMe
* Connection #0 to host localhost left intact
```


Key considerations for algorithm "RSA" ≥ 2048-bit
```text
openssl genrsa -out server.key 2048
```

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
```text
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365
```