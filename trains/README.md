## Simple Go App to read NRE STOMP Feed

### Current functionality
 Reads NRE feeds and writes out to a sql lite database. 
 Can be configured to write to mongo - though don't bother if you are doing this on a raspberry pi because you can only keep a 2G database on a 32 bit OS.

### Setup instructions
1. Add a the following to `etc/trains.gcfg`:
```
[nre]
user=d3user
pass=d3password
queue=<QueueName>
```
2. Create a var dirctory
3. run ./bin/trains

### Current plans
  - Stats around trains leaving / destined for winchester
  - Tweet/Web page of Waterloo departure

### NRE Feed
  - Schedule
  - Forecast
  - Deactivation
