// A concurrent prime sieve, copied from tinyurl.com/gosieve

package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

type oper string

const (
	reading     oper = "R"
	doneReading oper = "DR"
	writing     oper = "W"
)

type seqLog struct {
	time      int64
	routine   int
	channel   string
	worker    string
	operation oper
	value     int
}

func (s seqLog) String() string {
	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 10, 16, 5, '\t', 0)
	fmt.Fprintf(w, "%s-%d\t%s\t%s\t%d", s.worker, s.routine, s.channel, s.operation, s.value)
	w.Flush()
	return ""
	// return fmt.Sprintf("%d:%s:%s:%s:%d", s.routine, s.worker, s.channel, s.operation, s.value)
}

type logs []seqLog

var lg = make(logs, 0)

func (s logs) Less(i, j int) bool { return (s[i].time < s[j].time) }
func (s logs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s logs) Len() int           { return len(s) }

var channelNames = make(map[string]string)
var mapLock = &sync.Mutex{}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		chName := readFromMap2(ch)
		logWriting(chName, i, "Gen")
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		inName := readFromMap1(in)
		logReading(inName, "Fil")
		i := <-in // Receive value from 'in'.
		logDoneReading(inName, i, "Fil")
		if i%prime != 0 {
			outName := readFromMap2(out)
			logWriting(outName, i, "Fil")
			out <- i // Send 'i' to 'out'.
		}
	}
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

// The prime sieve: Daisy-chain Filter processes.
func primeSeiveAutopsied() {
	ch := make(chan int) // Create a new channel.
	writeToMap(ch, "ch")
	go generate(ch) // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		chName := readFromMap(ch)
		logReading(chName, "Main")
		prime := <-ch
		logDoneReading(chName, prime, "Main")
		print(prime, "\n")
		ch1 := make(chan int)
		writeToMap(ch1, fmt.Sprintf("ch%d", i))
		go filter(ch, ch1, prime)
		ch = ch1
	}
	dumbLog()
}

func dumbLog() {
	fmt.Println("")
	fmt.Println("============Log============")
	fmt.Println("")
	sort.Sort(lg)
	for _, l := range lg {
		fmt.Println(l)
	}
}

func logReading(c string, funcName string) {
	// fmt.Printf("%s: Reading from (%s), from [%d]\n", funcName, c, goid())
	log := seqLog{time.Now().UnixNano(), goid(), c, funcName, reading, -1}
	lg = append(lg, log)
}

func logDoneReading(c string, value int, funcName string) {
	// fmt.Printf("%s: Just read %d from (%s), from [%d]\n", funcName, value, c, goid())
	log := seqLog{time.Now().UnixNano(), goid(), c, funcName, doneReading, value}
	lg = append(lg, log)
}

func logWriting(c string, value int, funcName string) {
	// fmt.Printf("%s: Writing %d to (%s), from [%d]\n", funcName, value, c, goid())
	log := seqLog{time.Now().UnixNano(), goid(), c, funcName, writing, value}
	lg = append(lg, log)
}

func writeToMap(c chan int, value string) {
	mapLock.Lock()
	channelNames[fmt.Sprintf("%v", c)] = value
	mapLock.Unlock()
}

func readFromMap(c chan int) string {
	mapLock.Lock()
	v := channelNames[fmt.Sprintf("%v", c)]
	mapLock.Unlock()
	return v
}

func readFromMap1(c <-chan int) string {
	mapLock.Lock()
	v := channelNames[fmt.Sprintf("%v", c)]
	mapLock.Unlock()
	return v
}

func readFromMap2(c chan<- int) string {
	mapLock.Lock()
	v := channelNames[fmt.Sprintf("%v", c)]
	mapLock.Unlock()
	return v
}
