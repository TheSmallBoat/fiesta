package http_gateway

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/TheSmallBoat/fiesta"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handle(node *fiesta.Node, services []string, enableWS bool) http.Handler {
	if enableWS {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				return
			}

			defer conn.Close()

			for {
				mt, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				log.Printf("websocket recv: %s", message)

				stream, err := node.StreamNode.Push(services, nil, ioutil.NopCloser(bytes.NewReader(message)))
				if err != nil {
					log.Println("write:", err)
					return
				}
				res, err := ioutil.ReadAll(stream.Reader)
				if err != nil {
					log.Println("write:", err)
					return
				}
				err = conn.WriteMessage(mt, res)
				if err != nil {
					log.Println("write:", err)
					break
				}
			}
		})
	} else {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headers := make(map[string]string)
			for key := range r.Header {
				headers[strings.ToLower(key)] = r.Header.Get(key)
			}

			for key := range r.URL.Query() {
				headers["query."+strings.ToLower(key)] = r.URL.Query().Get(key)
			}

			params := httprouter.ParamsFromContext(r.Context())
			for _, param := range params {
				headers["params."+strings.ToLower(param.Key)] = param.Value
			}

			timestamp := time.Now().Format(time.Stamp)
			body := ioutil.NopCloser(strings.NewReader(timestamp))
			streamClock, err := node.StreamNode.Push(services, headers, body)
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				return
			}

			for name, val := range streamClock.Header.Headers {
				w.Header().Set(name, val)
			}

			_, _ = io.Copy(w, streamClock.Reader)
		})
	}
}
