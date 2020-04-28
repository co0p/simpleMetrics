Simple Metrics Service In Go
=============================


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
 
