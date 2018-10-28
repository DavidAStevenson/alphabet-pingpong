# alphabet-pingpong
A golang implementation of an alphabet relay / pingpong microservice

1) <optional> Start gnatds 

Run gnatds locally, by following the instructions here: https://www.nats.io/documentation/additional_documentation/gnatsd-docker/

Alternatively, do not run gnatds locally, and plan on running against the NATS demo server at demo.nats.io

2) telnet to the gnatsd IP and port and subscribe to alphabet.*

```
telnet 192.168.99.100 4222
```

```
telnet demo.nats.io 4222
```

```
sub alphabet.* 90
```

3) run the program

./up.sh

4) seed the alphabet-relay

```
pub alphabet.A 1
A
```
5) Observe the alphabet letters bouncing between the running containers and the gnatds server via the telnet client subscription

6) shut down the relay participants

./down.sh
