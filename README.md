[![Coverage Status](https://goreportcard.com/badge/github.com/involvestecnologia/mole)](https://goreportcard.com/report/github.com/involvestecnologia/mole)
[![Coverage Status](https://coveralls.io/repos/github/involvestecnologia/mole/badge.svg)](https://coveralls.io/github/involvestecnologia/mole)


<img src="./icon.png" width="100" height="100">  
  
# Mole

This project was born with the objective of facilitating the analysis of the oplog, exporting the information to elasticsearch, where we can use all the visualization tools of Kibana.

## How to execute the project?

1. In the project's root directory run the following command.

```
docker-compose -f deployments/development/docker-compose.yml up -d
```

2. Access MongoDB and run the command below to enable replication between nodes.

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
            "priority": 1,
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

3. Access the kibana http://localhost:5601

```
username: elastic
password: elastic
```

4. In the menu management-> Index Lifecycle Management create a policy that stores records for 5 days and performs daily rotation.

 

5. Through the Dev tools menu create a template for the oplog index:

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
