# GO REST API

## Rest api test with go

### Requirements

- go
- gin package
- pq package

### Description

This project consists in a rest api that allows to make crud operations from a postgres database with jwt authentication.
The database consists of 3 tables:

- users
- events
- registrations

Users can be created and can login to get a jwt token.
With this token they can manipulate events and register/cancel to them.
