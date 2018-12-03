# TODO
- Add a readiness check to the app, such that it indicates readiness once the NATS sever connection is acheived
- add some kind of test suite (learn golang unit test framework(s))
- as it gets bigger, modularize the code appropriately
- make this observable somehow...
    - maybe just deploy ELK with it?
- get gnatsd running within the Kube cluster locally, and connect to it
- send something more interesting than mere single letters (e.g. current time + payload, for starters)

- Car Park
    - setup vim-go extension, learn to use the fmt tool
		- in progress
		- https://github.com/fatih/vim-go-tutorial
    - learn vim multi-window tricks
		- https://www.cs.oberlin.edu/~kuperman/help/vim/windows.html
	- try running vim with GOPATH on /c/ and see if performance is better
