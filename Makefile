all:
	docker-compose up --build -d && docker-compose logs -f

stop:
	docker-compose stop

clean:
	docker-compose down

test:
	go test

.PHONY: all stop clean test
