docker-run:
	docker build -t app .
	docker run -d -p 80:80 app

docker-stop:
	docker rm $$(docker ps -aq) -f
	docker rmi $$(docker images -q) -f

docker-log:
	docker logs docker ps -a -q

docker-exec:
	docker exec -it $(id) /bin/sh
