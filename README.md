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


## TBD/TODO


## Done
- [x] AddSinglePersonAndMatch
- [x] RemoveSinglePerson
- [x] QuerySinglePeople
- [x] unit test for matching service
- [x] matching service test coverage 98.8%
- [x] golangci-lint pass
- [x] k6 load test script
- [x] containerize with docker