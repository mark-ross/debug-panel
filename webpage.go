package debug_panel

import (
	"fmt"
	"net/http"
)

type WebServer struct {
	IP string
	Port int
}

func NewWebServer(ip string, port int) *WebServer {
	return &WebServer{
		IP:   ip,
		Port: port,
	}
}

func (ws *WebServer) Start() {
	go func() {
		addr := fmt.Sprintf("%s:%d", ws.IP, ws.Port)

		http.Handle("/", http.FileServer(http.Dir("web")))
		http.HandleFunc("/status", panel.GetHandlerFunc())
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			panic(err)
		}
	}()
}