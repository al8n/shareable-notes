// This file is modified according to https://github.com/peterbourgon/ff/ffcli and https://github.com/ALiuGuanyan/ff/ffcli
package boot

import (
	"context"
	"errors"
	"flag"
	"fmt"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/spf13/pflag"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

// Command combines a main function with a flag.FlagSet, and zero or more
// sub-commands. A commandline program can be represented as a declarative tree
// of commands.
type Command struct {
	// Name of the command. Used for sub-command matching, and as a replacement
	// for Usage, if no Usage string is provided. Required for sub-commands,
	// optional for the root command.
	Name string

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

	// usageFunc generates a complete usage output, written to the io.Writer
	// returned by FlagSet.Output() when the -h flag is passed. The function is
	// invoked with its corresponding command, and its output should reflect the
	// command's short usage, short help, and long help strings, subcommands,
	// and available flags. Optional; if not provided, a suitable, compact
	// default is used.
	usageFunc func(c *Command) string

	// FlagSet associated with this command. Optional, but if none is provided,
	// an empty FlagSet will be defined and attached during the parse phase, so
	// that the -h flag works as expected.
	FlagSet *bootflag.FlagSet

	// Options provided to ff.Parse when parsing arguments for this command.
	// Optional.
	Options []Option

	// Subcommands accessible underneath (i.e. after) this command. Optional.
	Subcommands []*Command

	subcommands map[string]*Command

	// Hidden is used to make command invisible.
	Hidden bool

	// A successful Parse populates these unexported fields.
	selected *Command // the command itself (if terminal) or a subcommand
	args     []string // args that should be passed to Run, if any

	// Exec is invoked if this command has been determined to be the terminal
	// command selected by the arguments provided to Parse or ParseAndRun. The
	// args passed to Exec are the args left over after flags parsing. Optional.
	//
	// If Exec returns flag.ErrHelp, then Run (or ParseAndRun) will behave as if
	// -h were passed and emit the complete usage output.
	//
	// If Exec is nil, and this command is identified as the terminal command,
	// then Parse, Run, and ParseAndRun will all return noExecError. Callers may
	// check for this error and print e.g. help or usage text to the user, in
	// effect treating some commands as just collections of subcommands, rather
	// than being invocable themselves.
	Exec func(ctx context.Context, args []string) error
}

// Parse the commandline arguments for this command and all sub-commands
// recursively, defining flags along the way. If Parse returns without an error,
// the terminal command has been successfully identified, and may be invoked by
// calling Run.
//
// If the terminal command identified by Parse doesn't define an Exec function,
// then Parse will return noExecError.
func (c *Command) Parse(args []string) error {
	if c.selected != nil {
		return nil
	}

	if c.FlagSet == nil {
		c.FlagSet = bootflag.NewFlagSet(c.Name, bootflag.ExitOnError)
	}

	if c.usageFunc == nil {
		c.usageFunc = defaultusageFunc
	}

	c.FlagSet.GetFlagSet().Usage = func() {
		fmt.Fprintln(c.FlagSet.GetFlagSet().Output(), c.usageFunc(c))
	}

	numOfSubcommands := len(c.Subcommands)
	if numOfSubcommands > 0 {
		c.subcommands = make(map[string]*Command, numOfSubcommands)
		for _, subcommand := range c.Subcommands {
			c.subcommands[subcommand.Name] = subcommand
		}

		for index, arg := range args {
			sub, ok := c.subcommands[arg]
			if ok {
				if err := Parse(c.FlagSet.GetFlagSet(), args[:index], c.Options...); err != nil {
					return err
				}

				c.args = c.FlagSet.GetFlagSet().Args()

				err := sub.Parse(args[index + 1:])
				if err != nil {
					return err
				}
				c.selected = sub

				if sub.Exec == nil {
					return noExecError{Command: sub}
				}
				return nil
			}
		}
	}

	if err := Parse(c.FlagSet.GetFlagSet(), args, c.Options...); err != nil {
		return err
	}

	c.args = c.FlagSet.GetFlagSet().Args()

	c.selected = c

	if c.Exec == nil {
		return noExecError{Command: c}
	}

	return nil
}

// Run selects the terminal command in a command tree previously identified by a
// successful call to Parse, and calls that command's Exec function with the
// appropriate subset of commandline args.
//
// If the terminal command previously identified by Parse doesn't define an Exec
// function, then Run will return noExecError.
func (c *Command) Run(ctx context.Context) (err error) {
	var (
		unparsed = c.selected == nil
		terminal = c.selected == c && c.Exec != nil
		noop     = c.selected == c && c.Exec == nil
	)

	defer func() {
		if terminal && errors.Is(err, flag.ErrHelp) {
			c.FlagSet.GetFlagSet().Usage()
		}
	}()

	switch {
	case unparsed:
		return errUnparsed
	case terminal:
		return c.Exec(ctx, c.args)
	case noop:
		return noExecError{Command: c}
	default:
		return c.selected.Run(ctx)
	}
}

// ParseAndRun is a helper function that calls Parse and then Run in a single
// invocation. It's useful for simple command trees that don't need two-phase
// setup.
func (c *Command) ParseAndRun(ctx context.Context, args []string) error {
	if err := c.Parse(args); err != nil {
		return err
	}

	if err := c.Run(ctx); err != nil {
		return err
	}

	return nil
}

// errUnparsed is returned by Run if Parse hasn't been called first.
var errUnparsed = errors.New("command tree is unparsed, can't run")

// noExecError is returned if the terminal command selected during the parse
// phase doesn't define an Exec function.
type noExecError struct {
	Command *Command
}

// Error implements the error interface.
func (e noExecError) Error() string {
	return fmt.Sprintf("terminal command (%s) doesn't define an Exec function", e.Command.Name)
}


// DefaultusageFunc is the default usageFunc used for all commands
// if no custom usageFunc is provided.
func defaultusageFunc(c *Command) string {
	var b strings.Builder

	fmt.Fprintf(&b, "USAGE\n")
	if c.ShortUsage != "" {
		fmt.Fprintf(&b, "  %s\n", c.ShortUsage)
	} else {
		fmt.Fprintf(&b, "  %s\n", c.Name)
	}
	fmt.Fprintf(&b, "\n")

	if c.LongHelp != "" {
		fmt.Fprintf(&b, "%s\n\n", c.LongHelp)
	}

	hasVisibleSubcommands := false
	for _, subcommand := range c.Subcommands {
		if !subcommand.Hidden {
			hasVisibleSubcommands = true
			break
		}
	}

	if hasVisibleSubcommands {
		fmt.Fprintf(&b, "SUBCOMMANDS\n")
		tw := tabwriter.NewWriter(&b, 0, 4, 4, ' ', 0)
		for _, subcommand := range c.Subcommands {
			if !subcommand.Hidden {
				fmt.Fprintf(tw, "  %s\t%s\n", subcommand.Name, subcommand.ShortHelp)
			}
		}
		tw.Flush()
		fmt.Fprintf(&b, "\n")
	}

	if countFlags(c.FlagSet.GetFlagSet()) > 0 {
		fmt.Fprintf(&b, "FLAGS\n")
		tw := tabwriter.NewWriter(&b, 0, 2, 2, ' ', 0)
		c.FlagSet.GetFlagSet().VisitAll(func(f *pflag.Flag) {
			if strings.Contains(f.Name, "zzzzzz") {
				_, _ = fmt.Fprintf(tw, "\n\n")
			} else {
				if f.Shorthand != "" {
					_, _ = fmt.Fprintf(tw, "\t-%s, --%s\t%s\n", f.Shorthand, f.Name, f.Usage)
				} else {
					_, _ = fmt.Fprintf(tw, "\t    --%s\t%s\n", f.Name, f.Usage)
				}

			}
		})
		tw.Flush()
		fmt.Fprintf(&b, "\n")
	}

	return strings.TrimSpace(b.String())
}

func countFlags(fs *pflag.FlagSet) (n int) {
	fs.VisitAll(func(*pflag.Flag) { n++ })
	return n
}

// Parse the flags in the flag set from the provided (presumably commandline)
// args. Additional options may be provided to parse from a config file and/or
// environment variables in that priority order.
func Parse(fs *pflag.FlagSet, args []string, options ...Option) error {
	var c pContext
	for _, option := range options {
		option(&c)
	}

	// First priority: commandline flags (explicit user preference).
	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("error parsing commandline args: %w", err)
	}

	provided := map[string]bool{}
	fs.Visit(func(f *pflag.Flag) {
		provided[f.Name] = true
	})

	// Second priority: environment variables (session).
	if parseEnv := c.envVarPrefix != "" || c.envVarNoPrefix; parseEnv {
		var visitErr error
		fs.VisitAll(func(f *pflag.Flag) {
			if visitErr != nil {
				return
			}

			if provided[f.Name] {
				return
			}

			var key string
			key = strings.ToUpper(f.Name)
			key = envVarReplacer.Replace(key)
			key = maybePrefix(key, c.envVarNoPrefix, c.envVarPrefix)
			value := os.Getenv(key)
			if value == "" {
				return
			}

			for _, v := range maybeSplit(value, c.envVarSplit) {
				if err := fs.Set(f.Name, v); err != nil {
					visitErr = fmt.Errorf("error setting flag %q from env var %q: %w", f.Name, key, err)
					return
				}
			}
		})
		if visitErr != nil {
			return fmt.Errorf("error parsing env vars: %w", visitErr)
		}
	}

	fs.Visit(func(f *pflag.Flag) {
		provided[f.Name] = true
	})

	var configFile string
	if c.configFileVia != nil {
		configFile = *c.configFileVia
	}

	// Third priority: config file (host).
	if configFile == "" && c.configFileFlagName != "" {
		if f := fs.Lookup(c.configFileFlagName); f != nil {
			configFile = f.Value.String()
		}
	}

	if parseConfig := configFile != "" && c.configFileParser != nil; parseConfig {
		f, err := os.Open(configFile)
		switch {
		case err == nil:
			defer f.Close()
			if err := c.configFileParser(f, func(name, value string) error {
				if provided[name] {
					return nil
				}

				defined := fs.Lookup(name) != nil
				switch {
				case !defined && c.ignoreUndefined:
					return nil
				case !defined && !c.ignoreUndefined:
					fmt.Println(configFile, name)
					return fmt.Errorf("config file flag %q not defined in flag set", name)
				}

				if err := fs.Set(name, value); err != nil {
					return fmt.Errorf("error setting flag %q from config file: %w", name, err)
				}

				return nil
			}); err != nil {
				return err
			}

		case os.IsNotExist(err) && c.allowMissingConfigFile:
			// no problem

		default:
			return err
		}
	}

	fs.Visit(func(f *pflag.Flag) {
		provided[f.Name] = true
	})

	return nil
}

