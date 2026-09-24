build:
	go build
server:
	go run main.go
test:
	go test
db:
	psql fw_dev
dbtest:
	psql fw_test
stat:
	git diff --stat
