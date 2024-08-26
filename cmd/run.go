package cmd

import (
	"errors"
	"fmt"
	"github.com/PentestPad/subzy/runner"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
)

var opts = runner.Config{}

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run subzy",
	Aliases: []string{"r"},
	RunE: func(cmd *cobra.Command, args []string) error {
		fingerprintsPath, err := runner.GetFingerprintPath()
		if err != nil {
			return err
		}
		if _, err := os.Stat(fingerprintsPath); errors.Is(err, fs.ErrNotExist) {
			fmt.Printf("[ * ] Fingerprints not found; saving them to %q\n",
				fingerprintsPath)
			if err := runner.DownloadFingerprints(); err != nil {
				return err
			}
		} else {
			fmt.Printf("[ * ] Fingerprints found; checking integrity with an upstream\n")
			found, err := runner.CheckIntegrity()
			if err != nil {
				return err
			}
			if !found {
				fmt.Printf("[ * ] Integrity mismatch between local and upstream fingerprints; downloading\n")
				if err := runner.DownloadFingerprints(); err != nil {
					return err
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
