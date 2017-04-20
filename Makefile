inodes:
	go build -o bin/$@ src/$@.go

clean:
	rm -fv ./bin/*
