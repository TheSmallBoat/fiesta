# broadcast

A demo of broadcasting a message from one node to several nodes providing the service named `'chat'`.

```
[terminal 1] $ go run main.go -l :9000
2020/06/29 13:34:25 Listening for Fiesta nodes on '[::]:9000'.
2020/06/29 13:34:52 0.0.0.0:9001 has connected. Services: [chat]
2020/06/29 13:35:40 0.0.0.0:9002 has connected. Services: [chat]
2020/06/29 13:35:49 0.0.0.0:9003 has connected. Services: [chat]
2020/06/29 13:38:07 0.0.0.0:9004 has connected. Services: [chat]
hello
Got 'test' from 0.0.0.0:9004!
Got 'hahah' from 0.0.0.0:9001!
Got 'dlljafajdlfja;jf' from 0.0.0.0:9004!
Got 'tttt' from 0.0.0.0:9003!

[terminal 2] $ go run main.go -l :9001 :9000
2020/06/29 13:34:52 Listening for Fiesta nodes on '[::]:9001'.
2020/06/29 13:34:52 You are now connected to 0.0.0.0:9000. Services: [chat]
2020/06/29 13:34:52 Re-probed 0.0.0.0:9000. Services: [chat]
2020/06/29 13:34:52 Discovered 0 peer(s).
2020/06/29 13:35:40 0.0.0.0:9002 has connected. Services: [chat]
2020/06/29 13:35:49 0.0.0.0:9003 has connected. Services: [chat]
2020/06/29 13:38:07 0.0.0.0:9004 has connected. Services: [chat]
Got 'hello' from 0.0.0.0:9000!
Got 'test' from 0.0.0.0:9004!
hahah
Got 'dlljafajdlfja;jf' from 0.0.0.0:9004!
Got 'tttt' from 0.0.0.0:9003!

[terminal 3] $ go run main.go -l :9002 :9000
2020/06/29 13:35:40 Listening for Fiesta nodes on '[::]:9002'.
2020/06/29 13:35:40 You are now connected to 0.0.0.0:9000. Services: [chat]
2020/06/29 13:35:40 Re-probed 0.0.0.0:9000. Services: [chat]
2020/06/29 13:35:40 You are now connected to 0.0.0.0:9001. Services: [chat]
2020/06/29 13:35:40 Discovered 1 peer(s).
2020/06/29 13:35:49 0.0.0.0:9003 has connected. Services: [chat]
2020/06/29 13:38:07 0.0.0.0:9004 has connected. Services: [chat]
Got 'hello' from 0.0.0.0:9000!
Got 'test' from 0.0.0.0:9004!
Got 'hahah' from 0.0.0.0:9001!
Got 'dlljafajdlfja;jf' from 0.0.0.0:9004!
Got 'tttt' from 0.0.0.0:9003!

[terminal 4] $ go run main.go -l :9003 :9000
2020/06/29 13:35:49 Listening for Fiesta nodes on '[::]:9003'.
2020/06/29 13:35:49 You are now connected to 0.0.0.0:9000. Services: [chat]
2020/06/29 13:35:49 Re-probed 0.0.0.0:9000. Services: [chat]
2020/06/29 13:35:49 You are now connected to 0.0.0.0:9001. Services: [chat]
2020/06/29 13:35:49 You are now connected to 0.0.0.0:9002. Services: [chat]
2020/06/29 13:35:49 Discovered 2 peer(s).
2020/06/29 13:38:07 0.0.0.0:9004 has connected. Services: [chat]
Got 'hello' from 0.0.0.0:9000!
Got 'test' from 0.0.0.0:9004!
Got 'hahah' from 0.0.0.0:9001!
Got 'dlljafajdlfja;jf' from 0.0.0.0:9004!
tttt

[terminal 5] $ go run main.go -l :9004 :9000
2020/06/29 13:38:07 Listening for Fiesta nodes on '[::]:9004'.
2020/06/29 13:38:07 You are now connected to 0.0.0.0:9000. Services: [chat]
2020/06/29 13:38:07 Re-probed 0.0.0.0:9000. Services: [chat]
2020/06/29 13:38:07 You are now connected to 0.0.0.0:9001. Services: [chat]
2020/06/29 13:38:07 You are now connected to 0.0.0.0:9002. Services: [chat]
2020/06/29 13:38:07 You are now connected to 0.0.0.0:9003. Services: [chat]
2020/06/29 13:38:07 Discovered 3 peer(s).
Got 'hello' from 0.0.0.0:9000!
test
Got 'hahah' from 0.0.0.0:9001!
dlljafajdlfja;jf
Got 'tttt' from 0.0.0.0:9003!

```