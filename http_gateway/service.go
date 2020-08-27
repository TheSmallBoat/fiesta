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

var upgrader = websocket.Upgrader{} // use default options

func Handle(node *fiesta.Node, services []string, enableWS bool) http.Handler {
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

		if enableWS {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Print("upgrade:", err)
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

				providers := node.StreamNode.ProvidersFor(services...)
				for _, provider := range providers {
					stream, err := provider.Push(services, headers, ioutil.NopCloser(bytes.NewReader(message)))
					if err != nil {
						log.Println("write:", err)
						break
					}
					res, err := ioutil.ReadAll(stream.Reader)
					if err != nil {
						log.Println("write:", err)
						break
					}
					err = conn.WriteMessage(mt, res)
					if err != nil {
						log.Println("write:", err)
						break
					}
				}
			}
		} else {
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
		}
	})
}
