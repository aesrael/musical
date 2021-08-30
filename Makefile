build-run:
	go build -o main . && ./main & ./asynqmon.bin
run:
	./main & ./asynqmon.bin
