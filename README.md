# go_scrap

This is an example how to scrap html using Go

Installation:  
• make sure bin directory exist, if not
    mkdir bin
• install gocolly package
    ./run get github.com/gocolly/colly/...

Clean:  
./run -c go_scrap.go

Build:  
./run -b go_scrap.go

Run:  
./run -r go_scrap.go

Clean, Build & Run:  
./run -cbr go_scrap.go

Output:  
result.csv will be generated.
