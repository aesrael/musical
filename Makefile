build-run:
	go build -o main . && ./main & ./asynqmon
run:
	./main & ./asynqmon
run-prod:
	./main & ./asynqmon.bin --redis-addr=redis:6379
