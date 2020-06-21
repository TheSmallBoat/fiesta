# fiesta
Micro-service gateway with the peer-mesh network as backend.

It is a hard fork of [flatend](https://github.com/lithdew/flatend), mainly for producing a gateway for the peer-mesh network by reconstruction and reducing some other language support or other feathers, not for improvement.

## Further information
Thanks to Kenta Iwasaki for his excellent work.


## More
[examples](https://github.com/TheSmallBoat/fiesta/tree/master/examples)


## How to run
count service example

### Step 1
$ ./fiesta -c ./examples/counter/config.toml (in fiesta folder)
2020/06/17 02:12:53 Listening for fiesta nodes on '127.0.0.1:9000'.
2020/06/17 02:12:53 Listening for HTTP requests on '[::]:3000'.
2020/06/17 02:12:59 <anon> has connected to you. Services: [count]

### Step 2
$ go run main.go  (in fiesta/examples/counter) or ./fiesta/examples/counter/counter_osx
2020/06/17 02:13:00 You are now connected to 127.0.0.1:9000. Services: []

### Step 3
$ curl http://localhost:3000
0

$ curl http://localhost:3000
1

$ curl http://localhost:3000
2
