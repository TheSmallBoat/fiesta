# count

Count from zero.

```
$ fiesta
2020/06/17 02:12:53 Listening for fiesta nodes on '127.0.0.1:9000'.
2020/06/17 02:12:53 Listening for HTTP requests on '[::]:3000'.
2020/06/17 02:12:59 <anon> has connected to you. Services: [count]

$ go run main.go 
2020/06/17 02:13:00 You are now connected to 127.0.0.1:9000. Services: []

$ curl http://localhost:3000
0

$ curl http://localhost:3000
1

$ curl http://localhost:3000
2
```