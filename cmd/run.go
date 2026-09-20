package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/PentestPad/subzy/runner"
	"github.com/spf13/cobra"
)

var opts = runner.Config{}

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run subzy",
	Aliases: []string{"r"},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Flag errors still show usage, but operational network errors do not.
		cmd.Root().SilenceUsage = true

		fingerprintsPath, err := runner.GetFingerprintPath()
		if err != nil {
			return err
		}
		_, statErr := os.Stat(fingerprintsPath)
		if errors.Is(statErr, fs.ErrNotExist) {
			fmt.Printf("[ * ] Fingerprints not found; saving them to %q\n",
				fingerprintsPath)
			if err := runner.DownloadFingerprints(); err != nil {
				return err
			}
		} else if statErr != nil {
			return fmt.Errorf("access fingerprints: %w", statErr)
		} else {
			// Older versions truncated this file before making the request. Repair
			// such a cache before attempting the normal integrity check.
			if _, err := runner.Fingerprints(); err != nil {
				fmt.Printf("[ ! ] Cached fingerprints are invalid; downloading a replacement: %v\n", err)
				if downloadErr := runner.DownloadFingerprints(); downloadErr != nil {
					return fmt.Errorf("replace invalid cached fingerprints: %w", downloadErr)
				}
				return runner.Process(&opts)
			}

			fmt.Printf("[ * ] Fingerprints found; checking integrity with an upstream\n")
			found, err := runner.CheckIntegrity()
			if err != nil {
				// The integrity check is maintenance, not a prerequisite for a scan.
				// Keep using the cache when the upstream is temporarily unavailable.
				fmt.Printf("[ ! ] Unable to check upstream fingerprints; using cached copy: %v\n", err)
				found = true
			}
			if !found {
				fmt.Printf("[ * ] Integrity mismatch between local and upstream fingerprints; downloading\n")
				if err := runner.DownloadFingerprints(); err != nil {
					fmt.Printf("[ ! ] Unable to update fingerprints; using cached copy: %v\n", err)
				}
			}
		}

		if err := runner.Process(&opts); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	runCmd.Flags().StringVar(&opts.Target, "target", "", "Comma separated list of domains")
	runCmd.Flags().StringVar(&opts.Targets, "targets", "", "File containing the list of subdomains")
	runCmd.Flags().StringVar(&opts.Output, "output", "", "JSON output filename")
	runCmd.Flags().BoolVar(&opts.HTTPS, "https", false, "Force https protocol if not no protocol defined for target (default false)")
	runCmd.Flags().BoolVar(&opts.VerifySSL, "verify_ssl", false, "If set to true it won't check sites with insecure SSL and return HTTP Error")
	runCmd.Flags().BoolVar(&opts.HideFails, "hide_fails", false, "Don't display failed results")
	runCmd.Flags().BoolVar(&opts.OnlyVuln, "vuln", false, "Save only vulnerable subdomains")
	runCmd.Flags().IntVar(&opts.Concurrency, "concurrency", 10, "Number of concurrent checks")
	runCmd.Flags().IntVar(&opts.Timeout, "timeout", 10, "Request timeout in seconds")
	rootCmd.AddCommand(runCmd)
}
