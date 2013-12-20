alarm
=======

###SUMMARY

Alarm is a simple package to call a function via a go routine after a duration
has elapsed.  An application is to use it to timeout and kill a subprocess while 
your main groutine does other things.

###Example

```go
package main

import (
    "time"
    log "github.com/zdannar/flogger"
    "github.com/zdannar/alarm"
)

func main() {

    log.Info("Starting")
    doStuff := func() {
        log.Infof("(Alarm routine) Doing some %s things", "important")
        log.Info("...")
    }

    // NOTE: You should always have and call the cancel function
    cancel := alarm.OnAlarm(time.Second * time.Duration(3), doStuff)

    log.Info("(Master routine) : Doing some things in the master routine")
    time.Sleep(time.Second * time.Duration(5))
    cancel()
    log.Info("Quiting")
}
```

The output of the previous example

```

2013/12/19 15:12:49 alarmt.go:53: : INFO : Starting
2013/12/19 15:12:49 alarmt.go:61: : INFO : (Master routine) : Doing some things in the master routine
2013/12/19 15:12:52 alarmt.go:55: : INFO : (Alarm routine) Doing some important things
2013/12/19 15:12:52 alarmt.go:56: : INFO : ...
2013/12/19 15:12:54 alarmt.go:64: : INFO : Quiting

```

###INSTALL

go get github.com/zdannar/alarm
