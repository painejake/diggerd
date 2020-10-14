# Diggerd

Diggerd is a simple, lightweight, monitoring daemon serving the data via a JSON API written in Go designed to be super simple. Essentially a wrapper around [gopsutil](https://github.com/shirou/gopsutil). Used for simple monitoring of Raspberry Pi's and learning Go.

#### Building form source
For production release:
```sh
$ go build -o build/diggerd
```

#### Getting it running
Start the process:
```sh
$ ./diggerd
```
diggerd will now be listening on port :10000. The avaliable endpoints are:
|Endpoint|Usage|
|---|---|
|/stats|Returns all the collected system stats as JSON|
|/stats/cpu|Returns the CPU usage as a percentage|
|/stats/mem|Returns the current Memory usage in bytes|
|/stats/net|Returns the current Network usage for each interface in bytes|


### Todos

 - Expand offered data by net including error counts
 - Add disk stats output per disk basis
 - Create config options for CORS header, port binding etc

License
----

BSD