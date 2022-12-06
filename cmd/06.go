package cmd

import (
	"bufio"
	"io"
)

func day06(input *bufio.Reader) error {
	var buf [14]byte
	if _, err := io.ReadFull(input, buf[:4]); err != nil {
		return err
	}

	havePacketMarker := func() bool {
		for i, a := range buf[:3] {
			for _, b := range buf[i+1 : 4] {
				if a == b {
					return false
				}
			}
		}
		return true
	}

	haveMessageMarker := func() bool {
		for i, a := range buf[:13] {
			for _, b := range buf[i+1:] {
				if a == b {
					return false
				}
			}
		}
		return true
	}

	i := 4
	for ; !havePacketMarker(); i++ {
		var err error
		if buf[i&3], err = input.ReadByte(); err != nil {
			return err
		}
	}

	partOne(i)

	copy(buf[10+(i&3):], buf[i&3:4])
	if _, err := io.ReadFull(input, buf[i&3:10+i&3]); err != nil {
		return err
	}

	j := i &^ 3
	i &= 3
	i += 10
	for ; !haveMessageMarker(); i++ {
		var err error
		if buf[i%14], err = input.ReadByte(); err != nil {
			return err
		}
	}

	partTwo(i + j)

	return nil
}
