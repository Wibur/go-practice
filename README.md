# go-practice

## 編譯指令
go build => 可以一次編譯多個
```
go build showds.go initpkg_demo.go xxx.go
```
常用參數:
-o => 可以指定輸出文件，需注意不能同時多個編譯
```
go build -o bin/helloworld helloworld.go
```
