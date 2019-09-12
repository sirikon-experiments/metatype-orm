package db

import "github.com/sirikon-experiments/metatype-orm/scripts"

func Login()  {
	scripts.Pipe("./docker/metatype-devenv", "docker-compose", "exec", "db", "psql", "-U", "postgres", "metatype")
}
