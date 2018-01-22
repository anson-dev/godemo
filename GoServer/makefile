
#-----------------------------------------------------------------------

APP       := HUYA
TARGET    := GoServer
MFLAGS    :=
CONFIG    := 
STRIP_FLAG:= N
J2CPP_FLAG:= 
CFLAGS    := -fPIC
TARGET_GLUE 	:= ServantGlue.a
PRX_GLUE_FILE 	:= ServantPrxGlue.cpp

INCLUDE   += -I/usr/local/mqq/wbl/include 
LIB       += -L/usr/local/mqq/wbl/lib 

INCLUDE   += -I/home/tafjce/Comm/Util/hiredis
LIB       += /home/tafjce/Comm/Util/hiredis/libhiredis.a
LIB       += -lrt
#LIB   	  += ./ServantGlue.a

INCLUDE   += -I../Util/ChannelHelper
INCLUDE   += -I../Util/BroadcastHelper
#-----------------------------------------------------------------------

include ./makefile4go.taf
include /home/tafjce/Comm/Util/Util.mk
include /home/tafjce/HUYA/comm/comm.mk
include /home/tafjce/HUYA/ActivityBaseServer/ActivityBaseServer.mk
include /home/tafjce/HUYA/ActivityRankServer/ActivityRankServer.mk
include /home/tafjce/HUYA/ActivityTimerServer/ActivityTimerServer.mk
include /home/tafjce/HUYA/WebdbProxyServer/WebdbProxyServer.mk
include /home/tafjce/HUYA/LiveDeliverServer/LiveDeliverServer.mk
include /home/tafjce/HUYA/LiveDeliverServer/CommLiveDeliver.mk
include /home/zhengqiao/Util/LogHelper/LogHelper.mk
include /home/zhengqiao/Util/LockedSingleton/LockedSingleton.mk
include /home/zhengqiao/Util/JceHelper/JceHelper.mk

include ../Util/ChannelHelper/ChannelHelper.mk
include ../Util/BroadcastHelper/BroadcastHelper.mk
include ../Util/UserInfoHelper/UserInfoHelper.mk
#-----------------------------------------------------------------------
