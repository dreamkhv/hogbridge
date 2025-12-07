debug: down
	docker run -d --rm --name hogbridge-mailhog -p 8025:8025 -p 1025:1025 mailhog/mailhog
	air

run: down
	docker build -t dreamkhv/hogbridge -f docker/Dockerfile . && \
	docker run --rm --name hogbridge -p 8025:8025 -p 1025:1025 -p 8080:8080 dreamkhv/hogbridge

down:
	docker stop hogbridge hogbridge-mailhog || true

publish:
	docker build -t dreamkhv/hogbridge -f docker/Dockerfile . && \
	docker push dreamkhv/hogbridge