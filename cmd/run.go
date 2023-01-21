package cmd

import (
	"errors"
	"fmt"
	"github.com/LukaSikic/subzy/runner"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
)

var opts = runner.Settings{}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run subzy",
	RunE: func(cmd *cobra.Command, args []string) error {
		fingerprintsPath, err := runner.GetFingerprintPath()
		if err != nil {
			return err
		}
		if _, err := os.Stat(fingerprintsPath); errors.Is(err, fs.ErrNotExist) {
			fmt.Printf("[ * ] Fingerprints not found; saving them to %q\n",
				fingerprintsPath)
			if err := runner.CheckFingerprints(); err != nil {
				return err
			}
		}

		if err := runner.Process(&opts); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	runCmd.Flags().StringVarP(&opts.Target, "target", "s", "", "Comma separated list of domains")
	runCmd.Flags().StringVarP(&opts.Targets, "targets", "m", "", "File containing the list of subdomains")
	runCmd.Flags().BoolVarP(&opts.HTTPS, "https", "f", false, "Force https protocol if not no protocol defined for target (default false)")
	runCmd.Flags().BoolVarP(&opts.VerifySSL, "verify_ssl", "c", false, "If set to true it won't check sites with insecure SSL and return HTTP Error")
	runCmd.Flags().BoolVar(&opts.HideFails, "hide_fails", false, "Don't display failed results")
	runCmd.Flags().IntVar(&opts.Concurrency, "concurrency", 10, "Number of concurrent checks")
	runCmd.Flags().IntVar(&opts.Timeout, "timeout", 10, "Request timeout in seconds")
	rootCmd.AddCommand(runCmd)
}
