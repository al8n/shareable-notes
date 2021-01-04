package boot

import (
	"context"
	"fmt"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"google.golang.org/grpc"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	onceBoot sync.Once
	root = boot{}
)

var reservedCmds = map[string]bool{
	"start": true,
	"stop": true,
	"build": true,
}

func New(name string, server Booter, config Root) (bt Boot, err error) {
	onceBoot.Do(func() {
		err = constructRootCmd(name, server, config)
		if err != nil {
			return
		}
		constructDaemonCmd(name, config)
		constructStopCmd(name, config)
	})
	
	if err != nil {
		return nil, err
	}
	
	return root, nil
}

func constructStopCmd(name string, config Root)  {
	var stopCmd *Command
	{
		stopFlagSet := bootflag.NewFlagSet("stop", config.ErrorHandling)

		stopFlagSet.BoolVarP(&ls.force, "force", "f", false, fmt.Sprintf("force to stop %s daemon", name))

		stopCmd = newStopCmd(name, stopFlagSet,  config.Stop, config.Options...)

		root.cmd.Subcommands = append(root.cmd.Subcommands, stopCmd)
	}
}

func constructDaemonCmd(name string, config Root)  {
	var (
		startCmd, runCmd *Command
		logFilePath, configFilePath string
	)
	{
		logFilePath = logPath(defaultLogFileDir, name)

		configFilePath = filepath.Join(
			defaultConfigFileDirectory,
			defaultConfigFileName + defaultConfigFileType)

		daemonFlagSet := bootflag.NewFlagSet("daemon", config.ErrorHandling)

		daemonFlagSet.StringVarP(
			&root.logFile,
			defaultLogFileFlag,
			defaultLogFileFlagShortHand,
			logFilePath,
			fmt.Sprintf("%s (default is %s)", defaultLogFileFlagUsage, logFilePath))

		daemonFlagSet.StringVarP(
			&root.configFile,
			defaultConfigFileFlag,
			defaultConfigFileFlagShortHand,
			configFilePath,
			fmt.Sprintf("%s (default is %s)", defaultConfigFileFlagUsage, configFilePath))

		daemonFlagSet.BoolVarP(
			&ls.daemon,
			"daemon",
			"d",
			false,
			"run service in daemon mode")

		startCmd = newStartCmd(name, daemonFlagSet, config.Start)

		runCmd = newRunCmd(name, daemonFlagSet, config.Options...)

		root.cmd.Subcommands = append(root.cmd.Subcommands, startCmd, runCmd)
	}
}

func constructRootCmd(name string, server Booter, config Root) (err error) {
	var (
		dialOpts = []grpc.DialOption{grpc.WithInsecure()}

		subcommands []*Command

		execFn = defaultRootExec

		fs = bootflag.NewFlagSet(name, config.ErrorHandling)

		shortUsage = fmt.Sprintf("%s [flags] <subcommand>", name)

		shortHelp, longHelp string

		options = []Option{WithEnvVarNoPrefix(), WithAllowMissingConfigFile(true), WithConfigFile(root.configFile), WithConfigFileParser(configParser)}
	)
	{
		if strings.TrimSpace(config.LongHelp) != "" {
			longHelp = config.LongHelp
		}

		if strings.TrimSpace(config.ShortHelp) != "" {
			shortHelp = config.ShortHelp
		}

		if strings.TrimSpace(config.ShortUsage) != "" {
			shortUsage = config.ShortUsage
		}

		if len(config.Options) != 0 {
			options = config.Options
		}

		if config.Exec != nil {
			execFn = config.Exec
		}

		if len(config.LocalRPCDialOptions) > 0 {
			dialOpts = config.LocalRPCDialOptions
		}

		if len(config.CustomCommands) > 0 {
			for _, command := range config.CustomCommands {
				if _, ok := reservedCmds[strings.TrimSpace(command.Name)]; ok {
					err = fmt.Errorf("%s is a reserved command name by micro boot", command.Name)
					return err
				}
				subcommands = append(subcommands, command)
			}
		}
	}

	ls = &localServer{
		name: name,
		external: server,
		serverOpts: config.LocalRPCServerOptions,
		dialOpts: dialOpts,
		gracefulSignal: make(chan os.Signal, 1),
		forceSignal: make(chan os.Signal, 1),
		daemon: false,
		exiting: false,
	}

	root.cmd = &Command{
		Name: name,
		ShortUsage:  shortUsage,
		ShortHelp: shortHelp,
		LongHelp: longHelp,
		Options: options,
		FlagSet:     fs,
		Subcommands: subcommands,
		Exec: execFn,
	}
	return err
}

func defaultRootExec(ctx context.Context, args []string) (err error) {
	return nil
}