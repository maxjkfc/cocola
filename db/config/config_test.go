package config

import (
	"testing"
)

func Test_ConfigFormat(t *testing.T) {
	c := &Config{
		Host:    " dk12k3@jkgj",
		Account: "kjjk kfj akjskf",
	}

	c.Format()
}
