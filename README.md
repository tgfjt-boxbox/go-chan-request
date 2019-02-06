# sample with goroutine

Just a sample to request apis w/ goroutine.

## Spec

1. get top stories Ids from Hackernews API
2. get story url & title from Hackernews API
3. write `result/YYYY-MM-DD-H-i-s.md`: the links in order of top stories

## Usage

```go
go run main.go
```

## Test

```go
go test
```