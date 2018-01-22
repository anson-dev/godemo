#include "GoServantImp.h"
#include "servant/Application.h"
//#include "test.h"

using namespace std;
using namespace HUYA;

//////////////////////////////////////////////////////
void GoServantImp::initialize()
{
    //initialize servant here:
    //...
}

//////////////////////////////////////////////////////
void GoServantImp::destroy()
{
    //destroy servant here:
    //...
}

taf::Int32 GoServantImp::getRankGo(const HUYA::ActivityRankReq & tReq,HUYA::ActivityRankRsp &tRsp,taf::JceCurrentPtr current)
{
    return EC_OK;
}

taf::Int32 GoServantImp::getRankListGo(const HUYA::ActivityRankListReq & tReq,HUYA::ActivityRankListRsp &tRsp,taf::JceCurrentPtr current)
{
    return EC_OK;
}

taf::Int32 GoServantImp::updateRankGo(const HUYA::UpdateActivityRankReq & tReq,HUYA::ActivityRankRsp &tRsp,taf::JceCurrentPtr current)
{
    return EC_OK;
}
