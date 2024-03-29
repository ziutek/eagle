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
	bw := int(8e6)
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
	const numpkt = 1
	var buf [numpkt * 188]byte

	setRealtimeSched("output", 50)

	checkErr(d.TxSetModeEnable(true))
	checkErr(d.StartTransfer())

	const bitrate = 31668448
	delay := (numpkt*time.Second*188*8 + bitrate/2) / bitrate
	var delta time.Duration

	sendt := dtime()
	for i := 0; ; i++ {
		_, err := io.ReadFull(f, buf[:])
		checkErr(err)
		nanosleep(sendt - dtime())
		checkErr(d.TxSend(buf[:]))
		remain, err := d.TxRingBufRemain()
		checkErr(err)
		delta = (delta*9 + delay*(500*188-time.Duration(remain))/(1e5*188)) / 10
		sendt += delay + delta
		fmt.Println(remain, remain/188, delta)
	}

	checkErr(d.StopTransfer())
	checkErr(d.TxSetModeEnable(false))
}
