# Bouncer
> It's called that because it check ids.

* Bouncer is a database API server that receives an array of id's and checks for
their presence in a database, leveraging the speed and efficiency of a MySQL/Postgres
server.
* Written with Go, Echo (v3) and Postgres.

## 11/23/16
- [ ] Figure out why test is panicking
  * Seems to be related to the call to `db.Query()` when the database connection
  object is embedded in the struct.

## 11/28/16
* The test seems like it was panicking because `database/sql` needed a db connection.
* Had to switch over to the testing tools provided by the Echo team.

- [ ] Have `web.go` instantiate a `PGConn` struct that has a live PG connection
to query with as an attribute
