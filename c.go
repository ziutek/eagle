package eagle

/*
	#include <asm-generic/int-ll64.h>
	#include <sys/ioctl.h>
	#include "iocontrol.h"
	#include "error.h"
*/
import "C"

type cint C.int

const (
	modGetChipType      = C.IOCTL_ITE_MOD_GETCHIPTYPE
	modGetDriverInfo    = C.IOCTL_ITE_MOD_GETDRIVERINFO
	modAcquireChannel   = C.IOCTL_ITE_MOD_ACQUIRECHANNEL
	modSetModule        = C.IOCTL_ITE_MOD_SETMODULE
	modGetGainRange     = C.IOCTL_ITE_MOD_GETGAINRANGE
	modGetOutputGain    = C.IOCTL_ITE_MOD_GETOUTPUTGAIN
	modAdjustOutputGain = C.IOCTL_ITE_MOD_ADJUSTOUTPUTGAIN
	modEnableTxMode     = C.IOCTL_ITE_MOD_ENABLETXMODE
	modStartTransfer    = C.IOCTL_ITE_MOD_STARTTRANSFER
	modStopTransfer     = C.IOCTL_ITE_MOD_STOPTRANSFER
)

const (
	ErrWriteRegTimeout          Error = C.ERR_WRITE_REG_TIMEOUT
	ErrWriteTunerTimeout        Error = C.ERR_WRITE_TUNER_TIMEOUT
	ErrWriteTunerFail           Error = C.ERR_WRITE_TUNER_FAIL
	ErrReadTuner_Timeout        Error = C.ERR_READ_TUNER_TIMEOUT
	ErrRSDCounterNotReady       Error = C.ERR_RSD_COUNTER_NOT_READY
	ErrVTBCounterNotReady       Error = C.ERR_VTB_COUNTER_NOT_READY
	ErrFECMonNotEnabled         Error = C.ERR_FEC_MON_NOT_ENABLED
	ErrInvalidDevType           Error = C.ERR_INVALID_DEV_TYPE
	ErrInvalid_TunerType        Error = C.ERR_INVALID_TUNER_TYPE
	ErrOpenFile_Fail            Error = C.ERR_OPEN_FILE_FAIL
	ErrWriteFileFail            Error = C.ERR_WRITEFILE_FAIL
	ErrReadFileFail             Error = C.ERR_READFILE_FAIL
	ErrCreateFileFail           Error = C.ERR_CREATEFILE_FAIL
	ErrMallocFail               Error = C.ERR_MALLOC_FAIL
	ErrInvalidFileSize          Error = C.ERR_INVALID_FILE_SIZE
	ErrInvalidReadSize          Error = C.ERR_INVALID_READ_SIZE
	ErrLoadFWDoneButFail        Error = C.ERR_LOAD_FW_DONE_BUT_FAIL
	ErrNotImplemented           Error = C.ERR_NOT_IMPLEMENTED
	ErrWriteMBXTunerTimeout     Error = C.ERR_WRITE_MBX_TUNER_TIMEOUT
	ErrDivMoreThan8Chips        Error = C.ERR_DIV_MORE_THAN_8_CHIPS
	ErrDivNoChips               Error = C.ERR_DIV_NO_CHIPS
	ErrSuperFrameCnt0           Error = C.ERR_SUPER_FRAME_CNT_0
	ErrInvalidFFTMode           Error = C.ERR_INVALID_FFT_MODE
	ErrInvalidConstellationMode Error = C.ERR_INVALID_CONSTELLATION_MODE
	ErrRSDPktCnt0               Error = C.ERR_RSD_PKT_CNT_0
	ErrFFTShiftTimeout          Error = C.ERR_FFT_SHIFT_TIMEOUT
	ErrWaitTPSTimeout           Error = C.ERR_WAIT_TPS_TIMEOUT
	ErrInvalidBandwidth         Error = C.ERR_INVALID_BW
	ErrInvalidBufLen            Error = C.ERR_INVALID_BUF_LEN
	ErrNullPtr                  Error = C.ERR_NULL_PTR
	ErrMTtuneFail               Error = C.ERR_MT_TUNE_FAIL
	ErrMTOpenFail               Error = C.ERR_MT_OPEN_FAIL
	ErrInvalidAGCVolt           Error = C.ERR_INVALID_AGC_VOLT
	ErrReadTunerFail            Error = C.ERR_READ_TUNER_FAIL
	ErrEMBXIntNotCleared        Error = C.ERR_EMBX_INT_NOT_CLEARED
	ErrInvPullupVolt            Error = C.ERR_INV_PULLUP_VOLT
	ErrFreqOutOfRange           Error = C.ERR_FREQ_OUT_OF_RANGE
	ErrMTNotAvailable           Error = C.ERR_MT_NOT_AVAILABLE
	ErrBackToBootcoveFail       Error = C.ERR_BACK_TO_BOOTCODE_FAIL
	ErrGetBufferValueFail       Error = C.ERR_GET_BUFFER_VALUE_FAIL
	ErrMemAallocFail            Error = C.ERR_MEM_ALLOC_FAIL
	ErrInvalidPOS               Error = C.ERR_INVALID_POS
	ErrDynaTopFail              Error = C.ERR_DYNA_TOP_FAIL
	ErrInvalidIndex             Error = C.ERR_INVALID_INDEX
	ErrWaitPVITTimeout          Error = C.ERR_WAIT_PVIT_TIMEOUT
	ErrFuncInterrupted          Error = C.ERR_FUNC_INTERRUPTED
	ErrXtalNotSupported         Error = C.ERR_XTAL_NOT_SUPPORT
	ErrCantFindOrigTOPS         Error = C.ERR_CANT_FIND_ORIG_TOPS
	ErrInvalidRegValue          Error = C.ERR_INVALID_REG_VALUE
	ErrTunnerNotSupported       Error = C.ERR_TUNER_NOT_SUPPORT
	ErrUndefinedSAWBandwidth    Error = C.ERR_UNDEFINED_SAW_BW
	ErrInvalidChipRevision      Error = C.ERR_INVALID_CHIP_REV
	ErrBufferInsufficient       Error = C.ERR_BUFFER_INSUFFICIENT
	ErrCounterNotAvailable      Error = C.ERR_COUNTER_NOT_AVAILABLE
	ErrLoadFWCompFail           Error = C.ERR_LOADFW_COMP_FAIL
	ErrCantFindEEPROM           Error = C.ERR_CANT_FIND_EEPROM
	ErrTunerTypeNotSupported    Error = C.ERR_TUNER_TYPE_NOT_SUPPORT
	ErrInvalidMiscReg           Error = C.ERR_INV_MISC_REG
	ErrCantFindUSBDev           Error = C.ERR_CANT_FIND_USB_DEV
	ErrInvalidXtalFreq          Error = C.ERR_INVALID_XTAL_FREQ
	ErrInvalidDeviceCount       Error = C.ERR_INVALID_DEVICE_COUNT

	ErrI2CNullHandle  Error = C.ERR_I2C_NULL_HANDLE
	ErrI2CDontSupport Error = C.ERR_I2C_DONT_SUPPORT

	ErrComDataHighFail Error = C.ERR_COM_DATA_HIGH_FAIL
	ErrComClkHighFail  Error = C.ERR_COM_CLK_HIGH_FAIL
	ErrComWriteNoAck   Error = C.ERR_COM_WRITE_NO_ACK
	ErrComDataLowFail  Error = C.ERR_COM_DATA_LOW_FAIL

	ErrUSBNullHandle                Error = C.ERR_USB_NULL_HANDLE
	ErrUSBWriteFileFail             Error = C.ERR_USB_WRITEFILE_FAIL
	ErrUSBReadFileFail              Error = C.ERR_USB_READFILE_FAIL
	ErrUSBInvalidReadSize           Error = C.ERR_USB_INVALID_READ_SIZE
	ErrUSBBadStatus                 Error = C.ERR_USB_BAD_STATUS
	ErrUSBInvalidSN                 Error = C.ERR_USB_INVALID_SN
	ErrUSBInvalidPktSize            Error = C.ERR_USB_INVALID_PKT_SIZE
	ErrUSBInvalidHeader             Error = C.ERR_USB_INVALID_HEADER
	ErrUSBNoIRPkt                   Error = C.ERR_USB_NO_IR_PKT
	ErrUSBInvalidTALen              Error = C.ERR_USB_INVALID_DATA_LEN
	ErrUSBEP4ReadFileFail           Error = C.ERR_USB_EP4_READFILE_FAIL
	ErrUSBEP4InvalidReadSize        Error = C.ERR_USB_EP4_INVALID_READ_SIZE
	ErrUSBBootInvalidPktType        Error = C.ERR_USB_BOOT_INVALID_PKT_TYPE
	ErrUSBBootBadConfigHeader       Error = C.ERR_USB_BOOT_BAD_CONFIG_HEADER
	ErrUSBBootBadConfigSize         Error = C.ERR_USB_BOOT_BAD_CONFIG_SIZE
	ErrUSBBootBadConfigSN           Error = C.ERR_USB_BOOT_BAD_CONFIG_SN
	ErrUSBBootBadConfigSubtype      Error = C.ERR_USB_BOOT_BAD_CONFIG_SUBTYPE
	ErrUSBBootBadConfigValue        Error = C.ERR_USB_BOOT_BAD_CONFIG_VALUE
	ErrUSBBootBadConfigChksum       Error = C.ERR_USB_BOOT_BAD_CONFIG_CHKSUM
	ErrUSBBootBadConfirmHeader      Error = C.ERR_USB_BOOT_BAD_CONFIRM_HEADER
	ErrUSBBootBadConfirmSize        Error = C.ERR_USB_BOOT_BAD_CONFIRM_SIZE
	ErrUSBBootBadConfirmSN          Error = C.ERR_USB_BOOT_BAD_CONFIRM_SN
	ErrUSBBootBadConfirmSubtype     Error = C.ERR_USB_BOOT_BAD_CONFIRM_SUBTYPE
	ErrUSBBootBadConfirmValue       Error = C.ERR_USB_BOOT_BAD_CONFIRM_VALUE
	ErrUSBBootBadConfirmChksum      Error = C.ERR_USB_BOOT_BAD_CONFIRM_CHKSUM
	ErrUSBBootBadBootHeader         Error = C.ERR_USB_BOOT_BAD_BOOT_HEADER
	ErrUSBBootBadBootSize           Error = C.ERR_USB_BOOT_BAD_BOOT_SIZE
	ErrUSBBootBadBootSN             Error = C.ERR_USB_BOOT_BAD_BOOT_SN
	ErrUSBBootBadBootPattern01      Error = C.ERR_USB_BOOT_BAD_BOOT_PATTERN_01
	ErrUSBBootBadBootPattern10      Error = C.ERR_USB_BOOT_BAD_BOOT_PATTERN_10
	ErrUSBBootBadBootChksum         Error = C.ERR_USB_BOOT_BAD_BOOT_CHKSUM
	ErrUSBBootBadBootPktType        Error = C.ERR_USB_INVALID_BOOT_PKT_TYPE
	ErrUSBBootBadConfigValue1       Error = C.ERR_USB_BOOT_BAD_CONFIG_VAlUE
	ErrUSBCoInitializeExFail        Error = C.ERR_USB_COINITIALIZEEX_FAIL
	ErrUSBCoCreateInstanceFail      Error = C.ERR_USB_COCREATEINSTANCE_FAIL
	ErrUSBCoCreateLSEEnumeratorFail Error = C.ERR_USB_COCREATCLSEENUMERATOR_FAIL
	ErrUSBQueryInterfaceFail        Error = C.ERR_USB_QUERY_INTERFACE_FAIL
	ErrUSBPKSCtrlNull               Error = C.ERR_USB_PKSCTRL_NULL
	ErrUSBInvalidHandle             Error = C.ERR_USB_INVALID_HANDLE
	ErrUSBTooMuchWriteData          Error = C.ERR_USB_TOO_MUCH_WRITE_DATA
	ErrUSBNoBurstRead               Error = C.ERR_USB_NO_BURST_READ
	ErrUSBNullPenum                 Error = C.ERR_USB_NULL_PENUM
)
