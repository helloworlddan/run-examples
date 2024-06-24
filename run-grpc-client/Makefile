test-sigterm:
	go run main.go &
	sleep 5
	kill -15 $$(ps aux --sort start_time | grep 'go run main.go' | head -n 1 | awk '{print $$2}')

