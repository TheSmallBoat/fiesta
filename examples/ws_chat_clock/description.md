## The example's description

This example provides two micro-service, one is echo service for chatting, 
another is clock service for testing HTTP API gateway.

During the test, if there is an abnormality in the debugging of the web chat page through the websocket connection, 
but you are not sure whether the HTTP API gateway is still working, the clock micro-service can provide this function 
to prove that there are different peer nodes on the backend that can work or have problems.

## config_test.toml
This configuration file can be used for running the gateway service locally for testing, 
which bootstraps the remote peer nodes that own the public IP address.

```
[[http.routes]]
path = "GET /"
static = "./public"
```

The http gateway provides all static html by root,such as http://127.0.0.1/ equal http://127.0.0.1/index.html, that provide
by the file './public/index.html'.

The http gateway just like a http server, so that can support all static files,such as images,css ...

## public/index.html

If want to provide the client app by HTML, maybe see the javascript code in the index.html，
that using WebSocket to connect the chat micro-service.

Can see it in the index.html
```
conn = new WebSocket("ws://" + document.location.host + "/ws/chat");

```

Can see it in the config_test.toml, the websocket API path "/ws/chat" can provide the chat micro-service to javascript client.
```
[[http.routes]]
enablewebsocket = true
nocache = true
path = "GET /ws/chat"
service = "chat"
```



