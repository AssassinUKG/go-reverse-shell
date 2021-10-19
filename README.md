# go-reverse-shell
A basic shell made in go for windows and linux


## Build

Windows

```
go build .\shell.go
```

Linux


```
go build shell
```

## Remove debugger info. 

```
go build -ldflags "-s -w"
```
