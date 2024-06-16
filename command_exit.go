package main

import "os"

func callbackExit(_ *config) error {
	os.Exit(0)
	return nil
}
