go build -o rc main.go
go vet -vettool=./rc ./targe/test/dir
