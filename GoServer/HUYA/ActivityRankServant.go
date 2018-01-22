// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 2.1.6.6
// Generated from `ActivityRankServant.jce'
// **********************************************************************

package HUYA

import (
	"taf/jce"
)

const (
	EUpdateRank = iota + 0
	EDelRankList
)

type ActivityRankReq struct {
	iActivityId int32
	sKey string
	lUid int64
}

func (self *ActivityRankReq) ResetDefault() {
}

func (self *ActivityRankReq) ReadFrom(jis *jce.JceInputStream) (err error) {
	err = jis.ReadField(0, false, &(self.iActivityId))
	if err != nil {
		return err
	}
	err = jis.ReadField(1, false, &(self.sKey))
	if err != nil {
		return err
	}
	err = jis.ReadField(2, false, &(self.lUid))
	if err != nil {
		return err
	}
	return err
}

func (self *ActivityRankReq) WriteTo(jos *jce.JceOutputStream) (err error) {
	err = jos.WriteField(&(self.iActivityId), 0)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.sKey), 1)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.lUid), 2)
	if err != nil {
		return err
	}
	return err
}

type ActivityRankRsp struct {
	lUid int64
	iRank int32
	iScore int32
	lUpdateTime int64
}

func (self *ActivityRankRsp) ResetDefault() {
}

func (self *ActivityRankRsp) ReadFrom(jis *jce.JceInputStream) (err error) {
	err = jis.ReadField(1, false, &(self.lUid))
	if err != nil {
		return err
	}
	err = jis.ReadField(2, false, &(self.iRank))
	if err != nil {
		return err
	}
	err = jis.ReadField(3, false, &(self.iScore))
	if err != nil {
		return err
	}
	err = jis.ReadField(4, false, &(self.lUpdateTime))
	if err != nil {
		return err
	}
	return err
}

func (self *ActivityRankRsp) WriteTo(jos *jce.JceOutputStream) (err error) {
	err = jos.WriteField(&(self.lUid), 1)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.iRank), 2)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.iScore), 3)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.lUpdateTime), 4)
	if err != nil {
		return err
	}
	return err
}

type ActivityRankListReq struct {
	iActivityId int32
	sKey string
	iFrom int32
	iCount int32
}

func (self *ActivityRankListReq) ResetDefault() {
}

func (self *ActivityRankListReq) ReadFrom(jis *jce.JceInputStream) (err error) {
	err = jis.ReadField(0, false, &(self.iActivityId))
	if err != nil {
		return err
	}
	err = jis.ReadField(1, false, &(self.sKey))
	if err != nil {
		return err
	}
	err = jis.ReadField(2, false, &(self.iFrom))
	if err != nil {
		return err
	}
	err = jis.ReadField(3, false, &(self.iCount))
	if err != nil {
		return err
	}
	return err
}

func (self *ActivityRankListReq) WriteTo(jos *jce.JceOutputStream) (err error) {
	err = jos.WriteField(&(self.iActivityId), 0)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.sKey), 1)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.iFrom), 2)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.iCount), 3)
	if err != nil {
		return err
	}
	return err
}

type ActivityRankListRsp struct {
	iTotal int32
	vRankList []ActivityRankRsp
}

func (self *ActivityRankListRsp) ResetDefault() {
}

func (self *ActivityRankListRsp) ReadFrom(jis *jce.JceInputStream) (err error) {
	err = jis.ReadField(0, false, &(self.iTotal))
	if err != nil {
		return err
	}
	err = jis.ReadField(1, false, &(self.vRankList))
	if err != nil {
		return err
	}
	return err
}

func (self *ActivityRankListRsp) WriteTo(jos *jce.JceOutputStream) (err error) {
	err = jos.WriteField(&(self.iTotal), 0)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.vRankList), 1)
	if err != nil {
		return err
	}
	return err
}

type UpdateActivityRankReq struct {
	iActivityId int32
	sKey string
	lUid int64
	iScoreDelta int32
}

func (self *UpdateActivityRankReq) ResetDefault() {
}

func (self *UpdateActivityRankReq) ReadFrom(jis *jce.JceInputStream) (err error) {
	err = jis.ReadField(0, false, &(self.iActivityId))
	if err != nil {
		return err
	}
	err = jis.ReadField(1, false, &(self.sKey))
	if err != nil {
		return err
	}
	err = jis.ReadField(2, false, &(self.lUid))
	if err != nil {
		return err
	}
	err = jis.ReadField(3, false, &(self.iScoreDelta))
	if err != nil {
		return err
	}
	return err
}

func (self *UpdateActivityRankReq) WriteTo(jos *jce.JceOutputStream) (err error) {
	err = jos.WriteField(&(self.iActivityId), 0)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.sKey), 1)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.lUid), 2)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.iScoreDelta), 3)
	if err != nil {
		return err
	}
	return err
}

type DelActivityRankListReq struct {
	iActivityId int32
	sKey string
	vUid []int64
	iDelOp int32
}

func (self *DelActivityRankListReq) ResetDefault() {
}

func (self *DelActivityRankListReq) ReadFrom(jis *jce.JceInputStream) (err error) {
	err = jis.ReadField(0, false, &(self.iActivityId))
	if err != nil {
		return err
	}
	err = jis.ReadField(1, false, &(self.sKey))
	if err != nil {
		return err
	}
	err = jis.ReadField(2, false, &(self.vUid))
	if err != nil {
		return err
	}
	err = jis.ReadField(3, false, &(self.iDelOp))
	if err != nil {
		return err
	}
	return err
}

func (self *DelActivityRankListReq) WriteTo(jos *jce.JceOutputStream) (err error) {
	err = jos.WriteField(&(self.iActivityId), 0)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.sKey), 1)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.vUid), 2)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.iDelOp), 3)
	if err != nil {
		return err
	}
	return err
}

type ARSyncCacheReq struct {
	sSrcId string
	iOpType int32
	sReserved string
}

func (self *ARSyncCacheReq) ResetDefault() {
}

func (self *ARSyncCacheReq) ReadFrom(jis *jce.JceInputStream) (err error) {
	err = jis.ReadField(0, false, &(self.sSrcId))
	if err != nil {
		return err
	}
	err = jis.ReadField(1, false, &(self.iOpType))
	if err != nil {
		return err
	}
	err = jis.ReadField(2, false, &(self.sReserved))
	if err != nil {
		return err
	}
	return err
}

func (self *ARSyncCacheReq) WriteTo(jos *jce.JceOutputStream) (err error) {
	err = jos.WriteField(&(self.sSrcId), 0)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.iOpType), 1)
	if err != nil {
		return err
	}
	err = jos.WriteField(&(self.sReserved), 2)
	if err != nil {
		return err
	}
	return err
}

type ActivityRankServant interface {
	getRank(tReq ActivityRankReq, tRsp *ActivityRankRsp) int32
	getRankList(tReq ActivityRankListReq, tRsp *ActivityRankListRsp) int32
	updateRank(tReq UpdateActivityRankReq, tRsp *ActivityRankRsp) int32
	delRankList(tReq DelActivityRankListReq) int32
	syncCache(tReq ARSyncCacheReq) int32
}



