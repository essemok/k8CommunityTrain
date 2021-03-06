package shutdown

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"
)

// Callback is a function to provide service shutdown.
// This function should return status of shutdown and error in case of problems.
type Callback func() (status string, err error)

// Handler handles shutting process
type Handler struct {
	logger          logrus.FieldLogger
	shutdownSignals []os.Signal
}

// NewHandler creates an instance of Handler
func NewHandler(logger logrus.FieldLogger) *Handler {
	return &Handler{
		logger:          logger,
		shutdownSignals: []os.Signal{os.Interrupt, os.Kill, syscall.SIGTERM},
	}
}

// RegisterShutdown set up a channel where we can send signal notifications and listen for the signal.
func (h *Handler) RegisterShutdown(shutdown Callback) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, h.shutdownSignals...)

	killSignal := <-interrupt
	h.logger.Infof("Got signal: %+v", killSignal)

	status, err := shutdown()
	if err != nil {
		h.logger.Fatalf("Error during shutdown: %s Status: %s\n", err.Error(), status)
		os.Exit(-1)
	}

	if killSignal == os.Kill {
		h.logger.Infof("Service was killed with status: %s", status)
	} else {
		h.logger.Infof("Service was terminated by system signal with status: %s", status)
	}
	os.Exit(0)
}

// AddShutdownSignal adds a user-defined signals to shutdown.
func (h *Handler) AddShutdownSignal(sig os.Signal) {
	h.shutdownSignals = append(h.shutdownSignals, sig)
}
