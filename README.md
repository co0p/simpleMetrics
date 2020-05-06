Simple Metrics Service In Go
=============================

A in memory metrics store, that allows simple aggregations on events collected.

example binary & usage:
------------------------

build the example webserver:

    go build -o metricsServer cmd/metricsServer/main.go 

and then you can record events by using curl:
    
    curl -XGET localhost:8080/collect?label=user.visited

and then ask for an aggregation on those events with a certain sample rate:

    curl ...

Docker
-------

To build and run the example web server, you can use the provided docker file. 

Run the following command from the main directory in order to build the image: 

    docker build -t simplemetrics:latest .  

Run the following command to start the example web server listening on port 8888 in detached mode:

    docker run -d -8888:8080 simplemetrics:latest   


WIP  - brainstorming ahead
----------------------------

Event:
 - label : string
 - occured: time
 - value: int
 
e.g. Event("request.duration", 12/12/2021@12:12:00, 12)

service.collect(event)
...

Query:
- label
- time range (start, end)
- sampleRate in ms
- aggregation

Aggregation:
- sum
- min
- max (opt)
- avg (opt)

e.g:
Query("request.duration", 10/12/2021@12:12:00, 13/12/2021@12:12:00, 60*1000, "avg")


timeSeries = service.query(Query)

timeSeries: 
- [{value, startDate 0}, ...., {value, startDate N}]
 
