# nexushls

[![Test](https://github.com/hvantoan/nexushls/workflows/test/badge.svg)](https://github.com/hvantoan/nexushls/actions?query=workflow:test)
[![Lint](https://github.com/hvantoan/nexushls/workflows/lint/badge.svg)](https://github.com/hvantoan/nexushls/actions?query=workflow:lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/hvantoan/nexushls)](https://goreportcard.com/report/github.com/hvantoan/nexushls)
[![CodeCov](https://codecov.io/gh/bluenviron/nexushls/branch/main/graph/badge.svg)](https://app.codecov.io/gh/bluenviron/nexushls/branch/main)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hvantoan/nexushls)](https://pkg.go.dev/github.com/hvantoan/nexushls#pkg-index)

HLS client and muxer library for the Go programming language

Go &ge; 1.20 is required.

Features:

* Client

  * Read streams in MPEG-TS, fMP4 or Low-latency format
  * Read tracks encoded with AV1, VP9, H265, H264, Opus, MPEG-4 Audio (AAC)
  * Get absolute timestamp of incoming data

* Muxer

  * Generate streams in MPEG-TS, fMP4 or Low-latency format
  * Write tracks encoded with AV1, VP9, H265, H264, Opus, MPEG-4 audio (AAC)
  * Save generated segments on disk

* General

  * Parse and produce M3U8 playlists
  * Examples

## Table of contents
* [API Documentation](#api-documentation)
* [Specifications](#specifications)
* [Related projects](#related-projects)

## API Documentation

[Click to open the API Documentation](https://pkg.go.dev/github.com/hvantoan/nexushls#pkg-index)

## Specifications

|name|area|
|----|----|
|[RFC2616, HTTP 1.1](https://datatracker.ietf.org/doc/html/rfc2616)|protocol|
|[RFC8216, HLS](https://datatracker.ietf.org/doc/html/rfc8216)|protocol|
|[HLS v2](https://datatracker.ietf.org/doc/html/draft-pantos-hls-rfc8216bis)|protocol|
|[HTTP Live Streaming by Apple](https://developer.apple.com/documentation/http-live-streaming)|protocol|
|[Codec specifications](https://github.com/bluenviron/mediacommon#specifications)|codecs|
|[Golang project layout](https://github.com/golang-standards/project-layout)|project layout|