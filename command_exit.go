package main

import "os"

func callbackExit(_ *config, _ ...string) error {
	os.Exit(0)
	return nil
}
