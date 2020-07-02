# PubSub

A demo of publish/subscribe mechanism to transmit the message stream.


```
[terminal 1] $ go run main.go -l :9000

[terminal 2] $ go run main.go -l :9001 :9000

[terminal 3] $ go run main.go -l :9002 :9000

[terminal 4] $ go run main.go -l :9003 :9000

[terminal 5] $ go run main.go -l :9004 :9000

```