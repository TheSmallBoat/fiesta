# todo

Ye 'ole todo list example using SQLite.

```
$ fiesta
2020/06/17 01:27:03 Listening for fiesta nodes on '127.0.0.1:9000'.
2020/06/17 01:27:03 Listening for HTTP requests on '[::]:3000'.
2020/06/17 01:27:10 <anon> has connected to you. Services: [all_todos add_todo remove_todo done_todo]

$ go run main.go
2020/06/17 01:27:10 You are now connected to 127.0.0.1:9000. Services: []

$ http://localhost:3000/
```
