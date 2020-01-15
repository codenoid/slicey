# Slicey

Collection of slice modifier in Golang (Go)

## Install

```bash
go get -u github.com/codenoid/slicey
```

## Usage

```go
c := make(chan bool)

splitted := slicey.Splitl(bigSlice, 13)

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
- [ ] [Create Issue](https://github.com/codenoid/slicey/issues)
