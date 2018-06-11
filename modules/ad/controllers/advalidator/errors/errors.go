package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	// AdRateDimensionErr AdRateDimensionErr
	AdRateDimensionErr = t9e.G("rate not valid")
	// AdDurationErr AdDurationErr
	AdDurationErr = t9e.G("max duration mismatch")
	// AdMinDimensionErr AdMinDimensionErr
	AdMinDimensionErr = t9e.G("min dimension mismatch")
	// AdMaxDimensionErr AdMaxDimensionErr
	AdMaxDimensionErr = t9e.G("max dimension mismatch")
	// AdExactDimensionErr AdExactDimensionErr
	AdExactDimensionErr = t9e.G("exact dimension mismatch")
	// AdExtensionDimensionErr AdExtensionDimensionErr
	AdExtensionDimensionErr = t9e.G("extension not valid")
	// AdSizeErr AdSizeErr
	AdSizeErr = t9e.G("ad size not valid")
)
