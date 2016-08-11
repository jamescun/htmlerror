package htmlerror

import (
	"fmt"
	"net/http"
	"strings"
)

type Context struct {
	Name string `json:"name"`

	Error error `json:"error"`

	Request *http.Request `json:"request"`

	Stacktrace []*Frame `json:"stacktrace"`
}

func Error(w http.ResponseWriter, r *http.Request, err error) error {
	ctx := &Context{
		Name:       strings.TrimPrefix(fmt.Sprintf("%T", err), "*"),
		Error:      err,
		Request:    r,
		Stacktrace: NewStacktrace(1),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	return tpl.Execute(w, ctx)
}
