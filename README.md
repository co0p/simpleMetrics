Simple Metrics Service In Go
=============================

A in memory metrics store, that allows simple aggregations on events collected.

example binary & usage:
------------------------

build the example webserver:

    go build -o metricsServer cmd/metricsServer/main.go 

and then you can record events by using curl:
    
    curl ... 

and then ask for an aggregation on those events with a certain sample rate:

    curl ...

docker build if you have no golang installed (todo):

    docker build -t ... 


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
 
