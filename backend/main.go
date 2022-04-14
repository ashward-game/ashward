package main

import (
	"log"
	"orbit_nft/cmd"
	"orbit_nft/contract/rpc"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	envName := os.Getenv("ORBIT_ENV")
	if envName != "" {
		if _, err = os.Stat(".env." + envName); err == nil {
			err = godotenv.Load(".env." + os.Getenv("ORBIT_ENV"))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Printf("Cannot load file %v: %v, ignored", envName, err)
		}
	} else {
		log.Println("No ORBIT_ENV is specified, ignored")
	}
}

func main() {
	rpc.Initialize(os.Getenv("RPC_CONFIG_FILE"))
	cmd.Execute()
}
