package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/ziutek/dvb"
	"github.com/ziutek/eagle"
)

func main() {
	freq := int64(698e6)
	bw := int32(8e6)
	gain := -8

	d, err := eagle.Open(0, eagle.I)
	checkErr(err)
	typ, err := d.TxChipType()
	checkErr(err)
	fmt.Println("TxChipType: ", typ)
	info, err := d.TxModDriverInfo()
	checkErr(err)
	fmt.Printf("TxModDriverInfo: %+v\n", info)
	mingain, maxgain, err := d.TxGainRange(freq, bw)
	checkErr(err)
	fmt.Printf("GainRange: [%d, %d] dB\n", mingain, maxgain)

	gain, err = d.TxAdjustOutputGain(gain)
	checkErr(err)
	fmt.Printf("Gain set to: %d dB\n", gain)

	gain, err = d.TxOutputGain()
	checkErr(err)
	fmt.Printf("OutputGain: %d dB\n", gain)

	checkErr(d.TxSetChannel(freq, bw))
	checkErr(d.TxSetModulation(dvb.QAM64, dvb.TxMode8k, dvb.FEC78, dvb.Guard32))

	f, err := os.Open("/home/michal/out.ts")
	checkErr(err)
	const numpkt = 40
	var buf [numpkt * 188]byte
	//null := [188]byte{0x47, 0x1f, 0xff}

	setRealtimeSched("output", 50)

	checkErr(d.TxSetModeEnable(true))
	checkErr(d.StartTransfer())

	const bitrate = 31668448
	delay := (numpkt*time.Second*188*8 + bitrate/2) / bitrate
	var delta, sendt time.Duration

	sendt = dtime()
	for i := 0; ; i++ {
		_, err := io.ReadFull(f, buf[:])
		checkErr(err)
		nanosleep(sendt - dtime())
		remain, err := d.TxSend(buf[:])
		checkErr(err)
		delta = (delta + delay*(500*188-time.Duration(remain))/(2e5*188)) / 2
		sendt += delay + delta
		//fmt.Println(remain/188, delta)
	}

	checkErr(d.StopTransfer())
	checkErr(d.TxSetModeEnable(false))
}
