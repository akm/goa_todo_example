# goa TODO example app

## Run server

```
go run ./cmd/apisvr
```

## HTTP API call

### List

```
$ curl -i http://localhost:8000/api/todos
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 15 Oct 2023 08:29:04 GMT
Content-Length: 34

{"items":[],"total":0,"offset":0}
```

### Create

```
$ curl -i -X POST http://localhost:8000/api/todos --data-binary '{"title":"make check list", "state": "open"}'
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 15 Oct 2023 08:30:00 GMT
Content-Length: 132

{"id":1,"title":"make check list","State":"open","created_at":"2023-10-15T17:30:00+09:00","updated_at":"2023-10-15T17:30:00+09:00"}
```

```
$ curl -i -X POST http://localhost:8000/api/todos --data-binary '{"title":"go to shops", "state": "open"}'
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 15 Oct 2023 08:30:32 GMT
Content-Length: 128

{"id":2,"title":"go to shops","State":"open","created_at":"2023-10-15T17:30:32+09:00","updated_at":"2023-10-15T17:30:32+09:00"}
```

### Show

```
$ curl -i -X GET http://localhost:8000/api/todos/1                                                            
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 15 Oct 2023 08:30:59 GMT
Content-Length: 132

{"id":1,"title":"make check list","State":"open","created_at":"2023-10-15T17:30:00+09:00","updated_at":"2023-10-15T17:30:00+09:00"}
```

```
$ curl -i -X GET http://localhost:8000/api/todos/2
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 15 Oct 2023 08:31:04 GMT
Content-Length: 128

{"id":2,"title":"go to shops","State":"open","created_at":"2023-10-15T17:30:32+09:00","updated_at":"2023-10-15T17:30:32+09:00"}
```

### Update

```
$ curl -i -X PUT http://localhost:8000/api/todos/1 --data-binary '{"body": {"title":"make check list", "state": "closed"}}'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 15 Oct 2023 08:31:50 GMT
Content-Length: 134

{"id":1,"title":"make check list","State":"closed","created_at":"2023-10-15T17:30:00+09:00","updated_at":"2023-10-15T17:31:50+09:00"}
```


### List (again)

```
$ curl -i http://localhost:8000/api/todos
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 15 Oct 2023 08:32:27 GMT
Content-Length: 295

{"items":[{"id":1,"title":"make check list","State":"closed","created_at":"2023-10-15T17:30:00+09:00","updated_at":"2023-10-15T17:31:50+09:00"},{"id":2,"title":"go to shops","State":"open","created_at":"2023-10-15T17:30:32+09:00","updated_at":"2023-10-15T17:30:32+09:00"}],"total":2,"offset":0}
```


### Delete

```
$ curl -i -X DELETE http://localhost:8000/api/todos/1                                                                   
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 15 Oct 2023 08:32:54 GMT
Content-Length: 134

{"id":1,"title":"make check list","State":"closed","created_at":"2023-10-15T17:30:00+09:00","updated_at":"2023-10-15T17:31:50+09:00"}
```

```
$ curl -i http://localhost:8000/api/todos             
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 15 Oct 2023 08:32:57 GMT
Content-Length: 161

{"items":[{"id":2,"title":"go to shops","State":"open","created_at":"2023-10-15T17:30:32+09:00","updated_at":"2023-10-15T17:30:32+09:00"}],"total":1,"offset":0}
```



## gRPC API call

```
$ go run ./cmd/apisvr-cli -url grpc://localhost:8080 todos list
{
    "Items": [
        {
            "ID": 2,
            "Title": "go to shops",
            "State": "open",
            "CreatedAt": "2023-10-15T17:30:32+09:00",
            "UpdatedAt": "2023-10-15T17:30:32+09:00"
        },
        {
            "ID": 3,
            "Title": "pay",
            "State": "open",
            "CreatedAt": "2023-10-15T17:44:07+09:00",
            "UpdatedAt": "2023-10-15T17:44:07+09:00"
        }
    ],
    "Total": 2,
    "Offset": 0
}
```

Run `go run ./cmd/apisvr-cli --help` for more detail.


## gRPC API call from node.js client

### list

```
$ npx ts-node src/todos/list.ts
{
  id: 2,
  title: 'go to shops',
  state: 'open',
  createdAt: '2023-10-15T17:30:32+09:00',
  updatedAt: '2023-10-15T17:30:32+09:00'
}
{
  id: 3,
  title: 'pay',
  state: 'open',
  createdAt: '2023-10-15T17:44:07+09:00',
  updatedAt: '2023-10-15T17:44:07+09:00'
}
```

### show

```
$ npx ts-node src/todos/show.ts 2 
{
  id: 2,
  title: 'go to shops',
  state: 'open',
  createdAt: '2023-10-15T17:30:32+09:00',
  updatedAt: '2023-10-15T17:30:32+09:00'
}
```

### create

```
$ npx ts-node src/todos/create 'write to the notebook' opened 
Error: 2 UNKNOWN: value of message.State must be one of "open", "closed" but got value "opened"
(snip)
    at Module.load (node:internal/modules/cjs/loader:1091:32)
    at Function.Module._load (node:internal/modules/cjs/loader:938:12)
    at Function.executeUserEntryPoint [as runMain] (node:internal/modules/run_main:83:12) {
  code: 2,
  details: 'value of message.State must be one of "open", "closed" but got value "opened"',
  metadata: Metadata {
    internalRepr: Map(2) {
      'content-type' => [Array],
      'grpc-status-details-bin' => [Array]
    },
    options: {}
  }
}
```

```
$ npx ts-node src/todos/create 'write to the notebook' open 
{
  id: 4,
  title: 'write to the notebook',
  state: 'open',
  createdAt: '2023-10-15T20:20:44+09:00',
  updatedAt: '2023-10-15T20:20:44+09:00'
}
```


### update

```
$ npx ts-node src/todos/update 4 'write to the notebook' closed 
{
  id: 4,
  title: 'write to the notebook',
  state: 'closed',
  createdAt: '2023-10-15T20:20:44+09:00',
  updatedAt: '2023-10-15T20:21:18+09:00'
}
```

### delete

```
$ npx ts-node src/todos/delete.ts 3
{
  id: 3,
  title: 'pay',
  state: 'open',
  createdAt: '2023-10-15T17:44:07+09:00',
  updatedAt: '2023-10-15T17:44:07+09:00'
}
```


```
$ npx ts-node src/todos/list.ts
{
  id: 2,
  title: 'go to shops',
  state: 'open',
  createdAt: '2023-10-15T17:30:32+09:00',
  updatedAt: '2023-10-15T17:30:32+09:00'
}
{
  id: 4,
  title: 'write to the notebook',
  state: 'closed',
  createdAt: '2023-10-15T20:20:44+09:00',
  updatedAt: '2023-10-15T20:21:18+09:00'
}
```
