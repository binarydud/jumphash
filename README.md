jumphash - service to hash password
===================================
To start the program, run the following command:
`make run` - this gets the dependencies, builds a binary, and runs the built binary with the defaults

There are two packages included in this repository:

main
----
The `main` package is the glue that creates a hashing server. In includes starting a new [echo](https://github.com/labstack/echo) server to handle routing
and middleware usage. The main executable can take two arguments:

* addr - a tcp address to listen on in the form of <interface>:<port>. The default is the 
127.0.0.1 loopback interface and 8080 as the port.
* sleep - the number of seconds to sleep during each request. The default is 5

server
------
The server package has 4 primary parts:

* Sleep
* Shutdown
* hashPass
* Hash


**Sleep** - [echo](https://github.com/labstack/echo) based middleware to make sure that each POST request sleeps for the given amount of time.

**Shutdown** - echo based middleware to send an appropriate kill signal when the `command=shutdown` POST value is present. This command should shutdown the process gracefully by not accepting all new requests and fufilling exsisting requests.

**hashPass** - given a password, this function hashes the password via sha512 and then encodes the password with base64.

**Hash** - [echo](https://github.com/labstack/echo) based http handler that gets the form value and calls `hashPass` with the value.
