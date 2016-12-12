# REST Simulator for WBC

## Simulated APIs

- POST /profiles/enterprises
  - 201 - Succeed.
  - 400 - Error: Expecting Profile Name.
  - 409 - Error: Duplicate Profile.
- GET /profiles/enterprises/{profile}
  - 200 - Ok.
  - 404 - Error: Not Found.

## Compile and Run

- Download and install [Golang](https://golang.org/dl/).
- Set enviornment variable GOPATH.
  - Windows: `GOPATH=C:\gopath`
  - Linux: `export GOPATH=$HOME/go`
- Get 3rd party library.
  - `go get github.com/julienschmidt/httprouter`
- Run server.
  - `go run main.go`

## Example

```sh
$ curl -i -H "Content-Type: application/json" \
>      -d '{"profile": {"profileName": "bar"}, "enterpriseConfigParams": {"dbUsername": "barbar"}}' \
>      -X POST http://localhost:8888/profiles/enterprises
HTTP/1.1 201 Created
Date: Mon, 12 Dec 2016 16:41:01 GMT
Content-Length: 8
Content-Type: text/plain; charset=utf-8

Succeed.

$ curl -i -H "Content-Type: application/json" \
>      -d '{"profile": {"profileName": "foo"}, "enterpriseConfigParams": {"dbUsername": "foofoo"}}' \
>      -X POST http://localhost:8888/profiles/enterprises
HTTP/1.1 201 Created
Date: Mon, 12 Dec 2016 16:41:01 GMT
Content-Length: 8
Content-Type: text/plain; charset=utf-8

Succeed.

$ curl -i -H "Content-Type: application/json" \
>      -d '{"profile": {"profileName": "bar"}, "enterpriseConfigParams": {"dbUsername": "barbar"}}' \
>      -X POST http://localhost:8888/profiles/enterprises
HTTP/1.1 409 Conflict
Date: Mon, 12 Dec 2016 16:41:02 GMT
Content-Length: 25
Content-Type: text/plain; charset=utf-8

Error: Duplicate Profile.

$ curl -i -H "Accept: application/json" -H "Content-Type: application/json" \
>      -X GET http://localhost:8888/profiles/enterprises/foo
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 12 Dec 2016 16:30:26 GMT
Content-Length: 549

{"profile":{"profileName":"foo","accessPermission":"","lastUpdated":"2016/12/13 00:30:25"},"enterpriseConfigParams":{"databasePlatform":"","connectString":"","tableOwner":"","dbUsername":"foofoo","dbUserpasswd":"","odbcDataSource":"","sqlDatabase":"","sqlServer":"","db2DatabaseAlias":"","db2CurrentSQLID":"","encrypt":"","siebelEncryption":"","keyFileName":"","keyFilePassword":"","peerAuth":"","peerCertValidation":"","caCertFileName":"","certFileNameServer":"","requestServer":"","secAdptMode":"","serverFileSystem":"","cloudRegistryAddress":""}}

$ curl -i -H "Accept: application/json" -H "Content-Type: application/json" \
>      -X GET http://localhost:8888/profiles/enterprises/foobar
HTTP/1.1 404 Not Found
Date: Mon, 12 Dec 2016 16:31:06 GMT
Content-Length: 17
Content-Type: text/plain; charset=utf-8

Error: Not Found.
```