// Package browserdata is responsible for initializing all the necessary
// components that handle different types of browser data extraction.
// This file, imports.go, is specifically used to import various data
// handler packages to ensure their initialization logic is executed.
// These imports are crucial as they trigger the `init()` functions
// within each package, which typically handle registration of their
// specific data handlers to a central registry.
package browserdata

import (
	_ "github.com/GrpDsG20/CookiesExtractorv2/browserdata/bookmark"
	_ "github.com/GrpDsG20/CookiesExtractorv2/browserdata/cookie"
	_ "github.com/GrpDsG20/CookiesExtractorv2/browserdata/creditcard"
	_ "github.com/GrpDsG20/CookiesExtractorv2/browserdata/download"
	_ "github.com/GrpDsG20/CookiesExtractorv2/browserdata/extension"
	_ "github.com/GrpDsG20/CookiesExtractorv2/browserdata/history"
	_ "github.com/GrpDsG20/CookiesExtractorv2/browserdata/localstorage"
	_ "github.com/GrpDsG20/CookiesExtractorv2/browserdata/password"
	_ "github.com/GrpDsG20/CookiesExtractorv2/browserdata/sessionstorage"
)
