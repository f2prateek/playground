# playground

[![GoDoc](https://godoc.org/github.com/f2prateek/playground?status.svg)](https://godoc.org/github.com/f2prateek/playground)

Package playground provides a single method `Time` that returns the fixed current time in the go playground.

```go
t := playground.Time()
fmt.Println(t.Format(time.RFC1123))
// Output: Tue, 10 Nov 2009 23:00:00 UTC
```
