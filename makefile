dev:
	PORT=1323 go run main.go
connect-db:
	psql -h 0.0.0.0 -p 32768 -U postgres -d bouncer_dev
