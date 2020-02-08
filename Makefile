build:
	go build

clean:
	rm -rf json_rpc_server

up:
	docker-compose -f integration_tests/docker-compose.yml up -d

ps:
	docker-compose -f integration_tests/docker-compose.yml ps

down:
	docker-compose -f integration_tests/docker-compose.yml down

run:
	./json_rpc_server
