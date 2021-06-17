package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"github.com/robfig/cron/v3"
	"os"
	"time"
)

func Run(config *Config) {
	handler, errHandler := pcap.OpenLive(
		config.ListenInterface,
		defaultSnapLen,
		config.Promisc,
		pcap.BlockForever,
	)
	if errHandler != nil {
		panic(errHandler)
	}
	defer handler.Close()

	packets := gopacket.NewPacketSource(
		handler, handler.LinkType(),
	).Packets()

	fileLocation := config.WriteLocation
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc(config.CronSpec, func() {
		t := time.Now()
		tt := fmt.Sprintf("%d-%02d-%02d_%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second(),
		)
		fileName := fmt.Sprintf("%s-%s.pcap", "traffics", tt)
		path := fileLocation + fileName
		f, _ := os.Create(path)
		w := pcapgo.NewWriter(f)
		errWriteHeader := w.WriteFileHeader(65536, layers.LinkTypeEthernet) // new file, must do this.
		if errWriteHeader != nil {
			return
		}
		afterCh := time.After(time.Duration(config.Interval) * time.Second)
		for {
			select {
			case packet := <-packets:
				err := w.WritePacket(gopacket.CaptureInfo{
					Timestamp:      time.Time{},
					CaptureLength:  len(packet.Data()),
					Length:         len(packet.Data()),
					InterfaceIndex: 0,
					AncillaryData:  nil,
				}, packet.Data())
				if err != nil {
					return
				}
			case <-afterCh:
				__ := f.Close()
				if __ != nil {
					return
				}
				WriteLog("Create new file " + path)
				RunProcessor(config, path)
				return
			}
		}
	})

	c.Run()
	if err != nil {
		return
	}
}
