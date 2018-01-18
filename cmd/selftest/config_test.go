package main

import (
	"testing"
)

func testConfig() Config {
	// TODO: populate test config
	return Config{}
}

func TestConfigValidate(t *testing.T) {
	for _, c := range []struct {
		Expect    string
		Configure func(*Config)
	}{
		// TODO: add test cases
		{
			Configure: func(config *Config) {},
		},
	} {
		config := testConfig()
		c.Configure(&config)

		if err := config.Validate(); len(c.Expect) == 0 {
			if err == nil {
				continue
			}
			t.Fatalf("unexpected error: " + err.Error())
		} else {
			if err != nil {
				if expected, got := c.Expect, err.Error(); expected != got {
					t.Fatalf("expected %s, got %s", expected, got)
				}
				continue
			}
			if err == nil {
				t.Fatalf("missing expected error: " + c.Expect)
			}
		}
	}
}
