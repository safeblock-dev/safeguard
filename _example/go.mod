module example

go 1.22.0

replace github.com/safeblock-dev/safeguard => ./..

require (
	github.com/matchsystems/werr v0.1.3
	github.com/safeblock-dev/safeguard v0.0.2
)

require github.com/safeblock-dev/werr v0.0.7 // indirect