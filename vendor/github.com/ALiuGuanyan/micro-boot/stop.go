package boot

import (
	"context"
	"fmt"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	localrpc "github.com/ALiuGuanyan/micro-boot/internal/rpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func newStopCmd(name string, fs *bootflag.FlagSet, config *Config, options ...Option) *Command {
	var (
		shortUsage = fmt.Sprintf("%s stop [flags]", name)
		shortHelp, longHelp string
	)

	if config != nil {
		if config.Configurator != nil {
			root.stop = config.Configurator
			root.stop.BindFlags(fs)
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
		Name: "stop",
		ShortUsage: shortUsage,
		ShortHelp: shortHelp,
		LongHelp: longHelp,
		FlagSet: fs,
		Options: options,
		Exec: func(ctx context.Context, args []string) (err error) {
			var (
				conn *grpc.ClientConn
				client localrpc.BootClient
				stream localrpc.Boot_StopClient
				sockfile = socketPath(ls.name)
				resp *localrpc.StopResponse
				dialOpts = ls.dialOpts
			)

			dialOpts = append(dialOpts, grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
				return net.Dial("unix", sockfile)
			}))

			conn, err = grpc.Dial(sockfile, dialOpts...)

			if err != nil {
				return err
			}

			defer conn.Close()

			client = localrpc.NewBootClient(conn)

			stream, err = client.Stop(context.Background())
			if err != nil {
				return err
			}

			if !ls.force {
				c := make(chan os.Signal)
				signal.Notify(c,  syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
				go func() {
					for s := range c {
						switch s {
						case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
							stream.Send(&localrpc.StopRequest{Force: true})
							logrus.Error(forceStopMessage(ls.name))
							conn.Close()
							os.Exit(1)
						}
					}
				}()
			}

			err = stream.Send(&localrpc.StopRequest{Force: false})
			if err != nil{
				return err
			}

			for {
				resp, _ = stream.Recv()
				if resp == nil {
					return nil
				}

				if resp.Errno == -1 {
					logrus.Error(resp.Msg)
				} else if resp.Errno == 0 {
					logrus.Warn(resp.Msg)
				} else {
					logrus.Info(resp.Msg)
					break
				}
			}
			return nil
		},
	}
}