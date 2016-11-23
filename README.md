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
