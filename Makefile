.PHONY: db
db: redis mysql mongo
	#docker rm -f $(docker ps -aq)

redis:
	docker run -d --rm --name myredis -p 6379:6379 redis:6.2.5 --requirepass hpxhJvhnrIAbX6AWd9DVLhcaZAoaGD8m
mysql:
	docker run -d --rm --name mymysql -p 3306:3306 -e MYSQL_DATABASE=test -e MYSQL_ROOT_PASSWORD=hpxhJvhnrIAbX6AWd9DVLhcaZAoaGD8m mysql:5.7.35
mongo:
	docker run -d --rm --name mymongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=hpxhJvhnrIAbX6AWd9DVLhcaZAoaGD8m mongo:5.0.2 --auth
