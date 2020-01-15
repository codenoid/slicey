# Slicey

Collection of slice modifier in Golang (Go)

## Install

```bash
go get -u github.com/codenoid/slicey
```

## Usage

```go


bigSlice := []interface{}{"a", "b", "c", ".... so on"}
splitted := slicey.Splitl(bigSlice, 13)

c := make(chan bool)
for _, child := range splitted {
	// len(splitted) goroutine
	go DoTheJob(child, c)
}

var chans [len(splitted)]bool
for i := range chans {
   chans[i] = <- c
}
```

## Features

- [x] Split left
- [x] Split fill all
- [ ] Value sorting
- [ ] Auto split dynamically
- [ ] [Request feature](https://github.com/codenoid/slicey/issues)
