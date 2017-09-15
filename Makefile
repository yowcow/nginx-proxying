all:
	docker-compose up -d && docker-compose logs -f

stop:
	docker-compose stop

clean:
	docker-compose down

.PHONY: all stop clean
