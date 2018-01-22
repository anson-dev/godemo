#pragma once

#ifdef __cplusplus
extern "C" {
#endif

    //该函数仅供taf服务器调用
    int TafInvoke4Go(const char* szObjectName, const char* szFuncName, const char* req, int reqSize, char** rsp, int* rspSize);
#ifdef __cplusplus
}
#endif
