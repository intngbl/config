package config

import (
	"testing"
)

func TestAutoOpenFile(t *testing.T) {
	if Config == nil {
		t.Fatal("Config should not be nil.")
	}

	if _, err := Config.String("service-1", "url"); err != nil {
		t.Fatal("service-1/url should hold a value.")
	}

	if _, err := Config.String("service-1", "url-2"); err == nil {
		t.Fatal("service-1/url-2 should NOT hold any value.")
	}
}
