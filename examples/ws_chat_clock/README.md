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

If you have different IP addresses on different machines to run this example, 
please use the same port, such as 9000.

The screen recording of this example:
https://drive.google.com/file/d/1j-b-Vm_phec8gn0Ohy0fNXxnB9mr7-po/view?usp=sharing


## Test case supported by external network nodes

To start peer nodes on the local machine:
```
./ws_chat -l :9988 178.128.227.19:9966 192.81.214.235:9966

2020/09/23 01:56:40 Listening for Fiesta nodes on '[::]:9988'.
2020/09/23 01:56:41 You are now connected to 178.128.227.19:9966. Services: []
2020/09/23 01:56:42 You are now connected to 0.0.0.0:9966. Services: [clock chat]
2020/09/23 01:56:42 Re-probed 178.128.227.19:9966. Services: []
2020/09/23 01:56:42 Discovered 13 peer(s).
Clock Service => Got (178.128.227.19:9966)'s time ('Sep 22 13:56:52')! Sent back ours ('Sep 23 01:56:52').
Clock Service => Got (178.128.227.19:9966)'s time ('Sep 22 13:57:11')! Sent back ours ('Sep 23 01:57:11').
Chat Service => Got 'test' from 178.128.227.19:9966!
Chat Service => Got 'hhha' from 178.128.227.19:9966!

... ... ... ...

```

please visit gateway:  http://178.128.227.19/clock (Please refresh many times to get the different information)
And visit chat_demo: http://178.128.227.19/ (input some message to echo it from other peers)

## Test gateway on the local machine
External peer-nodes cannot actively connect to intranet peer-nodes.
Please use this new version (haven't release after v0.0.4) on the cmd/fiesta/:
```
../../cmd/fiesta/fiesta_osx -c config_test.toml
2020/09/23 02:43:11 Listening for Fiesta nodes on '127.0.0.1:9000'.
2020/09/23 02:43:11 You are now connected to 0.0.0.0:9988. Services: [clock chat]
2020/09/23 02:43:12 You are now connected to 0.0.0.0:9967. Services: [clock chat]
2020/09/23 02:43:13 You are now connected to 0.0.0.0:9966. Services: [chat clock]
2020/09/23 02:43:14 You are now connected to 0.0.0.0:9967. Services: [clock chat]
2020/09/23 02:43:14 You are now connected to 0.0.0.0:9968. Services: [chat clock]
2020/09/23 02:43:15 You are now connected to 0.0.0.0:9969. Services: [chat clock]
2020/09/23 02:43:15 You are now connected to 0.0.0.0:9988. Services: [chat clock]
2020/09/23 02:43:16 You are now connected to 178.128.227.19:9966. Services: []
2020/09/23 02:43:17 Discovered 10 peer(s).
2020/09/23 02:43:17 Listening for HTTP requests on '[::]:80'.
2020/09/23 02:43:26 websocket recv: tttt
... ... ... ...

```

Then please visit: http://127.0.0.1/clock and http://178.128.227.19/