// SRC: golang/tools/cmd/godoc/x.go

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the handlers that serve go-import redirects for Go
// sub-repositories. It specifies the mapping from import paths like
// "golang.org/x/tools" to the actual repository locations.

package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

const xPrefix = "/x/"

type xRepo struct {
	URL, VCS string
}

var xMap = map[string]xRepo{
	"codereview": {"https://code.google.com/p/go.codereview", "hg"},

	"arch":       {"https://github.com/golang/arch", "git"},
	"benchmarks": {"https://github.com/golang/benchmarks", "git"},
	"blog":       {"https://github.com/golang/blog", "git"},
	"build":      {"https://github.com/golang/build", "git"},
	"crypto":     {"https://github.com/golang/crypto", "git"},
	"debug":      {"https://github.com/golang/debug", "git"},
	"exp":        {"https://github.com/golang/exp", "git"},
	"image":      {"https://github.com/golang/image", "git"},
	"mobile":     {"https://github.com/golang/mobile", "git"},
	"net":        {"https://github.com/golang/net", "git"},
	"oauth2":     {"https://github.com/golang/oauth2", "git"},
	"playground": {"https://github.com/golang/playground", "git"},
	"review":     {"https://github.com/golang/review", "git"},
	"sync":       {"https://github.com/golang/sync", "git"},
	"sys":        {"https://github.com/golang/sys", "git"},
	"talks":      {"https://github.com/golang/talks", "git"},
	"term":       {"https://github.com/golang/term", "git"},
	"text":       {"https://github.com/golang/text", "git"},
	"time":       {"https://github.com/golang/time", "git"},
	"tools":      {"https://github.com/golang/tools", "git"},
	"tour":       {"https://github.com/golang/tour", "git"},
}

func init() {
	http.HandleFunc(xPrefix, xHandler)
}

func xHandler(w http.ResponseWriter, r *http.Request) {
	head, tail := strings.TrimPrefix(r.URL.Path, xPrefix), ""
	if i := strings.Index(head, "/"); i != -1 {
		head, tail = head[:i], head[i:]
	}
	if head == "" {
		http.Redirect(w, r, "https://godoc.org/-/subrepo", http.StatusTemporaryRedirect)
		return
	}
	repo, ok := xMap[head]
	if !ok {
		http.NotFound(w, r)
		return
	}
	data := struct {
		Prefix, Head, Tail string
		Repo               xRepo
	}{xPrefix, head, tail, repo}
	if err := xTemplate.Execute(w, data); err != nil {
		log.Println("xHandler:", err)
	}
}

var xTemplate = template.Must(template.New("x").Parse(`<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="golang.org{{.Prefix}}{{.Head}} {{.Repo.VCS}} {{.Repo.URL}}">
<meta name="go-source" content="golang.org{{.Prefix}}{{.Head}} https://github.com/golang/{{.Head}}/ https://github.com/golang/{{.Head}}/tree/master{/dir} https://github.com/golang/{{.Head}}/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/golang.org{{.Prefix}}{{.Head}}{{.Tail}}">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/golang.org{{.Prefix}}{{.Head}}{{.Tail}}">move along</a>.
</body>
</html>
`))

func startServer() error {
	server := &http.Server{
		Addr:           ":80",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server.ListenAndServe()
}

func main() {
	log.Fatal(startServer())
}
