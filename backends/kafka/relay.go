package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber-schemas/build/go/protos/records"
	"github.com/pkg/errors"

	ktypes "github.com/batchcorp/plumber/backends/kafka/types"
	"github.com/batchcorp/plumber/stats"
	"github.com/batchcorp/plumber/util"
)

const (
	RetryReadInterval = 5 * time.Second
)

var (
	ErrMissingTopic = errors.New("You must specify at least one topic")
)

// Relay sets up a new Kafka relayer
func (k *Kafka) Relay(ctx context.Context, relayOpts *opts.RelayOptions, relayCh chan interface{}, errorCh chan *records.ErrorRecord) error {
	if err := validateRelayOptions(relayOpts); err != nil {
		return errors.Wrap(err, "unable to verify options")
	}

	// TODO: Why is this here? ~ds 09.10.21
	k.log.Infof("Relaying Kafka messages from '%s' topic(s) -> '%s'",
		relayOpts.Kafka.Args.Topics, relayOpts.XBatchshGrpcAddress)

	// TODO: What is this? ~ds 09.10.21
	//k.log.Infof("HTTP server listening on '%s'", k.connOpts.Relay.HTTPListenAddress)

	reader, err := NewReaderForRelay(k.dialer, k.connArgs, relayOpts.Kafka.Args)
	if err != nil {
		return errors.Wrap(err, "unable to create new reader")
	}

	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			// Shutdown cancelled, exit so we don't spam logs with context cancelled errors
			if err == context.Canceled {
				k.log.Info("Received shutdown signal, exiting relayer")
				return nil
			}

			stats.Mute("kafka-relay-consumer")
			stats.Mute("kafka-relay-producer")

			stats.IncrPromCounter("plumber_read_errors", 1)

			wrappedErr := fmt.Errorf("unable to read kafka message: %s; retrying in %s", err, RetryReadInterval)
			util.WriteError(k.log, errorCh, wrappedErr)

			time.Sleep(RetryReadInterval)

			continue
		}

		stats.Incr("kafka-relay-consumer", 1)

		k.log.Debugf("Writing Kafka message to relay channel: %s", msg.Value)

		relayCh <- &ktypes.RelayMessage{
			Value:   &msg,
			Options: &ktypes.RelayMessageOptions{},
		}
	}
}

// validateRelayOptions ensures all required relay options are present
func validateRelayOptions(relayOpts *opts.RelayOptions) error {
	if relayOpts == nil {
		return errors.New("relay opts cannot be nil")
	}

	if relayOpts.Kafka == nil {
		return errors.New("kafka opts cannot be nil")
	}

	if relayOpts.Kafka.Args == nil {
		return errors.New("kafka args cannot be nil")
	}

	if len(relayOpts.Kafka.Args.Topics) == 0 {
		return ErrMissingTopic
	}

	return nil
}
