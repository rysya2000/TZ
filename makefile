.PHONY: run
run:
	docker run -d -p 27017:27017 --name mongo mongo:latest && docker-compose up --build


.DEFAULT_GOAL := run