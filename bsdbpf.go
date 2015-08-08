/*
 from 
  pcap.go
 and
  https://github.com/david415/HoneyBadger/blob/master/drivers/bpf.go
*/

package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/bsdbpf"
)

type BPFSniffer struct {
	handle *bsdbpf.BPFSniffer
}

func (s *BPFSniffer) Open(config *Config) error {
        // XXX TODO pass more options...
        handle, err := bsdbpf.NewBPFSniffer(*iface, nil)
        if err != nil {
		return fmt.Errorf("Error opening bsdbpf: %s", err)
        }
        s.handle = handle
        return nil
}

func (s *BPFSniffer) Close() {
	s.handle.Close()
}

func (s *BPFSniffer) ReadPacket() (data []byte, ci gopacket.CaptureInfo, err error) {
	return s.handle.ReadPacketData()
}
