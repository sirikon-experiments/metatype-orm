// +build mage

package main

import (
	// mage:import
	_ "github.com/sirikon-experiments/metatype-orm/scripts"
	// mage:import db
	_ "github.com/sirikon-experiments/metatype-orm/scripts/db"
	// mage:import devenv
	_ "github.com/sirikon-experiments/metatype-orm/scripts/devenv"
)
