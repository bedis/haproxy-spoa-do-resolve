===================================================================
SPOE Agent (SPOA) for HAProxy to perform DNS resolutions at runtime
===================================================================

This Agent performs DNS requests that HAProxy will ask it to do.
It embeds for now 2 engines for A and PTR resolutions.

You can customize `spoe-message` for each engine with the `event` and the `args` to be sent.
You can't change the name of the arg, only its value.

The file `spoa-do-resolve.conf` provided in this repo shows configuration example to use this agent.
Corresponding HAProxy configuration is also provided.

This is recommended to use with HAProxy 2.0+.

Testing
=======

First, do the `go get` thing.

You can easily test the solution using 3 shells:

1. run the agent: `make run`
2. run haproxy: `haproxy -d -f ./haproxy.cfg`
3. run `curl` against HAProxy and observe the result in the log line sent to stdout

Building
========

First, do the `go get` thing, then simply run `make build`

Help
====

Currently supported command line arguments:

```
Usage of ./spoa-do-resolve:
  -debug
        turn on debug messages
```

Known errors or limitations
===========================

None or check issues of this project.

TODO
====

* command line argument (flag) for:
  * custom DNS servers (default uses the OS ones)
  * files for Debug, Info, Warning and Error functions
* allow HAProxy to send a list of DNS server IPs to use for this request
* new SPOE engine for "full A+PTR" resolution
* cache the DNS respones in the agent
