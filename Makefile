.PHONY: db
db: redis mysql mongo
	#docker rm -f $(docker ps -aq)


redis:
	docker run -d --rm --name myredis -p 6379:6379 redis:6.2.5 --requirepass hpxhJvhnrIAbX6AWd9DVLhcaZAoaGD8m

mysql:
	docker run -d --rm --name mymysql -p 3306:3306 -e MYSQL_DATABASE=test -e MYSQL_ROOT_PASSWORD=hpxhJvhnrIAbX6AWd9DVLhcaZAoaGD8m mysql:5.7.35

mongo:
	docker run -d --rm --name mymongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=hpxhJvhnrIAbX6AWd9DVLhcaZAoaGD8m mongo:5.0.2 --auth

es:
	docker run -d --rm --name myes -p 9200:9200 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms128m -Xmx256m" elasticsearch:7.14.0

kibana:
	docker run -d --rm --name mykibana -p 5601:5601 --link myes kibana:7.14.0

kafka:
	docker run -d --rm --name myzookeeper -p 2181:2181 wurstmeister/zookeeper
	docker run -d --rm --name mykafka -p 9092:9092 --link myzookeeper -e KAFKA_HEAP_OPTS="-Xms128M -Xmx256M" -e KAFKA_ZOOKEEPER_CONNECT=myzookeeper:2181 -e KAFKA_ADVERTISED_HOST_NAME=0.0.0.0 -e KAFKA_ADVERTISED_PORT=9092 wurstmeister/kafka:2.13-2.7.0

etcd:
	docker run -d --rm --name myetcd -p 2379:2379 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:3.5.0

#.PHONY: protocenv
#protocenv:
#    go mod download
#    go install github.com/gogo/protobuf/protoc-gen-gogo
#    go install github.com/gogo/protobuf/protoc-gen-gofast
#    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
