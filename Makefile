all: escape bench
escape:
	go build -gcflags '-m -l' ./main.go
bench:
	go test -bench . -benchmem -benchtime=30000000x
