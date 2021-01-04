package boot

import (
	"fmt"
	localrpc "github.com/ALiuGuanyan/micro-boot/internal/rpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type localServer struct {
	name string
	external Booter
	server         *grpc.Server
	serverOpts     []grpc.ServerOption
	dialOpts       []grpc.DialOption
	ln             net.Listener
	mu             sync.Mutex
	exiting        bool
	daemon		   bool
	force      		bool
	gracefulSignal chan os.Signal
	forceSignal    chan os.Signal
}

func (l *localServer) Stop(stream localrpc.Boot_StopServer) error {
	for {
		req, _ := stream.Recv()
		if req.Force {
			_ = stream.Send(&localrpc.StopResponse{Errno: -1, Msg: forceStopMessage(l.name)})
			l.mu.Lock()
			l.forceStop()
			l.mu.Unlock()
			return nil
		} else {
			_ = stream.Send(&localrpc.StopResponse{Errno: 0, Msg: stopMessage(l.name)})
			l.mu.Lock()
			if l.exiting {
				l.forceSignal <- syscall.SIGTERM
			} else {
				l.exiting = true
				l.gracefulSignal <- syscall.SIGTERM
			}
			l.mu.Unlock()
		}
	}
}

func (l *localServer) gracefulStop() {
	logrus.Info(stopMessage(l.name))
	l.external.Close()
	logrus.Info(gracefulStopMessage(l.name))
	l.ln.Close()
	return
}

func (l *localServer) forceStop() {
	logrus.Error(forceStopMessage(l.name))
	l.ln.Close()
}

func (l *localServer) listenForcefulStop() {
	select {
	case <-l.forceSignal:
		go l.forceStop()
		break
	}
}

func (l *localServer) listenGracefulStop()  {
	select {
	case <-l.gracefulSignal:
		go l.gracefulStop()
		break
	}
}

func (l *localServer) listenSystemCall() {
	c := make(chan os.Signal, 2)

	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			l.mu.Lock()
			if l.exiting {
				l.forceSignal <- s
			} else {
				l.exiting = true
				l.gracefulSignal <- s
			}
			l.mu.Unlock()
		}
	}
}

func (l *localServer) serve() (err error) {

	go l.listenGracefulStop()
	go l.listenForcefulStop()
	go l.listenSystemCall()

	l.ln, err = net.Listen("unix", socketPath(l.name))
	if err != nil {
		return
	}

	l.server = grpc.NewServer(l.serverOpts...)
	localrpc.RegisterBootServer(l.server, l)
	err = l.external.Serve()
	if err != nil {
		return
	}

	_ = l.server.Serve(l.ln)

	return
}


func logPath(dir, name string) string {
	return fmt.Sprintf("%s/%s.log", dir, name)
}

func socketPath(name string) string {
	return fmt.Sprintf("%s/%s.sock", daemonPath, name)
}

func stopMessage(name string) string {
	return fmt.Sprintf("Stopping %s services...", name)
}

func forceStopMessage(name string) string {
	return fmt.Sprintf("%s services are forced to stop!", name)
}

func gracefulStopMessage(name string) string {
	return fmt.Sprintf("%s services are stopped gracefully", name)
}
