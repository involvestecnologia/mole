[![Coverage Status](https://goreportcard.com/badge/github.com/involvestecnologia/mole)](https://goreportcard.com/report/github.com/involvestecnologia/mole)


<img src="./icon.png" width="100" height="100">  
  
# Mole

Basically the project consists of replicating information from the oplog to elasticsearch, a stack developed to analyze information, there we can analyze the oplog to identify bottlenecks in MongoDB.

## How to execute the project on your local machine ?

To run the project locally you will need to perform 4 steps:

- Run a MongoDB cluster with the replication system enabled.
- Run an elasticsearch cluster.
- Run Kibana.
- Run the mole to start replicating the oplog.

To facilitate this process, I created a docker-compose that starts all this infrastructure, you will only need to activate some settings:

- In the project's root directory start the infrastructure by running the command below on the terminal:

```
docker-compose -f deployments/development/docker-compose.yml up -d
```

- Access MongoDB and activate replication using the command:

```
rs.initiate({
    "_id": "rs",
    "version": 1,
    "members": [
        {
            "_id": 0,
            "host": "localhost:27017",
            "priority": 1,
            "votes":1
        },
        {
            "_id": 1,
            "host": "localhost:27018",
            "priority": 0,
            "votes":1
        },
        {
            "_id": 2,
            "host": "localhost:27019",
            "arbiterOnly": true,
            "priority": 0,
            "votes":1
        }
    ],settings: {chainingAllowed: true}
});
```

- Configure Kibana
1. Access the link http: localhost: 5601 with the user "elastic" and the password "elastic".
2. In the "Management -> Index Lifecycle Management" menu, create a rotation policy with the name "" 5-day-storage-with-daily-rotation ".
3. In the "Dev Tools" menu, create a mapping for the oplog index using the command below:

```
PUT _template/oplog
{
    "index_patterns": [
        "oplog-*"
    ],
    "settings": {
        "number_of_shards": 2,
        "number_of_replicas": 1,
        "index.lifecycle.name": "5-day-storage-with-daily-rotation",
        "index.lifecycle.rollover_alias": "oplog"
    },
    "mappings": {
        "_source": {
            "enabled": true
        },
        "properties": {
            "timestamp": {
                "type": "date"
            },
            "database": {
                "type": "keyword"
            },
            "collection": {
                "type": "keyword"
            },
            "operation": {
                "type": "keyword"
            },
            "query": {
                "type": "text"
            }
        }
    }
}

PUT oplog-000001
PUT oplog-000001/_aliases/oplog

```
- Run the Mole application to start replication

```
go run main.go
```

## How to deploy in production ?

To run the Mole application, just replace the compose parameters and execute:

```
version: '2.2'
services:
  mole:
    image: involvestecnologia/mole:latest
    container_name: mole
    environment:
      MONGO_URI: mongodb://localhost:27017
      ELASTICSEARCH_HOSTS: http://localhost:9200
      ELASTICSEARCH_USERNAME: elastic
      ELASTICSEARCH_PASSWORD: elastic
      ELASTICSEARCH_SOURCE: oplog
      ELASTICSEARCH_BATCH_SIZE: 10000
      LOGSTASH_URL: localhost:12201
```