// pContext contains private fields used during parsing.
type pContext struct {
	configFileVia          *string
	configFileFlagName     string
	configFileParser       ConfigFileParser
	allowMissingConfigFile bool
	envVarPrefix           string
	envVarNoPrefix         bool
	envVarSplit            string
	ignoreUndefined        bool
}

// Option controls some aspect of Parse behavior.
type Option func(*pContext)

// WithConfigFile tells Parse to read the provided filename as a config file.
// Requires WithConfigFileParser, and overrides WithConfigFileFlag.
// Because config files should generally be user-specifiable, this option
// should be rarely used. Prefer WithConfigFileFlag.
func WithConfigFile(filename string) Option {
	return WithConfigFileVia(&filename)
}

// WithConfigFileVia tells Parse to read the provided filename as a config file.
// Requires WithConfigFileParser, and overrides WithConfigFileFlag.
// This is useful for sharing a single root level flag for config files among
// multiple ffcli subcommands.
func WithConfigFileVia(filename *string) Option {
	return func(c *pContext) {
		c.configFileVia = filename
	}
}

// WithConfigFileFlag tells Parse to treat the flag with the given name as a
// config file. Requires WithConfigFileParser, and is overridden by
// WithConfigFile.
//
// To specify a default config file, provide it as the default value of the
// corresponding flag -- and consider also using the WithAllowMissingConfigFile
// option.
func WithConfigFileFlag(flagname string) Option {
	return func(c *pContext) {
		c.configFileFlagName = flagname
	}
}

