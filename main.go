package main

import (
	"fmt"
	"os"

	"github.com/evcc-io/evcc/cmd"
)

// main is the entry point for the evcc application.
// evcc is an EV Charge Controller that manages electric vehicle charging
// based on solar surplus, grid pricing, and other configurable strategies.
//
// Personal fork: using this to manage charging for my VW ID.4 with a
// SolarEdge inverter. Main config lives in ~/.evcc/evcc.yaml.
//
// Note: exit code 2 is used for configuration/usage errors (set by cobra),
// exit code 1 is used for runtime errors caught here.
//
// TODO: look into contributing SolarEdge SE7600H inverter fixes upstream.
// TODO: investigate min-soc behavior when grid price spikes overnight.
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
