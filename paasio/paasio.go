package paasio

import (
	"io"
	"sync"
)

type counter struct {
	bytes int64
	ops   int
	mutex *sync.Mutex
}

type readCounter struct {
	r io.Reader
	counter
}

type writeCounter struct {
	w io.Writer
	counter
}

type rwCounter struct {
	WriteCounter
	ReadCounter
}

func newCounter() counter {
	return counter{mutex: new(sync.Mutex)}
}

func (c *counter) addBytes(n int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.bytes += int64(n)
	c.ops++
}

func (c *counter) count() (int64, int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.bytes, c.ops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	m, err := wc.w.Write(p)
	wc.addBytes(m)
	return m, err
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	return wc.count()
}

func (rc *readCounter) Read(p []byte) (int, error) {
	m, err := rc.r.Read(p)
	rc.addBytes(m)
	return m, err
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	return rc.count()
}

// NewReadWriteCounter returns an implementation of ReadWriteCounter.
// Calls to rw.Write() and rw.Read() are not guaranteed to be synchronized.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{
		NewWriteCounter(rw),
		NewReadCounter(rw),
	}
}

// NewReadCounter returns an implementation of ReadCounter.  Calls to
// r.Read() are not guaranteed to be synchronized.
func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		r:       reader,
		counter: newCounter(),
	}
}

// NewWriteCounter returns an implementation of WriteCounter.  Calls to
// w.Write() are not guaranteed to be synchronized.
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		w:       writer,
		counter: newCounter(),
	}
}
