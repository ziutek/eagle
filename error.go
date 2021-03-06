package eagle

type Error uint32

func (e Error) check() error {
	if e == 0 {
		return nil
	}
	return e
}

func (e Error) Error() string {
	if e == 0 {
		return "NoError"
	}
	if s := strerr[e]; s != "" {
		return s
	}
	return "ErrUnknown"
}

// egrep '^[[:space:]]+Err' c.go |awk '{printf("\t%s:\t\"%s\",\n", $1, $1);}' >>error.go
var strerr = map[Error]string{
	ErrWriteRegTimeout:              "ErrWriteRegTimeout",
	ErrWriteTunerTimeout:            "ErrWriteTunerTimeout",
	ErrWriteTunerFail:               "ErrWriteTunerFail",
	ErrReadTuner_Timeout:            "ErrReadTuner_Timeout",
	ErrRSDCounterNotReady:           "ErrRSDCounterNotReady",
	ErrVTBCounterNotReady:           "ErrVTBCounterNotReady",
	ErrFECMonNotEnabled:             "ErrFECMonNotEnabled",
	ErrInvalidDevType:               "ErrInvalidDevType",
	ErrInvalid_TunerType:            "ErrInvalid_TunerType",
	ErrOpenFile_Fail:                "ErrOpenFile_Fail",
	ErrWriteFileFail:                "ErrWriteFileFail",
	ErrReadFileFail:                 "ErrReadFileFail",
	ErrCreateFileFail:               "ErrCreateFileFail",
	ErrMallocFail:                   "ErrMallocFail",
	ErrInvalidFileSize:              "ErrInvalidFileSize",
	ErrInvalidReadSize:              "ErrInvalidReadSize",
	ErrLoadFWDoneButFail:            "ErrLoadFWDoneButFail",
	ErrNotImplemented:               "ErrNotImplemented",
	ErrWriteMBXTunerTimeout:         "ErrWriteMBXTunerTimeout",
	ErrDivMoreThan8Chips:            "ErrDivMoreThan8Chips",
	ErrDivNoChips:                   "ErrDivNoChips",
	ErrSuperFrameCnt0:               "ErrSuperFrameCnt0",
	ErrInvalidFFTMode:               "ErrInvalidFFTMode",
	ErrInvalidConstellationMode:     "ErrInvalidConstellationMode",
	ErrRSDPktCnt0:                   "ErrRSDPktCnt0",
	ErrFFTShiftTimeout:              "ErrFFTShiftTimeout",
	ErrWaitTPSTimeout:               "ErrWaitTPSTimeout",
	ErrInvalidBandwidth:             "ErrInvalidBandwidth",
	ErrInvalidBufLen:                "ErrInvalidBufLen",
	ErrNullPtr:                      "ErrNullPtr",
	ErrMTtuneFail:                   "ErrMTtuneFail",
	ErrMTOpenFail:                   "ErrMTOpenFail",
	ErrInvalidAGCVolt:               "ErrInvalidAGCVolt",
	ErrReadTunerFail:                "ErrReadTunerFail",
	ErrEMBXIntNotCleared:            "ErrEMBXIntNotCleared",
	ErrInvPullupVolt:                "ErrInvPullupVolt",
	ErrFreqOutOfRange:               "ErrFreqOutOfRange",
	ErrMTNotAvailable:               "ErrMTNotAvailable",
	ErrBackToBootcoveFail:           "ErrBackToBootcoveFail",
	ErrGetBufferValueFail:           "ErrGetBufferValueFail",
	ErrMemAallocFail:                "ErrMemAallocFail",
	ErrInvalidPOS:                   "ErrInvalidPOS",
	ErrDynaTopFail:                  "ErrDynaTopFail",
	ErrInvalidIndex:                 "ErrInvalidIndex",
	ErrWaitPVITTimeout:              "ErrWaitPVITTimeout",
	ErrFuncInterrupted:              "ErrFuncInterrupted",
	ErrXtalNotSupported:             "ErrXtalNotSupported",
	ErrCantFindOrigTOPS:             "ErrCantFindOrigTOPS",
	ErrInvalidRegValue:              "ErrInvalidRegValue",
	ErrTunnerNotSupported:           "ErrTunnerNotSupported",
	ErrUndefinedSAWBandwidth:        "ErrUndefinedSAWBandwidth",
	ErrInvalidChipRevision:          "ErrInvalidChipRevision",
	ErrBufferInsufficient:           "ErrBufferInsufficient",
	ErrCounterNotAvailable:          "ErrCounterNotAvailable",
	ErrLoadFWCompFail:               "ErrLoadFWCompFail",
	ErrCantFindEEPROM:               "ErrCantFindEEPROM",
	ErrTunerTypeNotSupported:        "ErrTunerTypeNotSupported",
	ErrInvalidMiscReg:               "ErrInvalidMiscReg",
	ErrCantFindUSBDev:               "ErrCantFindUSBDev",
	ErrInvalidXtalFreq:              "ErrInvalidXtalFreq",
	ErrInvalidDeviceCount:           "ErrInvalidDeviceCount",
	ErrI2CNullHandle:                "ErrI2CNullHandle",
	ErrI2CDontSupport:               "ErrI2CDontSupport",
	ErrComDataHighFail:              "ErrComDataHighFail",
	ErrComClkHighFail:               "ErrComClkHighFail",
	ErrComWriteNoAck:                "ErrComWriteNoAck",
	ErrComDataLowFail:               "ErrComDataLowFail",
	ErrUSBNullHandle:                "ErrUSBNullHandle",
	ErrUSBWriteFileFail:             "ErrUSBWriteFileFail",
	ErrUSBReadFileFail:              "ErrUSBReadFileFail",
	ErrUSBInvalidReadSize:           "ErrUSBInvalidReadSize",
	ErrUSBBadStatus:                 "ErrUSBBadStatus",
	ErrUSBInvalidSN:                 "ErrUSBInvalidSN",
	ErrUSBInvalidPktSize:            "ErrUSBInvalidPktSize",
	ErrUSBInvalidHeader:             "ErrUSBInvalidHeader",
	ErrUSBNoIRPkt:                   "ErrUSBNoIRPkt",
	ErrUSBInvalidTALen:              "ErrUSBInvalidTALen",
	ErrUSBEP4ReadFileFail:           "ErrUSBEP4ReadFileFail",
	ErrUSBEP4InvalidReadSize:        "ErrUSBEP4InvalidReadSize",
	ErrUSBBootInvalidPktType:        "ErrUSBBootInvalidPktType",
	ErrUSBBootBadConfigHeader:       "ErrUSBBootBadConfigHeader",
	ErrUSBBootBadConfigSize:         "ErrUSBBootBadConfigSize",
	ErrUSBBootBadConfigSN:           "ErrUSBBootBadConfigSN",
	ErrUSBBootBadConfigSubtype:      "ErrUSBBootBadConfigSubtype",
	ErrUSBBootBadConfigValue:        "ErrUSBBootBadConfigValue",
	ErrUSBBootBadConfigChksum:       "ErrUSBBootBadConfigChksum",
	ErrUSBBootBadConfirmHeader:      "ErrUSBBootBadConfirmHeader",
	ErrUSBBootBadConfirmSize:        "ErrUSBBootBadConfirmSize",
	ErrUSBBootBadConfirmSN:          "ErrUSBBootBadConfirmSN",
	ErrUSBBootBadConfirmSubtype:     "ErrUSBBootBadConfirmSubtype",
	ErrUSBBootBadConfirmValue:       "ErrUSBBootBadConfirmValue",
	ErrUSBBootBadConfirmChksum:      "ErrUSBBootBadConfirmChksum",
	ErrUSBBootBadBootHeader:         "ErrUSBBootBadBootHeader",
	ErrUSBBootBadBootSize:           "ErrUSBBootBadBootSize",
	ErrUSBBootBadBootSN:             "ErrUSBBootBadBootSN",
	ErrUSBBootBadBootPattern01:      "ErrUSBBootBadBootPattern01",
	ErrUSBBootBadBootPattern10:      "ErrUSBBootBadBootPattern10",
	ErrUSBBootBadBootChksum:         "ErrUSBBootBadBootChksum",
	ErrUSBBootBadBootPktType:        "ErrUSBBootBadBootPktType",
	ErrUSBBootBadConfigValue1:       "ErrUSBBootBadConfigValue1",
	ErrUSBCoInitializeExFail:        "ErrUSBCoInitializeExFail",
	ErrUSBCoCreateInstanceFail:      "ErrUSBCoCreateInstanceFail",
	ErrUSBCoCreateLSEEnumeratorFail: "ErrUSBCoCreateLSEEnumeratorFail",
	ErrUSBQueryInterfaceFail:        "ErrUSBQueryInterfaceFail",
	ErrUSBPKSCtrlNull:               "ErrUSBPKSCtrlNull",
	ErrUSBInvalidHandle:             "ErrUSBInvalidHandle",
	ErrUSBTooMuchWriteData:          "ErrUSBTooMuchWriteData",
	ErrUSBNoBurstRead:               "ErrUSBNoBurstRead",
	ErrUSBNullPenum:                 "ErrUSBNullPenum",
}
