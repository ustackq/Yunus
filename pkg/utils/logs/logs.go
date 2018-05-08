package logs

import (
	"flag"
	"log"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"github.com/ustack/Yunus/src/app/backend/pkg/utils/wait"
)

var logFlushFreq = pflag.Duration("log-flush-frequency", 5*time.Second, "Maxium num's between log flushes")

// TODO: need design standard log dir
func init() {
	flag.Set("logtostderr", "true")
}

// GlogWriter serves as bridge between stanard log with glog
type GlogWriter struct{}

// Writer implements the io.Writer interface
func (writer GlogWriter) Write(data []byte) (n int, err error) {
	glog.Info(string(data))
	return len(data), nil
}

// InitLogs initializes logs the way we want for Yunus
func InitLogs() {
	log.SetOutput(GlogWriter{})
	log.SetFlags(0)
	// The default glog flush interval is 30's, which is quitely long
	go wait.Until(glog.Flush, *logFlushFreq, wait.NeverStop)
}

// FlushLogs flush logs immediately
func FlushLogs() {
	glog.Flush()
}

// NewLogger create a new log.Logger which sends log to glog.Info
func NewLogger(prefix string) *log.Logger {
	return log.New(GlogWriter{}, prefix, 0)
}
