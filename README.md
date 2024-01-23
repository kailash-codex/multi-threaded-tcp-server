# Multi Threaded TCP Server

## This algorithm demonstrates how a Multi Threaded TCP Server works

We establish a TCP Server that listens to a port 2212

Requests made to port 2212, as simple as "curl https://localhost:2212" will be 'listened to' then 'accpeted by' the TCP Server. 

Then, the request is read, (in our example, nothing is passed to the TCP Server), and the TCP Server returns a response "Hello World!".

## Try:

Establish the TCP Server

```
$ go run main.go
```

On a seperate terminal, run command

```
$ curl http://localhost:1729
```

To test the concurrency of the multi-threaded system run the command multiple times

```
$ curl http://localhost:1729 &
$ curl http://localhost:1729 &
$ curl http://localhost:1729 &
```
