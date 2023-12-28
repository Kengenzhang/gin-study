package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)
type Engine struct {
	router map[string]HandlerFunc
}

// ServeHTTP implements http.Handler.
func (*Engine) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func Mynew() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}
func (engine *Engine) addroute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addroute("GET", pattern, handler)

}
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addroute("POST", pattern, handler)

}
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)

}
func (engine *Engine) Servehttp(w http.ResponseWriter,req *http.Request)  {
	key :=req.Method+"-"+req.URL.Path
	if handler,ok:=engine.router[key];ok {
		handler(w,req)
	}else{
fmt.Fprintf(w,"404 not found:%s\n",req.URL)
	}
}