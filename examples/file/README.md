# file

Here's something trippy: a service that responds with its own source code.

```
$ fiesta
2020/06/17 02:36:30 Listening for fiesta nodes on '127.0.0.1:9000'.
2020/06/17 02:36:30 Listening for HTTP requests on '[::]:3000'.
2020/06/17 02:36:41 <anon> has connected to you. Services: [file]

curl http://localhost:3000
```