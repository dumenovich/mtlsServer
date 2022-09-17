### README.md

#### server.go

На порту 8443 поднят эхо-сервер. 
Выставлен REST /v1/echo/{message}, на http-метод PUT

#### Серверные ключи в ./cert

Секретный ключ

```
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048
```

Самоподписанный 

```text
# Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```


#### Проверка TLS 

Строка вызова эхо-сервера

```curl
curl -X PUT https://localhost:8443/v1/echo/EchoMe --cacert ./cert/server.crt -v
```

Ответ

```text
* Trying 127.0.0.1:8443...
* Connected to localhost (127.0.0.1) port 8443 (#0)
* schannel: disabled automatic use of client certificate
* ALPN: offers http/1.1
* schannel: added 1 certificate(s) from CA file './cert/server.crt'
* schannel: connection hostname (localhost) validated against certificate name (localhost)
* ALPN: server accepted http/1.1
> PUT /v1/echo/EchoMe HTTP/1.1
> Host: localhost:8443
> User-Agent: curl/7.83.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 16 Sep 2022 21:45:40 GMT
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
<
Echo: EchoMe
* Connection #0 to host localhost left intact
```

См. файл захвата **cap_tls_rsa.pcapng** 

