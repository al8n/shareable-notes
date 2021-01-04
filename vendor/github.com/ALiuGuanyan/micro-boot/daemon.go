package boot

import (
	"context"
	"fmt"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	daemonPath = "/var/run"
	internalRunSuffix = "boot-in-daemon"
)

var (
	defaultConfigFileFlag = "config"
	defaultConfigFileName = "config"
	defaultConfigFileType = ".yaml"
	defaultConfigFileDirectory = ""
	defaultConfigFileFlagShortHand = "c"
	defaultConfigFileFlagUsage = "specify the config file"

	defaultLogFileFlag = "log"
	defaultLogFileDir = "/var/log"
	defaultLogFileFlagShortHand = "l"
	defaultLogFileFlagUsage = "specify the log file"
)

func SetDefaultConfigFileName(val string)  {
	defaultConfigFileName = val
}

func SetDefaultConfigFileType(val string)  {
	defaultConfigFileType = val
}

func SetDefaultConfigFileDirectory(val string)  {
	defaultConfigFileDirectory = val
}

func SetDefaultConfigFileFlagName(val string)  {
	defaultConfigFileFlag = val
}

// SetDefaultConfigFileFlagUsage will set the default config flag usage.
//
// eg. SetDefaultConfigFileFlagUsage("specify the config file")
//
// The output:
//
// -c, --config		specify the config file
func SetDefaultConfigFileFlagUsage(val string)  {
	defaultLogFileFlagUsage = val
}

func SetDefaultConfigFileFlagShortHand(val string)  {
	defaultConfigFileFlagShortHand = val
}

// SetDefaultLogFileDirectory will set the default directory to
// store the log file.
//
// eg. SetDefaultLogFileDirectory("/var/log")
func SetDefaultLogFileDirectory(val string)  {
	defaultLogFileDir = val
}

// SetDefaultLogFileFlagUsage will set the default log flag usage.
//
// eg. SetDefaultLogFileFlagUsage("specify the log file")
//
// The output:
//
// -l, --log		specify the log file
func SetDefaultLogFileFlagUsage(val string)  {
	defaultLogFileFlagUsage = val
}

func SetDefaultLogFileFlagName(val string)  {
	defaultLogFileFlag = val
}

func SetDefaultLogFileFlagShortHand(val string)  {
	defaultLogFileFlagShortHand = val
}

func newStartCmd(rootName string, fs *bootflag.FlagSet, config *Config) (cmd *Command) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil
	}

	var (
		shortUsage = fmt.Sprintf("%s start [flags]", rootName)
		shortHelp, longHelp string
		opts = []Option{WithEnvVarNoPrefix(), WithAllowMissingConfigFile(true), WithConfigFile(filepath.Join(pwd, root.configFile)), WithConfigFileParser(configParser)}
	)

	if config != nil {
		if config.Configurator != nil {
			root.start = config.Configurator
			root.start.BindFlags(fs)
		}

		if len(config.Options) > 0 {
			opts = config.Options
		}

		if strings.TrimSpace(config.ShortUsage) != "" {
			shortUsage = config.ShortUsage
		}

		if strings.TrimSpace(config.ShortHelp) != "" {
			shortHelp = config.ShortHelp
		}

		if strings.TrimSpace(config.LongHelp) != "" {
			longHelp = config.LongHelp
		}
	}

	return &Command{
		Name:       "start",
		ShortUsage: shortUsage,
		ShortHelp: shortHelp,
		LongHelp: longHelp,
		Options: opts,
		FlagSet:  fs,
		Exec: func(ctx context.Context, args []string) (err error) {
			if ls.daemon {
				var (
					abs string
					arguments = []string{getInternalCmdName(rootName), "-d"}
					begin = false
				)

				abs, err = filepath.Abs(os.Args[0])
				if err != nil{
					return err
				}

				for i := 0; i < len(os.Args); i++ {
					if begin {
						if os.Args[i] != "-d" && os.Args[i] != "--daemon" {
							arguments = append(arguments, os.Args[i])
						}
					} else {
						if os.Args[i] == "start" {
							begin = true
						}
					}
				}

				ccmd := exec.Command(abs, arguments...)
				ccmd.Env = os.Environ()
				ccmd.Stdout = os.Stdout
				ccmd.Stderr = os.Stderr
				ccmd.Dir = "/"

				if err = ccmd.Start(); err != nil {
					logrus.Error("services fail to start")
					logrus.Error(err)
					return err
				}

				err = ccmd.Wait()
				if err != nil {
					return err
				}

				return nil
			}


			pwd, err := os.Getwd()
			if err != nil {
				return err
			}

			fcfgFile := filepath.Join(pwd,root.configFile)

			logStartInfo(rootName, root.logFile, fcfgFile)

			if root.start != nil {
				err = root.start.Initialize(fcfgFile)
				if err != nil {
					return err
				}
			}

			err = ls.serve()
			if err != nil {
				return
			}

			return nil
		},
	}
}

func newRunCmd(name string, fs *bootflag.FlagSet, options ...Option) (cmd *Command) {
	return &Command{
		Name: getInternalCmdName(name),
		Hidden: true,
		FlagSet: fs,
		Options: options,
		Exec: func(ctx context.Context, args []string) (err error) {
			var (
				lf *os.File
				abs string
			)

			lf, err = openLogFile(root.logFile)
			if err != nil {
				return err
			}

			defer lf.Close()

			if ls.daemon {
				abs, err = filepath.Abs(os.Args[0])
				if err != nil{
					return err
				}

				arguments  := []string{getInternalCmdName(name)}
				arguments = append(arguments, os.Args[3:]...)

				ccmd := exec.Command(abs, arguments...)
				ccmd.Env = os.Environ()
				ccmd.Stdout = lf
				ccmd.Stderr = lf
				ccmd.Dir = "/"

				if err = ccmd.Start(); err != nil {
					logrus.SetOutput(lf)
					logrus.Error("services fail to start")
					logrus.Error(err)
					return err
				}
				return nil
			}


			fcfgPath := filepath.Join(abs,root.configFile)
			logStartInfo(name, root.logFile, fcfgPath)

			logrus.SetOutput(lf)
			if root.start != nil {
				err = root.start.Initialize(fcfgPath)
				if err != nil {
					return err
				}
			}

			err = ls.serve()
			if err != nil {
				return
			}

			return nil
		},
	}
}

func logStartInfo(name, logfilename, configfilename string)  {
	logrus.Info(fmt.Sprintf("starting %s service daemon...", name))

	logrus.Info("log: ", logfilename)
	logrus.Info("config: ", configfilename)
}

func getInternalCmdName(name string) string {
	return fmt.Sprintf("%s-%s", name, internalRunSuffix)
}

func openLogFile(filename string) (lf *os.File, err error) {
	dir, _ := filepath.Split(filename)
	if err = os.MkdirAll(dir, 0644); err != nil {
		return nil, err
	}

	lf, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return
}

