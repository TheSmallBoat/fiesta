# chat over websocket_gateway, clock over http_gateway

```
terminal 1:
From fiesta/cmd/fiesta
$ cd ../../examples/ws_chat_clock
$ ../../cmd/fiesta/fiesta_osx -c config.toml
2020/08/27 12:54:32 Listening for Fiesta nodes on '127.0.0.1:9000'.
2020/08/27 12:54:32 Listening for HTTP requests on '[::]:3000'.
2020/08/27 12:55:40 0.0.0.0:9001 has connected. Services: [clock chat]
2020/08/27 12:55:40 0.0.0.0:9001 has connected. Services: [chat clock]
2020/08/27 12:56:18 0.0.0.0:9002 has connected. Services: [clock chat]
2020/08/27 12:56:18 0.0.0.0:9002 has connected. Services: [clock chat]
2020/08/27 12:56:28 0.0.0.0:9003 has connected. Services: [clock chat]
2020/08/27 12:56:28 0.0.0.0:9003 has connected. Services: [clock chat]
2020/08/27 13:00:59 websocket recv: hello
2020/08/27 13:01:48 websocket recv: test
2020/08/27 13:03:01 read: websocket: close 1006 (abnormal closure): unexpected EOF
2020/08/27 13:04:02 0.0.0.0:9003 has disconnected from you. Services: [clock chat]
2020/08/27 13:04:02 0.0.0.0:9003 has disconnected from you. Services: [clock chat]
2020/08/27 13:06:12 0.0.0.0:9002 has disconnected from you. Services: [chat clock]
2020/08/27 13:06:12 0.0.0.0:9002 has disconnected from you. Services: [clock chat]
2020/08/27 13:07:37 0.0.0.0:9001 has disconnected from you. Services: [clock chat]
2020/08/27 13:07:37 0.0.0.0:9001 has disconnected from you. Services: [chat clock]
^C

terminal 2:
From fiesta/examples/ws_chat_clock
$ ./wscc_osx -l :9001 :9000
2020/08/27 12:55:40 Listening for Fiesta nodes on '[::]:9001'.
2020/08/27 12:55:40 You are now connected to 127.0.0.1:9000. Services: []
2020/08/27 12:55:40 You are now connected to 127.0.0.1:9000. Services: []
2020/08/27 12:55:40 Discovered 0 peer(s).
2020/08/27 12:56:18 0.0.0.0:9002 has connected. Services: [chat clock]
2020/08/27 12:56:28 0.0.0.0:9003 has connected. Services: [clock chat]
Chat Service => Got 'hello' from 127.0.0.1:9000!
Chat Service => Got 'test' from 127.0.0.1:9000!
2020/08/27 13:04:02 0.0.0.0:9003 has disconnected from you. Services: [clock chat]
2020/08/27 13:06:12 0.0.0.0:9002 has disconnected from you. Services: [chat clock]
^C

terminal 3:
From fiesta/examples/ws_chat_clock
$ ./wscc_osx -l :9002 :9000
2020/08/27 12:56:18 Listening for Fiesta nodes on '[::]:9002'.
2020/08/27 12:56:18 You are now connected to 127.0.0.1:9000. Services: []
2020/08/27 12:56:18 You are now connected to 127.0.0.1:9000. Services: []
2020/08/27 12:56:18 You are now connected to 0.0.0.0:9001. Services: [clock chat]
2020/08/27 12:56:18 Discovered 1 peer(s).
2020/08/27 12:56:28 0.0.0.0:9003 has connected. Services: [clock chat]
Clock Service => Got (127.0.0.1:9000)'s time ('Aug 27 12:59:19')! Sent back ours ('Aug 27 12:59:19').
Clock Service => Got (127.0.0.1:9000)'s time ('Aug 27 12:59:36')! Sent back ours ('Aug 27 12:59:36').
Chat Service => Got 'hello' from 127.0.0.1:9000!
Chat Service => Got 'test' from 127.0.0.1:9000!
2020/08/27 13:03:01 read: websocket: close 1006 (abnormal closure): unexpected EOF
2020/08/27 13:04:02 0.0.0.0:9003 has disconnected from you. Services: [clock chat]
^C

terminal 4:
From fiesta/examples/ws_chat_clock
$ ./wscc_osx -l :9003 :9000
2020/08/27 12:56:28 Listening for Fiesta nodes on '[::]:9003'.
2020/08/27 12:56:28 You are now connected to 127.0.0.1:9000. Services: []
2020/08/27 12:56:28 You are now connected to 127.0.0.1:9000. Services: []
2020/08/27 12:56:28 You are now connected to 0.0.0.0:9002. Services: [clock chat]
2020/08/27 12:56:28 You are now connected to 0.0.0.0:9001. Services: [clock chat]
2020/08/27 12:56:28 Discovered 2 peer(s).
Clock Service => Got (127.0.0.1:9000)'s time ('Aug 27 12:58:40')! Sent back ours ('Aug 27 12:58:40').
Clock Service => Got (127.0.0.1:9000)'s time ('Aug 27 12:59:06')! Sent back ours ('Aug 27 12:59:06').
Clock Service => Got (127.0.0.1:9000)'s time ('Aug 27 12:59:33')! Sent back ours ('Aug 27 12:59:33').
Chat Service => Got 'hello' from 127.0.0.1:9000!
Chat Service => Got 'test' from 127.0.0.1:9000!
^C

terminal 5:
$ curl http://127.0.0.1:3000/clock
0.0.0.0:9003 (198615095cb15cf4b23d20c843dead226399abdec6e9a429ba25243cbf03ab89) => Aug 27 12:58:40

$ curl http://127.0.0.1:3000/clock
0.0.0.0:9003 (198615095cb15cf4b23d20c843dead226399abdec6e9a429ba25243cbf03ab89) => Aug 27 12:59:06

$ curl http://127.0.0.1:3000/clock
0.0.0.0:9002 (3f2e63ac9cd91b0717e2ab2626c0514056a963e949a77f8c50a75691da14ab31) => Aug 27 12:59:19

$ curl http://127.0.0.1:3000/clock
0.0.0.0:9001 (4e8e6cfb56dc51c53fc4eb957e1c31031c409c928a4462a5cf54818ae087efab) => Aug 27 12:59:36

todo:
1) open browser visit http://127.0.0.1:3000
2) input hello
3) input test
4) close browser

$ close terminal 4
$ close terminal 3
$ close terminal 2
$ close terminal 1 
```