// doc.go

// Alpha Package steam is a wrapper for steam_api for go lang.
//
// Package Concept
//
// The steam_api is a c++ product.
//  wrapper.cpp uses the steam_api.
//  wrapper.h defines the extern "C" functions used by go.
//  gosteam.h defines the functions export by go.
//  steam.go is where we call our steam functions from go.
//  steam_test.go tests the functions.
//
// Note that to allow time for updates with the steam_api we are using
//  #define VERSION_SAFE_STEAM_API_INTERFACES
//
// Detailed instructions for the use of the steam_api is found at
// https://partner.steamgames.com/documentation/api
//
// Version Control
//
// Current Version is v129.0.0.0 and is compatable only with steamworks_sdk_129.
// We recommend that all developers use version control tags to avoid breaking their project.
// Our version format is SDK.MAJOR,MINOR,PATCH.
// We recommend the use of godep to preserve your projects dependencies.
//
// For more details about semantic versioning, see the documentation at:
//  http://semver.org
//
// Platforms
//
// We are developing on Linux AMD64.
// We will also be testing on the Windows AMD64.
//
package steam
