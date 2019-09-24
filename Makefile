build:
	go build
run:
	go run main.go dns.go spoe.go --debug

clean:
	rm -f ./sock haproxy-spoa-do-resolve
