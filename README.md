# go_scrap

This is an example how to scrap html using Go

## Installation
1. make sure bin directory exist, if not
```bash
mkdir bin
```
2. install gocolly package
```bash
./run get github.com/gocolly/colly/...
```

## Clean
```bash
./run -c go_scrap.go
```

## Build
```bash
./run -b go_scrap.go
```

## Run
```bash
./run -r go_scrap.go
```

## Clean, Build & Run:
```bash
./run -cbr go_scrap.go
```

## Output
result.csv will be generated.
