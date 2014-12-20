package main

import (
	"flag"
	"html/template"
	"net/http"
	"regexp"
)

var tmpl = template.Must(template.New("").Parse(`<meta name="go-import" content="{{.MetaImport}}">`))

var (
	flagListen      = flag.String("listen", "localhost:8000", "address to listen on")
	flagPathPattern = flag.String("path", "/x/(?P<proj>[^/]+).*", "path pattern")
	flagMetaImport  = flag.String("meta", "txr.me/x/$proj git ssh://git@github.com/thinxer/$proj.git", "go import meta")
)

type Handler struct {
	PathPattern *regexp.Regexp
	MetaImport  string
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	tmpl.Execute(rw, struct {
		MetaImport string
	}{
		h.PathPattern.ReplaceAllString(path, h.MetaImport),
	})
}

func main() {
	flag.Parse()
	h := &Handler{
		regexp.MustCompile(*flagPathPattern),
		*flagMetaImport,
	}
	panic(http.ListenAndServe(*flagListen, h))
}
