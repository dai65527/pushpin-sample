run server and 

subscribe channel

```
$ curl http://localhost:7999/stream/mychannel
```

post to channel from another terminal

```
$ curl -d '{"sender_name":"dnakano","message":"hello!\n"}' -H "content-type: application/json" http://localhost:8080/message/mychannel
"eyJzZW5kZXJfbmFtZSI6ImRuYWthbm8iLCJtZXNzYWdlIjoiaGVsbG8hXG4ifQ=="

$ curl -d '{"sender_name":"dnakano","message":"mychannel!\n"}' -H "content-type: application/json" http://localhost:8080/message/mychannel
"eyJzZW5kZXJfbmFtZSI6ImRuYWthbm8iLCJtZXNzYWdlIjoibXljaGFubmVsIVxuIn0="
```

subscriber will get messages.

```
$ curl http://localhost:7999/stream/mychannel
stream openedhello!
mychannel!
```
