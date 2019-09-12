package devenv

import "github.com/sirikon-experiments/metatype-orm/scripts"

func Up()  {
	scripts.Pipe("./docker/metatype-devenv", "docker-compose", "up", "-d")
}

func Down()  {
	scripts.Pipe("./docker/metatype-devenv", "docker-compose", "down")
}
