FROM  docker.elastic.co/elasticsearch/elasticsearch:7.5.2

COPY  ./certs/ /usr/share/elasticsearch/config/certs/

RUN echo 'elastic' | bin/elasticsearch-keystore add  -xf bootstrap.password

EXPOSE 9200 9300
