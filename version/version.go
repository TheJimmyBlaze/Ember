package version

import (
	"fmt"
)

var (
	BuildVersion = "dev-version"
	BuildHash    = "dev-version"
	BuildTime    = "dev-version"
)

func PrintVersionInformation() {
	fmt.Printf("Ember - X.509 Crypto Service - %s\n", BuildVersion)
	fmt.Printf("Build Time: %s\n", BuildTime)
	fmt.Printf("Build Hash: %s\n", BuildHash)
}
