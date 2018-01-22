#ifndef _GoServantImp_H_
#define _GoServantImp_H_

#include "servant/Application.h"
#include "GoServant.h"

struct ChannelId;
/**
 *
 *
 */
class GoServantImp : public taf::Coroutine<HUYA::GoServant>
//class GoServantImp : public HUYA::GoServant
{
public:
    /**
     *
     */
    virtual ~GoServantImp() {}

    /**
     *
     */
    virtual void initialize();

    /**
     *
     */
    virtual void destroy();

    virtual taf::Int32 getRankGo(const HUYA::ActivityRankReq & tReq,HUYA::ActivityRankRsp &tRsp,taf::JceCurrentPtr current);
    virtual taf::Int32 getRankListGo(const HUYA::ActivityRankListReq & tReq,HUYA::ActivityRankListRsp &tRsp,taf::JceCurrentPtr current);
    virtual taf::Int32 updateRankGo(const HUYA::UpdateActivityRankReq & tReq,HUYA::ActivityRankRsp &tRsp,taf::JceCurrentPtr current);
};
/////////////////////////////////////////////////////
#endif
