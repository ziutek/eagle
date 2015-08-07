package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/ziutek/sched"
	"github.com/ziutek/thread"
)

func dieErr(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err.Error())
	os.Exit(1)
}

func checkErr(err error) {
	if err == nil {
		return
	}
	dieErr(err)
}

var startTime time.Time

func init() {
	startTime = time.Now()
}

func dtime() time.Duration {
	return time.Now().Sub(startTime)
}

func nanosleep(d time.Duration) {
	if d <= 0 {
		return
	}
	to := syscall.NsecToTimespec(int64(d))
	for {
		err := syscall.Nanosleep(&to, &to)
		if err == nil {
			return
		}
		if err != syscall.EINTR {
			dieErr(err)
		}
	}
}

func setRealtimeSched(what string, prio100 int) {
	runtime.LockOSThread()
	if os.Geteuid() != 0 {
		fmt.Fprintln(os.Stderr, what+": can't set realtime scheduling: no root priv.")
		return
	}
	t := thread.Current()
	fmt.Fprintf(
		os.Stderr,
		"%s: realtime sheduling prio %d%% for thread %d",
		what, prio100, t,
	)
	p := sched.Param{
		Priority: sched.FIFO.MaxPriority() * prio100 / 100,
	}
	checkErr(t.SetSchedPolicy(sched.FIFO, &p))
}
