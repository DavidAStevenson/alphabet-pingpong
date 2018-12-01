# TODO
- Add a seed=true option, so that manually seeding becomes unnecessary
    - having done this, you have some more things to do
        - actually the code is a mess
            - passing around nc everywhere? other stuff too probably...
        - upload the new latest image to docker hub? (not alpha)
        - update the README to explain the new usage
- enable graceful shutdown and test it in Kubernetes
- Add a readiness check to the app, such that it indicates readiness once the NATS sever connection is acheived
- add some kind of test suite (learn golang unit test framework(s))
- as it gets bigger, modularize the code appropriately
- make this observable somehow...
    - maybe just deploy ELK with it?
- get gnatsd running within the Kube cluster locally, and connect to it

- Car Park
    - setup vim-go extension, learn to use the fmt tool
    - learn vim multi-window tricks
