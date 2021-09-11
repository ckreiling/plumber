package plumber

import (
	"context"

	"github.com/batchcorp/plumber-schemas/build/go/protos/records"
	"github.com/batchcorp/plumber/backends"
	"github.com/batchcorp/plumber/printer"
	"github.com/batchcorp/plumber/validate"
	"github.com/pkg/errors"
)

// HandleReadCmd handles CLI read mode
func (p *Plumber) HandleReadCmd() error {
	if err := validate.ReadConfig(p.CLIOptions.Read); err != nil {
		return errors.Wrap(err, "unable to validate read options")
	}

	backend, err := backends.New(p.CLIOptions.Global.XBackend, p.cliConnOpts)
	if err != nil {
		return errors.Wrap(err, "unable to create new backend")
	}

	resultCh := make(chan *records.ReadRecord, 1)
	errorCh := make(chan *records.ErrorRecord, 1)

	// backend.Read() blocks
	go func() {
		if err := backend.Read(context.Background(), p.CLIOptions.Read, resultCh, errorCh); err != nil {
			p.log.Errorf("unable to complete read for backend '%s': %s", backend.Name(), err)
		}
	}()

MAIN:
	for {
		var err error

		select {
		case msg := <-resultCh:
			p.log.Debug("HandleReadCmd: received message on resultCh")

			err = backend.DisplayMessage(msg)
		case errorMsg := <-errorCh:
			err = backend.DisplayError(errorMsg)
		}

		if err != nil {
			printer.Errorf("unable to display message with '%s' backend: %s", backend.Name(), err)
		}

		if !p.CLIOptions.Read.Continuous {
			break MAIN
		}
	}

	return nil
}
