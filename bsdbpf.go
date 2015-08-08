/*
 from 
  pcap.go
 and
  https://github.com/david415/HoneyBadger/blob/master/drivers/bpf.go
*/

package main

import (
	"fmt"
	"time"
	"github.com/google/gopacket"
	"github.com/google/gopacket/bsdbpf"
)

type BPFSniffer struct {
	handle *bsdbpf.BPFSniffer
}

// types.xxx
//  from https://github.com/hb9cwp/HoneyBadger/blob/master/types/packet_source.go
type SnifferDriverOptions struct {
	DAQ          string
	Filename     string
	Device       string
	Snaplen      int32
	WireDuration time.Duration
	Filter       string
}

// PacketDataSource is an interface for some source of packet data.
type PacketDataSourceCloser interface {
	// ReadPacketData returns the next packet available from this data source.
	// It returns:
	//  data:  The bytes of an individual packet.
	//  ci:  Metadata about the capture
	//  err:  An error encountered while reading packet data.  If err != nil,
	//    then data/ci will be ignored.
	ReadPacketData() (data []byte, ci gopacket.CaptureInfo, err error)
	// Close closes the ethernet sniffer and returns nil if no error was found.
	Close() error
}

/*
type Config struct {
	iface      string
	pcapOut    string
	enableAF   bool
	pcapFile   *os.File
	pcapWriter *pcapgo.Writer
	sniffer    Sniffer
	isRunning  bool
}

func (s *PcapSniffer) Open(config *Config) error {
	// Capture settings
	const (
		// Max packet length
		snaplen int32 = 65536
		// Set the interface in promiscuous mode
		promisc bool = true
		// Timeout duration
		flushAfter string = "10s"
		//BPF filter when capturing packets
		filter string = "ip"
	)

	// Open the interface
	flushDuration, err := time.ParseDuration(flushAfter)
	if err != nil {
		return fmt.Errorf("Invalid flush duration: %s", flushAfter)
	}
	handle, err := pcap.OpenLive(*iface, snaplen, promisc, flushDuration/2)
	if err != nil {
		return fmt.Errorf("Error opening pcap handle: %s", err)
	}
	if err := handle.SetBPFFilter(filter); err != nil {
		return fmt.Errorf("Error setting BPF filter: %s", err)
	}
	s.handle = handle

	return nil
}
func NewBPFHandle(options *types.SnifferDriverOptions) (types.PacketDataSourceCloser, error) {
	// XXX TODO pass more options...
	bpfSniffer, err := bsdbpf.NewBPFSniffer(options.Device, nil)
	return &BPFHandle{
		bpfSniffer: bpfSniffer,
	}, err
}
*/
//func (s *BPFSniffer) Open(options *SnifferDriverOptions) error {
func (s *BPFSniffer) Open(config *Config) error {
        // XXX TODO pass more options...
        handle, err := bsdbpf.NewBPFSniffer(*iface, nil)
        if err != nil {
		return fmt.Errorf("Error opening bsdbpf: %s", err)
        }
        s.handle = handle
        return nil
}

/*
func (s *PcapSniffer) Close() {
	s.handle.Close()
}
*/
func (s *BPFSniffer) Close() {
	s.handle.Close()
}

/*
func (s *PcapSniffer) ReadPacket() (data []byte, ci gopacket.CaptureInfo, err error) {
	return s.handle.ZeroCopyReadPacketData()
}
*/
func (s *BPFSniffer) ReadPacket() (data []byte, ci gopacket.CaptureInfo, err error) {
	return s.handle.ReadPacketData()
}
