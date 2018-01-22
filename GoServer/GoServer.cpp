#include "GoServer.h"
#include "GoServantImp.h"

using namespace std;

GoServer g_app;

/////////////////////////////////////////////////////////////////
void
GoServer::initialize()
{
	//initialize application here:
	//...

    addConfig(ServerConfig::ServerName + ".conf");
	addServant<GoServantImp>(ServerConfig::Application + "." + ServerConfig::ServerName + ".GoServantObj");
}
/////////////////////////////////////////////////////////////////
void
GoServer::destroyApp()
{
	//destroy application here:
	//...
}
/////////////////////////////////////////////////////////////////
int
main(int argc, char* argv[])
{
	try
	{
		g_app.main(argc, argv);
		g_app.waitForShutdown();
	}
	catch (std::exception& e)
	{
		cerr << "std::exception:" << e.what() << std::endl;
	}
	catch (...)
	{
		cerr << "unknown exception." << std::endl;
	}
	return -1;
}
/////////////////////////////////////////////////////////////////
