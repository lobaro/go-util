package threadsafe

import (
	"bytes"
	"sync"
	"io"
)

type Buffer struct {
	bytes.Buffer
	sync.Mutex
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.Read(p)
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.Write(p)
}

func (b *Buffer) String() string {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.String()
}

func (b *Buffer) Bytes() []byte {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.Bytes()
}

func (b *Buffer) Cap() int {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.Cap()
}

func (b *Buffer) Grow(n int) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.Buffer.Grow(n)
}

func (b *Buffer) Len() int {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.Len()
}

func (b *Buffer) Next(n int) []byte {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.Next(n)
}

func (b *Buffer) ReadByte() (c byte, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.ReadByte()
}

func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.ReadBytes(delim)
}

func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.ReadFrom(r)
}

func (b *Buffer) ReadRune() (r rune, size int, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.ReadRune()
}

func (b *Buffer) ReadString(delim byte) (line string, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.ReadString(delim)
}

func (b *Buffer) Reset() {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.Buffer.Reset()
}

func (b *Buffer) Truncate(n int) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.Buffer.Truncate(n)
}

func (b *Buffer) UnreadByte() error {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.UnreadByte()
}

func (b *Buffer) UnreadRune() error {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.UnreadRune()
}

func (b *Buffer) WriteByte(c byte) error {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.WriteByte(c)
}

func (b *Buffer) WriteRune(r rune) (n int, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.WriteRune(r)
}

func (b *Buffer) WriteString(s string) (n int, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.WriteString(s)
}

func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.Buffer.WriteTo(w)
}
