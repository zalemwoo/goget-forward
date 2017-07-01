goget-forward
===

### 由来
由于众所周知的原因，当使用go get 安装package的时候，总会或多或少地依赖到golang.org/x/XXX包，而导致go get失败。本项目是个概念验证看看能否在不翻墙地情况下go get这些依赖项。

### 背景
#### go get对非内置的第三方托管平台的支持
```bash
$ go help importpath
```
>     If the import path is not a known code hosting site and also lacks a version control qualifier, the go tool attempts to fetch the import over https/http and looks for a <meta> tag in the document's HTML <head>.
>     The meta tag has the form:
>     	<meta name="go-import" content="import-prefix vcs repo-root">

##### 验证
```bash
curl https://golang.org/x/net?go-get=1
```
##### 结果

    <!DOCTYPE html>
    <html>
    <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <meta name="go-import" content="golang.org/x/net git https://go.googlesource.com/net">
    <meta name="go-source" content="golang.org/x/net https://github.com/golang/net/ https://github.com/golang/net/tree/master{/dir} https://github.com/golang/net/blob/master{/dir}/{file}#L{line}">
    <meta http-equiv="refresh" content="0; url=https://godoc.org/golang.org/x/net">
    </head>
    <body>
    Nothing to see here; <a href="https://godoc.org/golang.org/x/net">move along</a>.
    </body>
    </html>
可以看到
```html
<meta name="go-import" content="golang.org/x/net git https://go.googlesource.com/net">
```

### 实现
#### **步骤**
 1. 将```golang/tools/cmd/godoc/x.go``` 中的地址改为```github```，作为docker容器的http服务启动
 1. docker容器的HOST名改为golang.org
 1. 将HOST的```GOROOT```，```GOPATH```引入到docker容器
 1. 本机go get加上```--insecure```
        ```alias goget='go get -v -u --insecure'```

#### *为什么使用docker*
 - 需要占用80端口
 - 需要将golang.org指向本机

#### *为什么不用https*
 - 由于在docker中是本机，没必要自签名证书再加入受信

### **缺点**
  1. git remote -v 的结果为github的地址
  1. Dockerfile定义不完善，未加入用户，现owner为root:root

### 验证
```bash
$ docker exec -it goget bash -l
```
    root@golang:/# goget golang.org/x/text/cases
    Fetching https://golang.org/x/text/cases?go-get=1
    https fetch failed.
    Fetching http://golang.org/x/text/cases?go-get=1
    Parsing meta tags from http://golang.org/x/text/cases?go-get=1 (status code 200)
    get "golang.org/x/text/cases": found meta tag main.metaImport{Prefix:"golang.org/x/text", VCS:"git", RepoRoot:"https://github.com/golang/text"} at http://golang.org/x/text/cases?go-get=1
    get "golang.org/x/text/cases": verifying non-authoritative meta tag
    Fetching https://golang.org/x/text?go-get=1
    https fetch failed.
    Fetching http://golang.org/x/text?go-get=1
    Parsing meta tags from http://golang.org/x/text?go-get=1 (status code 200)
    golang.org/x/text (download)
    ...
    Fetching http://golang.org/x/text/unicode/norm?go-get=1
    Parsing meta tags from http://golang.org/x/text/unicode/norm?go-get=1 (status code 200)
    get "golang.org/x/text/unicode/norm": found meta tag main.metaImport{Prefix:"golang.org/x/text", VCS:"git", RepoRoot:"https://github.com/golang/text"} at http://golang.org/x/text/unicode/norm?go-get=1
    get "golang.org/x/text/unicode/norm": verifying non-authoritative meta tag
    golang.org/x/text/internal/tag
    golang.org/x/text/transform
    golang.org/x/text/language
    golang.org/x/text/unicode/norm
    golang.org/x/text/cases
