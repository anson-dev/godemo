#ifndef __UserInfoProxy_H__
#define __UserInfoProxy_H__

//#include <memory>
//#include "util/tc_thread_rwlock.h"
//#include "util/tc_hashmap.h"
//#include "jmem/jmem_hashmap.h"
#include "WebdbProxyServant.h"
#include "JceTuple.h"
#include "JceCache.h"

class UserInfoProxy: public taf::TC_Singleton<UserInfoProxy>
{
public:
    UserInfoProxy()
    {
        init();
    }

    int getNickNames(const std::vector<taf::Int64>& vUids, std::map<taf::Int64,std::string>& mapUid2NickName)
    {
        taf::Int64 lNow = time(0);
        std::vector<taf::Int64> vToRefreshUid;
        for(auto& lUid: vUids)
        {
            UserInfoTuple tUserInfo;
            if(_mapUid2UserInfo.key(lUid).get(tUserInfo) && tUserInfo.get<4>() > lNow)
            {
                mapUid2NickName[lUid] = tUserInfo.get<0>();
            }
            else
            {
                vToRefreshUid.push_back(lUid);
            }
        }

        int iRet = refreshUserInfo(vToRefreshUid);
        for(auto& lUid : vToRefreshUid)
        {
            UserInfoTuple tUserInfo;
            if(_mapUid2UserInfo.key(lUid).get(tUserInfo) && tUserInfo.get<4>() > lNow)
            {
                mapUid2NickName[lUid] = tUserInfo.get<0>();
            }
        }
        return iRet;
    }

    std::string getNickName(long lUid)
    {
        std::string sNickName;
        getUserInfoVal<0>(lUid, sNickName);
        return sNickName;
    }

    //uid转yy号
    long getImid(long lUid)
    {
        long lImid = 0;
        getUserInfoVal<3>(lUid, lImid);
        return lImid;
    }

    //yy号转uid
    long getUidByImid(long lImid)
    {
        long lUid = 0;
        HUYA::getUserInfoReq tReq;
        tReq.lImid = lImid;
        HUYA::getUserInfoRsp tRsp;
        int iRet = -1;
        try
        {
            iRet = _prx->getUserInfo(tReq, tRsp);
            lUid = tRsp.tUserInfo.lUid;
        }
        catch(std::exception &ex)
        {
            LOG->error()<<__FUNCTION__<<"|ex:"<<ex.what()<<endl;
        }
        if(iRet != HUYA::EC_OK)
        {
            LOG->error()<<__FUNCTION__<<"|ret "<<iRet<<endl;
        }
        return lUid;
    }

    //获取用户自定义的高清头像
    std::string getUserHDLogo(long lUid)
    {
        std::string sLogo;
        getUserInfoVal<1>(lUid, sLogo);
        return sLogo;
    }

    //获取用户系统头像索引
    int getUserLogoIndex(long lUid)
    {
        int iLogoIndex;
        getUserInfoVal<2>(lUid, iLogoIndex);
        return iLogoIndex;
    }

    //获取主播头像
    std::string getPresenterLogo(unsigned int uid)
    {
        std::string url;
        if( uid == 0 )
        {
            return url;
        }


        int mod = 100;
        try
        {
            char char_uid[32];
            snprintf( char_uid, sizeof( char_uid ), "%u", uid );
            std::string str_uid( char_uid );
            std::string hashid = taf::TC_MD5::md5str( str_uid );

            int32_t hash = 0;
            char str_hash[32];
            sscanf( hashid.substr( 0, 7 ).c_str(), "%x", &hash );
            hash = hash % mod;

            if( hash < 0 )
            {
                hash = 0 - hash;
            }

            hash += 1000;
            snprintf( str_hash, sizeof( str_hash ), "%d", hash );

            std::string filePath = hashid.substr( 0, 2 );
            std::string fileName = hashid.substr( 2, 30 );
            url = std::string( str_hash ) + "/" + filePath + "/" + fileName + "_180_135.jpg";
        }
        catch (exception& ex)
        {
            LOG->error()<<__FUNCTION__<<"|"<<uid<<"|exception: |" << ex.what()<<endl;
            url = "";
        }
        catch (...)
        {
            LOG->error()<<__FUNCTION__<<"|"<<uid<<"|exception: unknown.|"<<endl;
            url = "";
        }

        if(!url.empty())
        {
            url = "http://img.huya.com/avatar/" + url;
        }
        else
        {
            url = "http://assets.dwstatic.com/amkit/p/duya/common/img/default_profile.jpg";
        }
        return url;
    }

private:
    bool init()
    {
        try
        {
            /*
            TC_Config tConf;
            tConf.parseFile(ServerConfig::BasePath + "/" + ServerConfig::ServerName + ".conf");
            _mapUid2UserInfo.init("UserInfo", tConf);
            */
            _mapUid2UserInfo.init("UserInfo");

            _comm.stringToProxy<HUYA::WebdbProxyServantPrx>("HUYA.WebdbProxyServer.WebdbProxyServantObj@tcp -h 10.21.40.167 -t 60000 -p 3869",_prx);
            return true;
        }
        catch(std::exception &ex)
        {
            LOG->error()<<__FUNCTION__<<"|ex:"<<ex.what()<<endl;
            return false;
        }
    }

