package main

/*
go env -w GO111MODULE=on
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
*/
import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Msg Abc
type Msg interface {
	SaveType() string
	ToString(int64) string
}

// Decoder Abc
type Decoder interface {
	Decode([]byte) (Msg, error)
}

func checkMainArgs() (string, string) {
	var pcapFile string
	var outDir string
	flag.StringVar(&pcapFile, "i", "", "inputFile.pcap")
	flag.StringVar(&outDir, "o", "", "please choose in : rs, hs, ak")
	flag.Parse()
	if pcapFile == "" || outDir == "" {
		flag.Usage()
		os.Exit(1)
	}
	if (outDir != "rs") && (outDir != "ak") {
		flag.Usage()
		os.Exit(1)
	}
	return pcapFile, outDir
}

func panicWhenErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	pcapFileName, outDir := checkMainArgs()
	handler, err := pcap.OpenOffline(pcapFileName)
	panicWhenErr(err)
	defer handler.Close()

	if err := os.Mkdir(outDir, os.ModePerm); err != nil {
		if os.IsExist(err) {
			fmt.Println("dir ", outDir, "already exist")
		} else {
			panic(err)
		}
	}

	year, month, day := time.Now().Date()
	depthFileName := fmt.Sprintf("%d%02d%d_Depth.csv", year, month, day)
	tickFileName := fmt.Sprintf("%d%02d%d_Tick.csv", year, month, day)

	depthFile, err := os.OpenFile(path.Join(outDir, depthFileName), os.O_RDWR|os.O_CREATE, 0755)
	panicWhenErr(err)
	tickFile, err := os.OpenFile(path.Join(outDir, tickFileName), os.O_RDWR|os.O_CREATE, 0755)
	panicWhenErr(err)

	var decoder Decoder
	switch outDir {
	case "rs":
		decoder = RsDecoder{binary.BigEndian}
	case "ak":
		decoder = AkDecoder{binary.BigEndian}
	}

	/*
		packetSource := gopacket.NewPacketSource(handler, handler.LinkType())
		for packet := range packetSource.Packets() {
			if app := packet.ApplicationLayer(); app != nil {
				ts := packet.Metadata().CaptureInfo.Timestamp.UnixNano()
				msg, err := decoder.Decode(app.Payload())
				if err != nil {
					fmt.Println("failed to decode: ", app.Payload())
					continue
				}
				switch msg.SaveType() {
				case "Depth":
					depthFile.WriteString(fmt.Sprintf("%s\n", msg.ToString(ts)))
				case "Tick":
					tickFile.WriteString(fmt.Sprintf("%s\n", msg.ToString(ts)))
				}
			}
		}
	*/
	type im struct {
		ts  int64
		msg Msg
	}
	chdepth := make(chan *im, 100)
	chtick := make(chan *im, 100)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		packetSource := gopacket.NewPacketSource(handler, handler.LinkType())
		for packet := range packetSource.Packets() {
			if app := packet.ApplicationLayer(); app != nil {
				ts := packet.Metadata().CaptureInfo.Timestamp.UnixNano()
				msg, err := decoder.Decode(app.Payload())
				if err != nil {
					fmt.Println("failed to decode: ", app.Payload())
					continue
				}
				switch msg.SaveType() {
				case "Depth":
					chdepth <- &im{ts, msg}
				case "Tick":
					chtick <- &im{ts, msg}
				}
			}
		}
		close(chtick)
		close(chdepth)
		wg.Done()
	}()
	go func() {
		for im := range chdepth {
			depthFile.WriteString(fmt.Sprintf("%s\n", im.msg.ToString(im.ts)))
		}
		wg.Done()
	}()
	go func() {
		for im := range chtick {
			tickFile.WriteString(fmt.Sprintf("%s\n", im.msg.ToString(im.ts)))
		}
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("decode success")

}