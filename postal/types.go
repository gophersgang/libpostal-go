// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sun, 28 Feb 2016 04:34:07 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package postal

/*
#cgo pkg-config: libpostal
#include "libpostal.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// NormalizeOptions as declared in postal/libpostal.h:63
type NormalizeOptions struct {
	Languages              [][]byte
	NumLanguages           int32
	AddressComponents      uint16
	LatinAscii             bool
	Transliterate          bool
	StripAccents           bool
	Decompose              bool
	Lowercase              bool
	TrimString             bool
	DropParentheticals     bool
	ReplaceNumericHyphens  bool
	DeleteNumericHyphens   bool
	SplitAlphaFromNumeric  bool
	ReplaceWordHyphens     bool
	DeleteWordHyphens      bool
	DeleteFinalPeriods     bool
	DeleteAcronymPeriods   bool
	DropEnglishPossessives bool
	DeleteApostrophes      bool
	ExpandNumex            bool
	RomanNumerals          bool
	ref23955150            *C.normalize_options_t
	allocs23955150         interface{}
}

// AddressParserResponse as declared in postal/libpostal.h:79
type AddressParserResponse struct {
	NumComponents  uint
	Components     [][]byte
	Labels         [][]byte
	ref50c6200b    *C.address_parser_response_t
	allocs50c6200b interface{}
}

// AddressParserOptions as declared in postal/libpostal.h:84
type AddressParserOptions struct {
	Language       []byte
	Country        []byte
	refcd8497af    *C.address_parser_options_t
	allocscd8497af interface{}
}
