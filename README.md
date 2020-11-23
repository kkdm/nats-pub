# README

Simple NATS Streaming Server Client API server for publishing messages

## How to use

```bash
$ ./nats-pub -h
Usage:
  nats-pub [OPTIONS]

Application Options:
  -s, --server=   nats server url
  -c, --cluster=  nats cluster name
  -l, --log-path= log file path (default: ./server.log)
  -v, --verbose   verbose log

Help Options:
  -h, --help      Show this help message

```

### Example

```bash
$ nats-pub -s nats://example.com:4222 -c mycluster -l /var/log/app/server.log
```

#### From another terminal

```bash
$ curl localhost:8080/publish -d '{"subject": "test", "message": "12345"}'
```
