#ifndef _GoServer_H_
#define _GoServer_H_

#include <iostream>
#include "servant/Application.h"

using namespace taf;

/**
 *
 **/
class GoServer : public Application
{
public:
	/**
	 *
	 **/
	virtual ~GoServer() {};

	/**
	 *
	 **/
	virtual void initialize();

	/**
	 *
	 **/
	virtual void destroyApp();
};

extern GoServer g_app;

////////////////////////////////////////////
#endif
