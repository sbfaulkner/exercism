package paasio

import (
	"io"
	"sync"
)

const testVersion = 3

type readCounter struct {
	r    io.Reader
	n    int64
	nops int
	mux  sync.RWMutex
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	rc.mux.RLock()
	defer rc.mux.RUnlock()
	n, nops = rc.n, rc.nops
	return
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	n, err = rc.r.Read(p)
	rc.mux.Lock()
	defer rc.mux.Unlock()
	rc.nops++
	rc.n += int64(n)
	return
}

// NewReadCounter wraps the provided io.Reader to count read operations and bytes read
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{r: r}
}

type writeCounter struct {
	w    io.Writer
	n    int64
	nops int
	mux  sync.RWMutex
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	wc.mux.RLock()
	defer wc.mux.RUnlock()
	n, nops = wc.n, wc.nops
	return
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.w.Write(p)
	wc.mux.Lock()
	defer wc.mux.Unlock()
	wc.nops++
	wc.n += int64(n)
	return
}

// NewWriteCounter wraps the provided io.Writer to count write operations and bytes written
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{w: w}
}

type readWriteCounter struct {
	r readCounter
	w writeCounter
}

func (rwc *readWriteCounter) Read(p []byte) (n int, err error) {
	return rwc.r.Read(p)
}

func (rwc *readWriteCounter) ReadCount() (n int64, nops int) {
	return rwc.r.ReadCount()
}

func (rwc *readWriteCounter) Write(p []byte) (n int, err error) {
	return rwc.w.Write(p)
}

func (rwc *readWriteCounter) WriteCount() (n int64, nops int) {
	return rwc.w.WriteCount()
}

// NewReadWriteCounter wraps the provided io.Reader/io.Writer to count read/write operations and bytes read/written
func NewReadWriteCounter(rw interface{}) ReadWriteCounter {
	return &readWriteCounter{readCounter{r: rw.(io.Reader)}, writeCounter{w: rw.(io.Writer)}}
}
