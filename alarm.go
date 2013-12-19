// Alarm provides a method to call a function within a go routine after a duration
// of time. Further information can be found at https://github.com/zdannar/alarm.
package alarm

import (
    "time"
    "sync"
)

// Initialize OnAlarm with a time duration, and a function that doesn't take 
// arguments or returns anything.  State changes can be managed with closures
// and structures.  OnAlarm returns a function to cancel the alarm.  The 
// returned cancel() should be always be called.
func OnAlarm(dur time.Duration, f func()) func() {

    var quit chan bool = make(chan bool)

    // variable to make sure that deadlock doesn't occur
    var done bool = false
    mu := sync.Mutex{}

    go func() () {
        ac := time.After(dur)
SELECT:
        for {
            select {
            case <-quit:
                mu.Lock(); done = true; mu.Unlock()
                return 
            case <- ac:
                break SELECT
            }
        }
        mu.Lock(); done = true; mu.Unlock()
        f()
    }()

    // Return the cancel function
    return func() {
        mu.Lock()
        if !done {
            quit <- true
        }
        mu.Unlock()
    }
}
