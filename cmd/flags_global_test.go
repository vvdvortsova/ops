package cmd_test

import (
	"testing"

	"github.com/nanovms/ops/cmd"
	"github.com/nanovms/ops/lepton"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestCreateGlobalFlags(t *testing.T) {

	flagSet := pflag.NewFlagSet("test", 0)

	cmd.PersistGlobalCommandFlags(flagSet)

	flagSet.Set("show-debug", "true")
	flagSet.Set("show-errors", "true")
	flagSet.Set("show-warnings", "true")

	globalFlags := cmd.NewGlobalCommandFlags(flagSet)

	assert.Equal(t, globalFlags.ShowDebug, true)
	assert.Equal(t, globalFlags.ShowErrors, true)
	assert.Equal(t, globalFlags.ShowWarnings, true)
}

func TestGlobalFlagsMergeToConfig(t *testing.T) {
	flagSet := pflag.NewFlagSet("test", 0)

	cmd.PersistGlobalCommandFlags(flagSet)

	flagSet.Set("show-debug", "true")
	flagSet.Set("show-errors", "false")
	flagSet.Set("show-warnings", "true")

	globalFlags := cmd.NewGlobalCommandFlags(flagSet)

	config := &lepton.Config{}

	err := globalFlags.MergeToConfig(config)

	assert.Nil(t, err)

	assert.Equal(t, config, &lepton.Config{
		RunConfig: lepton.RunConfig{
			ShowDebug:    true,
			ShowErrors:   false,
			ShowWarnings: true,
		},
	})
}
