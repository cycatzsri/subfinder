package runner

import (
	"errors"
	"os"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/subfinder/v2/pkg/passive"
	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
)

const banner = ``

// Version is the current version of subfinder
const Version = `v2.5.3`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("", banner)
	gologger.Print().Msgf("\t\t")

	gologger.Print().Msgf("")
	gologger.Print().Msgf("")
	gologger.Print().Msgf("")
}

// loadProvidersFrom runs the app with source config
func (options *Options) loadProvidersFrom(location string) {
	if len(options.AllSources) == 0 {
		options.AllSources = passive.DefaultAllSources
	}
	if len(options.Recursive) == 0 {
		options.Recursive = passive.DefaultRecursiveSources
	}
	// todo: move elsewhere
	if len(options.Resolvers) == 0 {
		options.Recursive = resolve.DefaultResolvers
	}
	if len(options.Sources) == 0 {
		options.Sources = passive.DefaultSources
	}

	options.Providers = &Providers{}
	// We skip bailing out if file doesn't exist because we'll create it
	// at the end of options parsing from default via goflags.
	if err := options.Providers.UnmarshalFrom(location); isFatalErr(err) && !errors.Is(err, os.ErrNotExist) {
		gologger.Fatal().Msgf("Could not read providers from %s: %s\n", location, err)
	}
}
