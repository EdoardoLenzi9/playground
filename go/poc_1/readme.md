# Proof of Concept 1 - Interfaces

Make in communication a go microservice `m1` with:
1. another go microservice `m2`
2. a message broker `mb`
3. a database `db`

## [Ogen](https://ogen.dev/docs/intro)

```sh
go mod init m2
go install -v github.com/ogen-go/ogen/cmd/ogen@latest
```
