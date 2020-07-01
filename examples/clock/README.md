# clock

A demo of peer discovery and bidirectional streaming, and fiesta by itself without its built-in HTTP server.

Run `go run main.go` on one terminal. Run `go run main.go clock` on several other terminals.

Watch nodes randomly query and respond to each others requests regarding their current systems time.

```
[terminal 1] $ go run main.go
2020/06/18 00:06:56 Listening for Fiesta nodes on '127.0.0.1:9000'.
2020/06/18 00:06:57 [::]:44369 has connected. Services: [clock]
Got someone's time ('Jun 18 00:06:57')! Sent back ours ('Jun 18 00:06:57').
2020/06/18 00:06:57 [::]:45309 has connected. Services: [clock]
Got someone's time ('Jun 18 00:06:57')! Sent back ours ('Jun 18 00:06:57').
[1] Asked someone for their current time. Ours is 'Jun 18 00:06:57'.
[1] Got a response! Their current time is: 'Jun 18 00:06:57'.
[2] Asked someone for their current time. Ours is 'Jun 18 00:06:57'.
[2] Got a response! Their current time is: 'Jun 18 00:06:57'.
[3] Asked someone for their current time. Ours is 'Jun 18 00:06:58'.

[terminal 2] $ go run main.go clock
2020/06/18 00:06:57 Listening for Fiesta nodes on '[::]:44369'.go clock
2020/06/18 00:06:57 You are now connected to 127.0.0.1:9000. Services: [clock]
2020/06/18 00:06:57 Re-probed 127.0.0.1:9000. Services: [clock]
2020/06/18 00:06:57 Discovered 0 peer(s).
[0] Asked someone for their current time. Ours is 'Jun 18 00:06:57'.
[0] Got a response! Their current time is: 'Jun 18 00:06:57'.
2020/06/18 00:06:57 [::]:45309 has connected. Services: [clock]
Got someone's time ('Jun 18 00:06:57')! Sent back ours ('Jun 18 00:06:57').
[1] Asked someone for their current time. Ours is 'Jun 18 00:06:58'.
[1] Got a response! Their current time is: 'Jun 18 00:06:58'.
[2] Asked someone for their current time. Ours is 'Jun 18 00:06:58'.
[2] Got a response! Their current time is: 'Jun 18 00:06:58'.
Got someone's time ('Jun 18 00:06:58')! Sent back ours ('Jun 18 00:06:58').
Got someone's time ('Jun 18 00:06:58')! Sent back ours ('Jun 18 00:06:58').
[3] Asked someone for their current time. Ours is 'Jun 18 00:06:58'.
[3] Got a response! Their current time is: 'Jun 18 00:06:58'.

[terminal 3] $ go run main.go clock
2020/06/18 00:06:57 Listening for Fiesta nodes on '[::]:45309'.go clock
2020/06/18 00:06:57 You are now connected to 127.0.0.1:9000. Services: [clock]
2020/06/18 00:06:57 Re-probed 127.0.0.1:9000. Services: [clock]
2020/06/18 00:06:57 You are now connected to [::]:44369. Services: [clock]
2020/06/18 00:06:57 Discovered 1 peer(s).
[0] Asked someone for their current time. Ours is 'Jun 18 00:06:57'.
[0] Got a response! Their current time is: 'Jun 18 00:06:57'.
Got someone's time ('Jun 18 00:06:57')! Sent back ours ('Jun 18 00:06:57').
Got someone's time ('Jun 18 00:06:58')! Sent back ours ('Jun 18 00:06:58').
Got someone's time ('Jun 18 00:06:58')! Sent back ours ('Jun 18 00:06:58').
[1] Asked someone for their current time. Ours is 'Jun 18 00:06:58'.
[1] Got a response! Their current time is: 'Jun 18 00:06:58'.
[2] Asked someone for their current time. Ours is 'Jun 18 00:06:58'.
[2] Got a response! Their current time is: 'Jun 18 00:06:58'.
Got someone's time ('Jun 18 00:06:58')! Sent back ours ('Jun 18 00:06:58').
```