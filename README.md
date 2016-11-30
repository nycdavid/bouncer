# Bouncer
> It's called that because it check ids.

* Bouncer is a database API server that receives an array of id's and checks for
their presence in a database, leveraging the speed and efficiency of a MySQL/Postgres
server.
* Written with Go, Echo (v2.2.0) and Postgres.

## 11/23/16
- [ ] Figure out why test is panicking
  * Seems to be related to the call to `db.Query()` when the database connection
  object is embedded in the struct.

## 11/28/16
* The test seems like it was panicking because `database/sql` needed a db connection.
* Had to switch over to the testing tools provided by the Echo team.

- [x] Have `web.go` instantiate a `PGConn` struct that has a live PG connection
to query with as an attribute

## 11/29/16
* ~~How do we get the Post handler function to gain access to the db connection without
having to make multiple connections to PG?~~

- [x] Test if the request is actually being made in the test by making a dead simple Post handler

## 11/30/16
* `rec.Body` response from the http recorder is returning some weird nonsense
string and not the json object that we expect. `eyJtYXRjaGVkQ291bnQiOjEsIm1hdGNoZWRJZHMiOlsxLDJdfQ==`
