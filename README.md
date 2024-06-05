# tinder-match

## Description
HTTP server for the Tinder matching system. The HTTP server must support the following three APIs:
1. AddSinglePersonAndMatch : Add a new user to the matching system and find any possible matches for the new user.
2. RemoveSinglePerson : Remove a user from the matching system so that the user cannot be matched anymore.
3. QuerySinglePeople : Find the most N possible matched single people, where N is a request parameter.

Here is the matching rule:
- A single person has four input parameters: name, height, gender, and number of
wanted dates.
- Boys can only match girls who have lower height. Conversely, girls match boys who
are taller.
- Once the girl and boy match, they both use up one date. When their number of dates
becomes zero, they should be removed from the matching system.


## Setup and Run
### with go
```
# compile and run
$ make run-bin
```

### with docker
```
# build image and run
$ make run-docker
```

## System Design


## Complexity Analysis


## file structure
```
├── cmd
│   └── server      # server entry point
├── docs
│   └── api         # API documentation
├── internal        # sub-folder no need to export
│   ├── config      # config package
│   ├── controller  # controller/handler package
│   │   └── dto     # controller data transfer object
│   ├── http        # gin engine for http 
│   ├── logger      # logger package
│   ├── route       # route path register package
│   ├── server      # app entry point
│   └── service     # service business logic package
├── k6              # k6 load test script
└── model           # entity model able to export
```

## API Documentation
for more detail, you can
find the API documentation in [`docs/api`](https://github.com/jynychen/tinder-match/tree/main/docs/api) folder

* AddSinglePersonAndMatch
    - add a new user and find any possible matches
    - POST /api/v1/people
    - Request
    ```
    {
        "name": "string",
        "height": 0,
        "gender": "string",
        "wantedDates": 0
    }
    ```
    - Response
    ```
    {
       "matched": [ ... ]
    }
    ```

* RemoveSinglePerson
    - Remove a user from the matching system
    - DELETE /api/v1/people/:id

* QuerySinglePeople
    -  Query a list of users from the matching system
    - GET /api/v1/people?limit=0
    - Response
    ```
    {
        "people": [ ... ]
    }
    ```


## TBD/TODO
- [ ] event driven archt. with mq/channel for async processing
- [ ] add user first and matching later for API quick response
- [ ] lock free data structure for concurrent access
- [ ] refactor red black tree to support same key nodes
- [ ] log matching process for validation and failover recovery
- [ ] persist data store for data recovery
- [ ] add unit test / mock for service
- [ ] add integration test for API
- [ ] level log for different log level
- [ ] debug/release build


## Done
- [x] AddSinglePersonAndMatch
- [x] RemoveSinglePerson
- [x] QuerySinglePeople
- [x] unit test for matching service
- [x] matching service test coverage 98.8%
- [x] golangci-lint pass
- [x] k6 load test script
- [x] containerize with docker
- [x] API documentation


## k6 test result
```
$ k6 run k6/match.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: k6/match.js
        output: -

     scenarios: (100.00%) 1 scenario, 10000 max VUs, 1m10s max duration (incl. graceful stop):
              * default: Up to 10000 looping VUs for 40s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


     ✓ is status 200

     checks.........................: 100.00% ✓ 1202716      ✗ 0      
     data_received..................: 526 MB  13 MB/s
     data_sent......................: 419 MB  10 MB/s
     http_req_blocked...............: avg=5.07µs  min=234ns   med=857ns    max=61.86ms  p(90)=2.07µs  p(95)=2.9µs  
     http_req_connecting............: avg=2.96µs  min=0s      med=0s       max=51.06ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.33ms  min=18.4µs  med=103.26µs max=109.45ms p(90)=3.64ms  p(95)=7.96ms 
       { expected_response:true }...: avg=1.33ms  min=18.4µs  med=103.26µs max=109.45ms p(90)=3.64ms  p(95)=7.96ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 2104753
     http_req_receiving.............: avg=29.32µs min=2.88µs  med=10.44µs  max=71.78ms  p(90)=20.91µs p(95)=26.86µs
     http_req_sending...............: avg=30.32µs min=1.25µs  med=4.38µs   max=93.31ms  p(90)=10.03µs p(95)=15.2µs 
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s       max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.27ms  min=11.69µs med=85µs     max=78.11ms  p(90)=3.49ms  p(95)=7.7ms  
     http_reqs......................: 2104753 51541.292945/s
     iteration_duration.............: avg=1.01s   min=1s      med=1s       max=1.35s    p(90)=1.04s   p(95)=1.08s  
     iterations.....................: 300679  7363.041849/s
     vus............................: 43      min=43         max=10000
     vus_max........................: 10000   min=10000      max=10000


running (0m40.8s), 00000/10000 VUs, 300679 complete and 0 interrupted iterations
default ✓ [======================================] 00000/10000 VUs  40s
```