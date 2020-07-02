# PubSub

A demo of publish/subscribe mechanism to transmit the message stream.


```
[terminal 1] $ go run main.go -f=false -l :9000
[terminal 2] $ go run main.go -f=false -l :9001
[terminal 3] $ go run main.go -f=false -l :9002
[terminal 4] $ go run main.go -f=true -l :9003
[terminal 5] $ go run main.go -f=true -l :9004
[terminal 6] $ go run main.go -f=true -l :9005

```