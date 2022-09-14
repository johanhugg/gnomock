package main

// all known presets should go right here so that they are available when
// requested over HTTP:
import (
	_ "github.com/johanhugg/gnomock/preset/cassandra"
	_ "github.com/johanhugg/gnomock/preset/cockroachdb"
	_ "github.com/johanhugg/gnomock/preset/elastic"
	_ "github.com/johanhugg/gnomock/preset/influxdb"
	_ "github.com/johanhugg/gnomock/preset/k3s"
	_ "github.com/johanhugg/gnomock/preset/kafka"
	_ "github.com/johanhugg/gnomock/preset/localstack"
	_ "github.com/johanhugg/gnomock/preset/mariadb"
	_ "github.com/johanhugg/gnomock/preset/memcached"
	_ "github.com/johanhugg/gnomock/preset/mongo"
	_ "github.com/johanhugg/gnomock/preset/mssql"
	_ "github.com/johanhugg/gnomock/preset/mysql"
	_ "github.com/johanhugg/gnomock/preset/postgres"
	_ "github.com/johanhugg/gnomock/preset/rabbitmq"
	_ "github.com/johanhugg/gnomock/preset/redis"
	_ "github.com/johanhugg/gnomock/preset/splunk"
	// new presets go here
)