// WithConfigFileParser tells Parse how to interpret the config file provided
// via WithConfigFile or WithConfigFileFlag.
func WithConfigFileParser(p ConfigFileParser) Option {
	return func(c *pContext) {
		c.configFileParser = p
	}
}

// WithAllowMissingConfigFile tells Parse to permit the case where a config file
// is specified but doesn't exist. By default, missing config files result in an
// error.
func WithAllowMissingConfigFile(allow bool) Option {
	return func(c *pContext) {
		c.allowMissingConfigFile = allow
	}
}

// WithEnvVarPrefix tells Parse to try to set flags from environment variables
// with the given prefix. Flag names are matched to environment variables with
// the given prefix, followed by an underscore, followed by the capitalized flag
// names, with separator characters like periods or hyphens replaced with
// underscores. By default, flags are not set from environment variables at all.
func WithEnvVarPrefix(prefix string) Option {
	return func(c *pContext) {
		c.envVarPrefix = prefix
	}
}

// WithEnvVarNoPrefix tells Parse to try to set flags from environment variables
// without any specific prefix. Flag names are matched to environment variables
// by capitalizing the flag name, and replacing separator characters like
// periods or hyphens with underscores. By default, flags are not set from
// environment variables at all.
func WithEnvVarNoPrefix() Option {
	return func(c *pContext) {
		c.envVarNoPrefix = true
	}
}

// WithEnvVarSplit tells Parse to split environment variables on the given
// delimiter, and to make a call to Set on the corresponding flag with each
// split token.
func WithEnvVarSplit(delimiter string) Option {
	return func(c *pContext) {
		c.envVarSplit = delimiter
	}
}

// WithIgnoreUndefined tells Parse to ignore undefined flags that it encounters
// in config files. By default, if Parse encounters an undefined flag in a
// config file, it will return an error. Note that this setting does not apply
// to undefined flags passed as arguments.
func WithIgnoreUndefined(ignore bool) Option {
	return func(c *pContext) {
		c.ignoreUndefined = ignore
	}
}

// ConfigFileParser interprets the config file represented by the reader
// and calls the set function for each parsed flag pair.
type ConfigFileParser func(r io.Reader, set func(name, value string) error) error



var envVarReplacer = strings.NewReplacer(
	"-", "_",
	".", "_",
	"/", "_",
)

func maybePrefix(key string, noPrefix bool, prefix string) string {
	if noPrefix {
		return key
	}
	return strings.ToUpper(prefix) + "_" + key
}

func maybeSplit(value, split string) []string {
	if split == "" {
		return []string{value}
	}
	return strings.Split(value, split)
}
