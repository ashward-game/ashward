package util

import "os"

func GetEnv() string {
	env := os.Getenv("ORBIT_ENV")
	if env != "" {
		return env
	}

	return "development" // default
}
