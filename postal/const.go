package postal

const (
	// MaxLanguageLen as defined in postal/libpostal.h:13
	MaxLanguageLen = 4
	// AddressNone as defined in postal/libpostal.h:19
	AddressNone = 0
	// AddressAny as defined in postal/libpostal.h:20
	AddressAny = 1
	// AddressName as defined in postal/libpostal.h:21
	AddressName = 2
	// AddressHouseNumber as defined in postal/libpostal.h:22
	AddressHouseNumber = 4
	// AddressStreet as defined in postal/libpostal.h:23
	AddressStreet = 8
	// AddressUnit as defined in postal/libpostal.h:24
	AddressUnit = 16
	// AddressLocality as defined in postal/libpostal.h:26
	AddressLocality = 128
	// AddressAdmin1 as defined in postal/libpostal.h:27
	AddressAdmin1 = 256
	// AddressAdmin2 as defined in postal/libpostal.h:28
	AddressAdmin2 = 512
	// AddressAdmin3 as defined in postal/libpostal.h:29
	AddressAdmin3 = 1024
	// AddressAdmin4 as defined in postal/libpostal.h:30
	AddressAdmin4 = 2048
	// AddressAdminOther as defined in postal/libpostal.h:31
	AddressAdminOther = 4096
	// AddressCountry as defined in postal/libpostal.h:32
	AddressCountry = 8192
	// AddressPostalCode as defined in postal/libpostal.h:33
	AddressPostalCode = 16384
	// AddressNeighborhood as defined in postal/libpostal.h:34
	AddressNeighborhood = 32768
	// AddressAll as defined in postal/libpostal.h:35
	AddressAll = 65535
)
