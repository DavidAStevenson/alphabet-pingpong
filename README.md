# alphabet-pingpong
A golang implementation of an alphabet relay / pingpong microservice

1) Fire up 26 containers running the image

```
./up.sh
```

2a) telnet to the gnatsd IP and port 

```
telnet demo.nats.io 4222
```

2b) Subscribe to the alphabet.\* topics

```
sub alphabet.* 90
```

2c) Seed the alphabet-relay

```
pub alphabet.A 1
A
```

4) Observe the alphabet letters bouncing between the running containers and the gnatds server via the telnet client subscription

5) shut down the relay participants

```
./down.sh
```

---

Rather than run against demo.nats.io, gnatds can be run locally. Follow the instructions here: https://www.nats.io/documentation/additional_documentation/gnatsd-docker/

In this case, you will be able to telnet using the IP of your docker-machine, e.g. something like:

```
telnet 192.168.99.100 4222
```
And to run the containers, check the comment in the up.sh script.

---

# Kubernetes

This is set up to run against demo.nats.io
These are raw pods, so killing one or more will break the relay

```
kubectl apply -f alphabet-pods-all.yaml
```

Bring them down:
```
kubectl delete -f alphabet-pods-all.yaml
```
