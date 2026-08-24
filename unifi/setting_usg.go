package unifi

// SettingUsg's flat GeoIPFiltering* fields (GeoIPFilteringEnabled,
// GeoIPFilteringBlock, GeoIPFilteringCountries, GeoIPFilteringTrafficDirection)
// come from a pre-10 controller spec. On the primary compatibility target
// (UDM-Pro / UniFi Network 10.x) Region Blocking lives in the SEPARATE `usg_geo`
// setting: use SettingUsgGeo with GetSettingUsgGeo / UpdateSettingUsgGeo instead.
// The flat fields are retained for backward compatibility with 6.x–9.x
// controllers but are NOT persisted on 10.x, so writing them there is a no-op.
//
// Deprecated: use SettingUsgGeo (key `usg_geo`) for Region Blocking on UDM-Pro /
// UniFi Network 10.x.
const SettingUsgGeoDeprecationNote = "SettingUsg GeoIPFiltering* fields are not persisted on UDM-Pro / Network 10.x; use SettingUsgGeo (usg_geo)"
