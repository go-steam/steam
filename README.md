# steam_api wrapper for the Go language

This is a pre-Alpha release
---------------------------

This is a work in progress. At this time I am using Linux64. Contributors would be welcomed.


Package Concept
===============

The steam_api is a c++ product. 
  * wrapper.cpp uses the steam_api.
  * wrapper.h defines the extern "C" functions used by go. 
  * gosteam.h defines the functions export by go. 
  * steam.go is where we call our steam functions from go. 
  * steam_test.go tests the functions.


Installation
------------

We currently support only steamworks_sdk_129
		$ go get github.com/go-steam/steam

Begin by: 
  * downloading steamworks_sdk_ver to your GOPATH, ver is the version number.
  * extract steamworks_sdk_ver
  * rename sdk folder to steamworks_sdk_ver
  * you may need to copy the appropriate contents of the redistributable_bin of the sdk to your /usr/lib so that the .so files can be found by the test programs.
  * copy the header files in steamworks_sdk_ver/public/steam to the davsk/steam/v128/linux64/steam
  * go test will fail if the steam client is not running on your machine.


Version Control
---------------

Current Version is v129.0.0.0 and is compatable only with steamworks_sdk_129.
We recommend that all developers use version control tags to avoid breaking their project.
Our version format is SDK.MAJOR,MINOR,PATCH.
We recommend the use of godep to preserve your projects dependencies.

For more details about semantic versioning, see the documentation at:
    [http://semver.org](http://semver.org)


Platforms
---------

We are developing on Linux AMD64.
We will also be testing on the Windows AMD64.


Documentation
-------------

Detailed instructions for the use of the steam_api is found at 
[https://partner.steamgames.com/documentation/api](https://partner.steamgames.com/documentation/api)

Addition documentation is available through package documention.