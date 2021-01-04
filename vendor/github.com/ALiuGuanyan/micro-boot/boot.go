package boot

import (
	"context"
	"github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"sync"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

var (
	ls *localServer
	onceExecute sync.Once
)

type Config struct {
	// ShortUsage string for this command. Consumed by the DefaultusageFunc and
	// printed at the top of the help output. Recommended but not required.
	// Should be one line of the form
	//
	//     cmd [flags] subcmd [flags] <required> [<optional> ...]
	//
	// If it's not provided, the DefaultusageFunc will use Name instead.
	// Optional, but recommended.
	ShortUsage string

	// ShortHelp is printed next to the command name when it appears as a
	// sub-command, in the help output of its parent command. Optional, but
	// recommended.
	ShortHelp string

	// LongHelp is consumed by the DefaultusageFunc and printed in the help
	// output, after ShortUsage and before flags. Typically a paragraph or more
	// of prose-like text, providing more explicit context and guidance than
	// what is implied by flags and arguments. Optional.
	LongHelp string

	Configurator Configurator

	Options  []Option
}

type Configurator interface {
	Initialize(name string) (err error)
	flag.Flags
}

type Boot interface {
	Execute() (err error)
}

type boot struct {
	cmd *Command
	logFile, configFile string
	start Configurator
	stop Configurator
}

func (b boot) Execute() (err error) {
	onceExecute.Do(func() {
		err = b.cmd.ParseAndRun(context.Background(), os.Args[1:])
	})
	return
}

type Root struct {
	Exec 				func(ctx context.Context, args []string) error
	ShortUsage          string
	ShortHelp			string
	LongHelp 			string
	ErrorHandling		flag.ErrorHandling
	Start 			*Config
	Stop   *Config
	Options []Option
	CustomCommands []*Command
	LocalRPCServerOptions []grpc.ServerOption
	LocalRPCDialOptions []grpc.DialOption
}

type Booter interface {
	Serve() (err error)
	Close() (err error)
}

