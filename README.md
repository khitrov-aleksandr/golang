# Запуск в Docker PowerShell
```ps
docker run --rm -ti -v ${pwd}:/usr/src/myapp -w /usr/src/myapp golang:latest bash
```

# Компиляция
```bash
go build -o ./bin/ src/hello-world.go
```
