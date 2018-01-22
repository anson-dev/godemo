#include "ServantPrxGlue.h"
#include "servant/Application.h"
#include "servant/ServantProxy.h"
#include "servant/Servant.h"

using namespace std;

int TafInvoke4Go(const char* szObjectName, const char* szFuncName, const char* reqBuf, int reqBufSize, char** rspBuf, int* rspBufSize)
{
    try
    {
        ServantPrx prx = Application::getCommunicator()->stringToProxy<ServantPrx>(szObjectName);

        vector<char> vReqBuf(reqBuf, reqBuf+reqBufSize);
        ResponsePacket rsp;

        prx->taf_invoke(taf::JCENORMAL, szFuncName, vReqBuf, map<string, string>(), map<string,string>(),rsp);

        *rspBufSize = rsp.sBuffer.size();
        *rspBuf = (char*) malloc(rsp.sBuffer.size());
        memcpy(*rspBuf,&((rsp.sBuffer)[0]),*rspBufSize);
        return 0;
    }
    catch(exception &ex)
    {
        LOG->debug()<<__FUNCTION__<<"|"<<ex.what()<<endl;
        return -1;
    }
}
