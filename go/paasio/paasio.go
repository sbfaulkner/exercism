package paasio

import (
	"io"
	"sync"
)

const testVersion = 3

type ioCounter struct {
	n    int64
	nops int
	mux  sync.RWMutex
}

func (ioc *ioCounter) count() (int64, int) {
	ioc.mux.RLock()
	defer ioc.mux.RUnlock()
	return ioc.n, ioc.nops
}

func (ioc *ioCounter) add(n int) {
	ioc.mux.Lock()
	defer ioc.mux.Unlock()
	ioc.nops++
	ioc.n += int64(n)
}

type readCounter struct {
	r io.Reader
	c ioCounter
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.c.count()
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	n, err = rc.r.Read(p)
	rc.c.add(n)
	return
}

// NewReadCounter wraps the provided io.Reader to count read operations and bytes read
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{r: r}
}

type writeCounter struct {
	w io.Writer
	c ioCounter
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	return wc.c.count()
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.w.Write(p)
	wc.c.add(n)
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
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{readCounter{r: rw}, writeCounter{w: rw}}
}
