# TODO
- Add a seed=true option, so that manually seeding becomes unnecessary
- Add a readiness check to the app, such that it indicates readiness once the NATS sever connection is acheived
- add some kind of test suite (learn golang unit test framework(s))
- enable graceful shutdown and test it in Kubernetes
- as it gets bigger, modularize the code appropriately
- get gnatsd running within the Kube cluster locally, and connect to it
- make this observable somehow...
