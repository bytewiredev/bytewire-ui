module github.com/bytewiredev/bytewire-ui/showcase

go 1.25.0

require (
	github.com/bytewiredev/bytewire v0.0.0-20260304230457-4a0a540671ba
	github.com/bytewiredev/bytewire-ui v0.0.0
)

require (
	github.com/dunglas/httpsfv v1.1.0 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/quic-go/quic-go v0.59.0 // indirect
	github.com/quic-go/webtransport-go v0.10.0 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
)

replace (
	github.com/bytewiredev/bytewire => ../../bytewire
	github.com/bytewiredev/bytewire-ui => ../
)
