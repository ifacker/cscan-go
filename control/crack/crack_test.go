package crack

import (
	"cscan/config"
	"testing"
)

func TestStartCrack(t *testing.T) {
	StartCrack(&config.IpOptions{})
}