    int refreshUserInfo(const std::vector<taf::Int64>& vUid)
    {
        if(vUid.empty())
        {
            return HUYA::EC_OK;
        }

        //获取用户信息
        HUYA::batchGetUserInfoReq tReq;
        tReq.vUids = vUid;
        HUYA::batchGetUserInfoRsp tRsp;
        int iRet = -1;
        try
        {
            iRet = _prx->batchGetUserInfo(tReq, tRsp);
        }
        catch(std::exception &ex)
        {
            LOG->error()<<__FUNCTION__<<"|ex:"<<ex.what()<<endl;
        }
        if(iRet != HUYA::EC_OK)
        {
            LOG->error()<<__FUNCTION__<<"|ret "<<iRet<<endl;
            return iRet;
        }

        //对于头像信息为空的用户获取其头像信息
        std::vector<long> vLogoUid;
        for(auto& kv: tRsp.mUserInfos)
        {
            if(kv.second.sHdlogo.empty() && kv.second.iLogoIndex == 0)
            {
                vLogoUid.push_back(kv.first);
            }
        }
        if(!vLogoUid.empty())
        {
            HUYA::BatchGetUserLogoReq tLogoReq;
            tLogoReq.vUids = vLogoUid;
            HUYA::BatchGetUserLogoRsp tLogoRsp;
            try
            {
                iRet = _prx->batchGetUserLogo(tLogoReq, tLogoRsp);
            }
            catch(std::exception &ex)
            {
                LOG->error()<<__FUNCTION__<<"|ex:"<<ex.what()<<endl;
                iRet = -1;
            }
            for(auto& kv : tLogoRsp.mUserLogoInfo)
            {
                if(kv.second.iLogoType == 0 && !kv.second.sLogo.empty())
                {
                    tRsp.mUserInfos[kv.first].sHdlogo = kv.second.sLogo;
                }
            }
        }

        taf::Int64 lExpiredTime = time(0) + kEffectiveDuration;
        for(auto& kv: tRsp.mUserInfos)
        {
            taf::Int64 lUid = kv.first;
            _mapUid2UserInfo.key(lUid).replace(kv.second.sNick, kv.second.sHdlogo, kv.second.iLogoIndex, kv.second.lImid, lExpiredTime);
        }

        return iRet;
    }

    template<int iIndex, typename T>
    bool getUserInfoVal(long lUid, T& tVal)
    {
        long lNow = time(0);
        UserInfoTuple tUserInfo;
        if(_mapUid2UserInfo.key(lUid).get(tUserInfo) && tUserInfo.get<4>() > lNow)
        {
            tVal = tUserInfo.get<iIndex>();
            return true;
        }

        refreshUserInfo({lUid});
        if(_mapUid2UserInfo.key(lUid).get(tUserInfo) && tUserInfo.get<4>() > lNow)
        {
            tVal = tUserInfo.get<iIndex>();
            return true;
        }
        else
        {
            return false;
        }
    }

    //0: 昵称
    //1: 高清头像
    //2: 系统头像索引
    //3: yy号
    //4: 失效时间
    typedef JceTuple<std::string, std::string, int, long, long>  UserInfoTuple;

    static const int kEffectiveDuration = 600; //s
    JceCache<JceTuple<long>, UserInfoTuple> _mapUid2UserInfo;
    HUYA::WebdbProxyServantPrx _prx;
    Communicator _comm;
};

#endif
