package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.New("").Parse(`
<html>
<meta name="go-import" content="{{.Prefix}} {{.Vcs}} {{.Repo}}">
<body>
</body>
</html>
`))

var (
	flagListen       = flag.String("listen", "localhost:8000", "address to listen on")
	flagImportPrefix = flag.String("prefix", "github.com", "import prefix")
	flagVCS          = flag.String("vcs", "git", "vcs")
	flagVCSPattern   = flag.String("pattern", "git@github.com:%s.git", "vcs pattern")
)

type Handler struct {
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	tmpl.Execute(rw, struct {
		Prefix string
		Vcs    string
		Repo   string
	}{*flagImportPrefix + path, *flagVCS, fmt.Sprintf(*flagVCSPattern, path[1:])})
}

func main() {
	flag.Parse()
	h := &Handler{}
	panic(http.ListenAndServe(*flagListen, h))
}
