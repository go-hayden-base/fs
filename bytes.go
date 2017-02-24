package fs

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
)

type EnumerableBytes struct {
	bytes      []byte
	filterReg  *regexp.Regexp
	filterFunc func(line string, stop *bool) bool
}

func (s *EnumerableBytes) Bytes() []byte {
	return s.bytes
}

func (s *EnumerableBytes) Filter(reg string) *EnumerableBytes {
	s.filterReg = regexp.MustCompile(reg)
	return s
}

func (s *EnumerableBytes) FilterFunc(f func(line string, stop *bool) bool) *EnumerableBytes {
	s.filterFunc = f
	return s
}

func (s *EnumerableBytes) Enumerate(f func(line string, err error)) {
	if f == nil {
		return
	}
	reader := bytes.NewReader(s.bytes)
	bufReader := bufio.NewReader(reader)
	for {
		b, _, e := bufReader.ReadLine()
		if e != nil {
			if e == io.EOF {
				break
			} else {
				f("", e)
				return
			}
		}
		line := string(b)
		if s.filterReg != nil && !s.filterReg.MatchString(line) {
			continue
		}
		stop := false
		if s.filterFunc != nil && !s.filterFunc(line, &stop) {
			if stop {
				break
			}
			continue
		}
		f(string(b), nil)
		if stop {
			break
		}
	}
}

func NewEnumerableBytes(b []byte) *EnumerableBytes {
	eb := new(EnumerableBytes)
	eb.bytes = b
	return eb
}
