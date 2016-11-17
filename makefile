dev:
	PORT=1323 go run main.go
connect-db:
	psql -h 0.0.0.0 -p 32768 -U postgres -d bouncer_dev
pg:
	docker build -f Dockerfile.db . \
	&& docker run --name bouncer_dev -d -P postgres:9.6.1-alpine
