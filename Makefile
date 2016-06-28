

gaSort: *.go 
	go build -o gaSort  *.go

clean: 
	rm -f gaSort 

test: gaSort 
	./gaSort < input 
