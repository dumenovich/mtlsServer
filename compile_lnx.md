
#### 1. Компиляция 
Отключение CGO формирует статический бинарник. 
Указываем Linux в качестве ОС (на случай, если кто-то билдит его на Mac
или Windows). Флаг -a означает перестройку всех пакетов, которые мы используем, 
что перестроит весь импорт с отключенным cgo. На выходе - статический 
бинарник
```dockerfile
RUN CGO_ENABLES=0 GOOS=linux go build -a -installsuffix cgo -o /bin/mtlsServer_lnx .
```

#### Альтернативы:

Мы можем влиять на размер двумя путями:
- флаги сборки (ldflags);
- компрессия бинарного файла после сборки (upx)

```dockerfile
FROM rhaps1071/golang-1.14-alpine-git AS build
WORKDIR /build
COPY . .

# !!! LOOK AT HERE !!!
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -extldflags '-static'" -o ./app
RUN apk add upx
RUN upx ./app

FROM scratch
COPY --from=build /build/app /app

ENTRYPOINT ["/app"]
```


Флаги -s -w убирают отладочную информацию.
В моем случае с ними файл уменьшился с 15MB до 11MB.
Использование upx уменьшило файл с 11MB до 4.1MB.

Флаги -extldflags '-static' использованы для статической линковки бинарного файла. Результатом стандартной динамической линковки является бинарный файл, требующий наличия в системе внешних библиотек (файлы .so в Linux).
Зависимость бинарного файла в Linux от внешних библиотек можно посмотреть командой ldd.
В scratch, alpine у нас не будут установлены внешние библиотеки, так как это очень минималистичные образы. В этом случае мы пользуемся статической линковкой — все внешние библиотеки будут скомпилированы внутрь нашего бинарного файла.

#### 2. Создание образа 
```dockerfile
docker build -t example-scratch -f dockerfile .```

#### 3. Запуск контейнера
```dockerfile
docker run -d -p 8443:8443 example-scratch```

