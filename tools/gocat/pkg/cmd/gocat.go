package cmd

// Gocat receives a filepath and outputs the file's contents
// to stdout. It is used to test fopen latency on machines
// suspected of slow io.

import (
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/spf13/cobra"
)

type GocatOpts struct {
	filepath string
	stats    bool

	stdout io.Writer
	stderr io.Writer
}

func NewCmdGocat(parent string, out, err io.Writer) *cobra.Command {
	opts := &GocatOpts{
		stdout: out,
		stderr: err,
	}

	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%[1]s <filepath> [--stats]", parent),
		Short: "output file contents to stdout",
		Long:  "output file contents to stdout",
		RunE: func(c *cobra.Command, args []string) error {
			if err := opts.Complete(c, args); err != nil {
				return err
			}
			if err := opts.Validate(); err != nil {
				return err
			}

			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.filepath, "filepath", "f", "", "path of file whose contents should be printed.")
	cmd.Flags().BoolVarP(&opts.stats, "stats", "s", false, "display statistics of file io, instead of printing file contents.")

	return cmd
}

func (o *GocatOpts) Complete(c *cobra.Command, args []string) error {
	return nil
}

func (o *GocatOpts) Validate() error {
	if len(o.filepath) == 0 {
		return fmt.Errorf("a filepath is required")
	}

	return nil
}

func (o *GocatOpts) Run() error {
	timeStart := time.Now()

	b, err := ioutil.ReadFile(o.filepath)
	if err != nil {
		return err
	}

	if o.stats {
		ellapsed := time.Since(timeStart)
		fmt.Printf("opened and read %v bytes in %s\n", len(b), ellapsed)
		return nil
	}

	fmt.Fprintf(o.stdout, "%s", string(b))
	return nil
}
