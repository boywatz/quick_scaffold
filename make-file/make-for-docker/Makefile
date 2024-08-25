run:
	docker build -t fullstack-nginx-golang-fasthttp-nuxt3 .
	docker run -d -p 80:80 fullstack-nginx-golang-fasthttp-nuxt3
stop:
	docker rm $$(docker ps -aq) -f
	docker rmi $$(docker images -q) -f
log:
	docker logs docker ps -a -q
exec:
	docker exec -it $(id) /bin/sh