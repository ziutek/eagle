package eagle

import (
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"syscall"

	"github.com/ziutek/dvb"
)

type Family byte

const (
	I  Family = iota // IT950x series
	II               // IT951x series)
)

type Device struct {
	fd     int
	family Family
}

func Open(n int, f Family) (*Device, error) {
	if uint(n) >= 16 {
		return nil, errors.New("not supported device number")
	}
	path := strconv.Itoa(n)
	switch f {
	case I:
		path = "/dev/usb-it950x" + path
	case II:
		path = "/dev/usb-it950x" + path
	default:
		return nil, errors.New("unknown device type")
	}
	fd, err := syscall.Open(path, syscall.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	d := &Device{fd: fd, family: f}
	runtime.SetFinalizer(d, (*Device).Close)
	return d, nil
}

func (d *Device) Close() error {
	return syscall.Close(d.fd)
}

func (d *Device) ioctl(cmd uintptr, dataptr interface{}) error {
	var a3 uintptr
	if dataptr != nil {
		a3 = reflect.ValueOf(dataptr).Pointer()
	}
	if _, _, e := syscall.Syscall(syscall.SYS_IOCTL, uintptr(d.fd), cmd, a3); e != 0 {
		return e
	}
	return nil
}

func (d *Device) checkFreqBand(frequency int64, bandwidth int32) error {
	var lowf, higf int64
	switch d.family {
	case I:
		lowf = 173e6
		higf = 900e6
	default:
		lowf = 93e6
		higf = 803e6
	}
	if frequency < lowf || frequency > higf {
		return ErrFreqOutOfRange
	}
	if bandwidth < 2e6 || bandwidth > 8e6 {
		return ErrInvalidBandwidth
	}
	return nil
}

type Type uint16

func (t Type) String() string {
	return "IT" + strconv.FormatUint(uint64(t), 16)
}

func (d *Device) TxChipType() (Type, error) {
	var data struct {
		chipType Type
		error    Error
		_        [16]byte
	}
	if err := d.ioctl(modGetChipType, &data); err != nil {
		return 0, err
	}
	return data.chipType, data.error.check()
}

type TxModDriverInfo struct {
	DriverVersion string
	APIVersion    string
	FWVersionLink string
	FWVersionOFDM string
	DateTime      string
	Company       string
	SupportHWInfo string
}

func (d *Device) TxModDriverInfo() (TxModDriverInfo, error) {
	var data struct {
		DriverVersion [16]byte
		APIVersion    [32]byte
		FWVersionLink [16]byte
		FWVersionOFDM [16]byte
		DateTime      [24]byte
		Company       [8]byte
		SupportHWInfo [32]byte
		Error         Error
		_             [128]byte
	}
	if err := d.ioctl(modGetDriverInfo, &data); err != nil {
		return TxModDriverInfo{}, err
	}
	info := TxModDriverInfo{
		btos(data.DriverVersion[:]),
		btos(data.APIVersion[:]),
		btos(data.FWVersionLink[:]),
		btos(data.FWVersionOFDM[:]),
		btos(data.DateTime[:]),
		btos(data.Company[:]),
		btos(data.SupportHWInfo[:]),
	}
	return info, data.Error.check()
}

// TxSetChannel sets channel parameters (Hz unit).
func (d *Device) TxSetChannel(frequency int64, bandwidth int32) error {
	if err := d.checkFreqBand(frequency, bandwidth); err != nil {
		return err
	}
	var data struct {
		chip      byte
		bandwidth uint16
		frequency uint32
		error     Error
		_         [16]byte
	}
	data.bandwidth = uint16(bandwidth / 1e3)
	data.frequency = uint32(frequency / 1e3)
	if err := d.ioctl(modAcquireChannel, &data); err != nil {
		return err
	}
	return data.error.check()
}

func (d *Device) TxSetModulation(mod dvb.Modulation, txmode dvb.TxMode, coderate dvb.CodeRate, guard dvb.Guard) error {
	var data struct {
		chip             byte
		transmissionMode byte
		constellation    byte
		interval         byte
		highCodeRate     byte
		error            Error
		_                [16]byte
	}
	switch mod {
	case dvb.QPSK:
		data.constellation = 0
	case dvb.QAM16:
		data.constellation = 1
	case dvb.QAM64:
		data.constellation = 2
	default:
		return errors.New("not supported modulation")
	}
	switch txmode {
	case dvb.TxMode2k:
		data.transmissionMode = 0
	case dvb.TxMode8k:
		data.transmissionMode = 1
	default:
		return errors.New("not supported transmission mode")
	}
	switch coderate {
	case dvb.FEC12:
		data.highCodeRate = 0
	case dvb.FEC23:
		data.highCodeRate = 1
	case dvb.FEC34:
		data.highCodeRate = 2
	case dvb.FEC56:
		data.highCodeRate = 3
	case dvb.FEC78:
		data.highCodeRate = 4
	default:
		return errors.New("not supported code rate")
	}
	switch guard {
	case dvb.Guard32:
		data.interval = 0
	case dvb.Guard16:
		data.interval = 1
	case dvb.Guard8:
		data.interval = 2
	case dvb.Guard4:
		data.interval = 3
	default:
		return errors.New("not supported guard interval")
	}
	if err := d.ioctl(modSetModule, &data); err != nil {
		return err
	}
	return data.error.check()
}

func (d *Device) TxGainRange(frequency int64, bandwidth int32) (min, max int, err error) {
	if err = d.checkFreqBand(frequency, bandwidth); err != nil {
		return
	}
	var data struct {
		error     Error
		frequency uint32
		bandwidth uint16
		maxGain   cint
		minGain   cint
		_         [16]byte
	}
	data.frequency = uint32(frequency / 1e3)
	data.bandwidth = uint16(bandwidth / 1e3)
	if err = d.ioctl(modGetGainRange, &data); err != nil {
		return
	}
	min = int(data.minGain)
	max = int(data.maxGain)
	err = data.error.check()
	return
}

func (d *Device) TxOutputGain() (int, error) {
	var data struct {
		gain  cint
		error Error
		_     [16]byte
	}
	if err := d.ioctl(modGetOutputGain, &data); err != nil {
		return 0, err
	}
	return int(data.gain), data.error.check()
}

func (d *Device) TxAdjustOutputGain(gain int) (int, error) {
	var data struct {
		GainValue cint
		error     Error
	}
	data.GainValue = cint(gain)
	if err := d.ioctl(modAdjustOutputGain, &data); err != nil {
		return 0, err
	}
	return int(data.GainValue), data.error.check()
}

func (d *Device) TxSetModeEnable(on bool) error {
	var data struct {
		OnOff byte
		error Error
		_     [16]byte
	}
	if on {
		data.OnOff = 1
	}
	if err := d.ioctl(modEnableTxMode, &data); err != nil {
		return err
	}
	return data.error.check()
}

func (d *Device) StartTransfer() error {
	return d.ioctl(modStartTransfer, nil)
}

func (d *Device) StopTransfer() error {
	return d.ioctl(modStopTransfer, nil)
}

// TxSend sends b and return remaining buffer size (only modified driver returns
// correct value).
func (d *Device) TxSend(b []byte) (int, error) {
	n, err := syscall.Write(d.fd, b)
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, Error(-n)
	}
	return n, nil
}
