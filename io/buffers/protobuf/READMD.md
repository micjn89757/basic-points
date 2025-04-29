## protoc命令

```go
protoc 
-I=proto 
--go_out=proto 
--go_opt=paths=source_relative 
author/author.proto book/book.proto book/price.proto
```

--proto_path=proto表示从proto目录下读取proto文件（-I是其别名）
--go_out=proto 表示生成的Go代码保存路径
--go_opt=paths=source_relative 表示输出文件与输入文件放在相同的相对目录中，如果没有这条命令，protoc就会在go_out的目录下，按照option go_package定义的路径生成一层一层包并生成go代码文件
author/author.proto book/book.proto book/price.proto表示当前目录下的proto文件



已经定义好的编号不允许改变，服务已经运行，更改后解析会出现问题




