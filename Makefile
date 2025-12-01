debug:
	air

run:
	docker build -t dreamkhv/hog-bridge -f docker/Dockerfile . &&
	docker run -p 8025:8025 -p 1025:1025 -p 8080:8080 dreamkhv/hog-bridge