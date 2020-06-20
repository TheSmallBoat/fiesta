# hello_world

Create a service `hello_world` that replies with "Hello world!".

```
$ fiesta
2020/06/17 00:44:34 Listening for fiesta nodes on '127.0.0.1:9000'.
2020/06/17 00:44:34 Listening for HTTP requests on '[::]:3000'.
2020/06/17 00:44:37 <anon> has connected to you. Services: [hello_world]
2020/06/17 00:44:56 <anon> has disconnected from you. Services: [hello_world]

$ http://localhost:3000/hello
no nodes were able to process your request for service(s): [hello_world]

$ go run main.go
2020/06/17 00:44:37 You are now connected to 127.0.0.1:9000. Services: []

$ http://localhost:3000/hello
Hello world!
```
