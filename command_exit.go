package main

import "os"

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}