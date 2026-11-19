package unifi

// This will generate the *.generated.go files in this package for the specified
// client controller version. The positional argument is the Official-API spec
// version; -legacy-version pins the Internal/legacy controller resources.
// Pinned to 11.0.99 (spec snapshot taken from a console; not yet in fw-update): back to `latest` once a public 11.x >= 11.0.99 ships.
//go:generate go run ../codegen/ -version-base-dir=../codegen/ -legacy-version=9.5.21 11.0.99
