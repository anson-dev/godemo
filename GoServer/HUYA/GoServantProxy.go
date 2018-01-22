// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 2.1.6.6
// Generated from `GoServant.jce'
// **********************************************************************

package HUYA

// #cgo LDFLAGS: -L../ -lServantPrxGlue -lstdc++ -Wl,-rpath=./
// #cgo CFLAGS: -I../
// #include "ServantPrxGlue.h"
// #include "stdlib.h"
import "C"
import "unsafe"
import "bytes"
import "errors"
import "taf/jce"
type GoServantProxy struct {
	objName *C.char
	getRankGoFN *C.char
	getRankListGoFN *C.char
	updateRankGoFN *C.char
}

func NewGoServantProxy(objectName string) *GoServantProxy{
	return &GoServantProxy{
		objName: C.CString(objectName),
		getRankGoFN: C.CString("getRankGo"),
		getRankListGoFN: C.CString("getRankListGo"),
		updateRankGoFN: C.CString("updateRankGo"),
	}
}

func DelGoServantProxy(prx *GoServantProxy) {
	C.free(unsafe.Pointer(prx.objName))
	C.free(unsafe.Pointer(prx.getRankGoFN))
	C.free(unsafe.Pointer(prx.getRankListGoFN))
	C.free(unsafe.Pointer(prx.updateRankGoFN))
}

func (self *GoServantProxy) getRankGo(tReq ActivityRankReq, tRsp *ActivityRankRsp) (ret int32, err error) {
	var jos jce.JceOutputStream
	jos.Buffer = bytes.NewBuffer(nil)
	err = jos.WriteField(&tReq,1)
	if err != nil {
		return jce.JCESERVERENCODEERR, err
	}
	err = jos.WriteField(tRsp,2)
	if err != nil {
		return jce.JCESERVERENCODEERR, err
	}
	reqBufSize := C.int(jos.Len())
	reqBuf := C.CString(jos.String())
	defer C.free(unsafe.Pointer(reqBuf))
	rspBufSize := C.int(0)
	var rspBuf *C.char
	defer C.free(unsafe.Pointer(rspBuf))
	iRet := C.TafInvoke4Go(self.objName, self.getRankGoFN, reqBuf, reqBufSize, &rspBuf, &rspBufSize)
	if iRet != 0 {
		return ret, errors.New("Failed in taf invoke.")
	}
	var jis jce.JceInputStream
	jis.Reader = bytes.NewReader(C.GoBytes(unsafe.Pointer(rspBuf), rspBufSize))
	jis.ReadField(0, true, &ret)
	if err != nil {
		return jce.JCESERVERDECODEERR, err
	}
	jis.ReadField(2, true, tRsp)
	if err != nil {
		return jce.JCESERVERDECODEERR, err
	}
	return ret, err
}

func (self *GoServantProxy) getRankListGo(tReq ActivityRankListReq, tRsp *ActivityRankListRsp) (ret int32, err error) {
	var jos jce.JceOutputStream
	jos.Buffer = bytes.NewBuffer(nil)
	err = jos.WriteField(&tReq,1)
	if err != nil {
		return jce.JCESERVERENCODEERR, err
	}
	err = jos.WriteField(tRsp,2)
	if err != nil {
		return jce.JCESERVERENCODEERR, err
	}
	reqBufSize := C.int(jos.Len())
	reqBuf := C.CString(jos.String())
	defer C.free(unsafe.Pointer(reqBuf))
	rspBufSize := C.int(0)
	var rspBuf *C.char
	defer C.free(unsafe.Pointer(rspBuf))
	iRet := C.TafInvoke4Go(self.objName, self.getRankListGoFN, reqBuf, reqBufSize, &rspBuf, &rspBufSize)
	if iRet != 0 {
		return ret, errors.New("Failed in taf invoke.")
	}
	var jis jce.JceInputStream
	jis.Reader = bytes.NewReader(C.GoBytes(unsafe.Pointer(rspBuf), rspBufSize))
	jis.ReadField(0, true, &ret)
	if err != nil {
		return jce.JCESERVERDECODEERR, err
	}
	jis.ReadField(2, true, tRsp)
	if err != nil {
		return jce.JCESERVERDECODEERR, err
	}
	return ret, err
}

func (self *GoServantProxy) updateRankGo(tReq UpdateActivityRankReq, tRsp *ActivityRankRsp) (ret int32, err error) {
	var jos jce.JceOutputStream
	jos.Buffer = bytes.NewBuffer(nil)
	err = jos.WriteField(&tReq,1)
	if err != nil {
		return jce.JCESERVERENCODEERR, err
	}
	err = jos.WriteField(tRsp,2)
	if err != nil {
		return jce.JCESERVERENCODEERR, err
	}
	reqBufSize := C.int(jos.Len())
	reqBuf := C.CString(jos.String())
	defer C.free(unsafe.Pointer(reqBuf))
	rspBufSize := C.int(0)
	var rspBuf *C.char
	defer C.free(unsafe.Pointer(rspBuf))
	iRet := C.TafInvoke4Go(self.objName, self.updateRankGoFN, reqBuf, reqBufSize, &rspBuf, &rspBufSize)
	if iRet != 0 {
		return ret, errors.New("Failed in taf invoke.")
	}
	var jis jce.JceInputStream
	jis.Reader = bytes.NewReader(C.GoBytes(unsafe.Pointer(rspBuf), rspBufSize))
	jis.ReadField(0, true, &ret)
	if err != nil {
		return jce.JCESERVERDECODEERR, err
	}
	jis.ReadField(2, true, tRsp)
	if err != nil {
		return jce.JCESERVERDECODEERR, err
	}
	return ret, err
}


