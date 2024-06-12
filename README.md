# gohlslib

[![Test](https://nexus_hls/workflows/test/badge.svg)](https://nexus_hls/actions?query=workflow:test)
[![Lint](https://nexus_hls/workflows/lint/badge.svg)](https://nexus_hls/actions?query=workflow:lint)
[![Go Report Card](https://goreportcard.com/badge/nexus_hls)](https://goreportcard.com/report/nexus_hls)
[![CodeCov](https://codecov.io/gh/bluenviron/gohlslib/branch/main/graph/badge.svg)](https://app.codecov.io/gh/bluenviron/gohlslib/branch/main)
[![PkgGoDev](https://pkg.go.dev/badge/nexus_hls)](https://pkg.go.dev/nexus_hls#pkg-index)

HLS client and muxer library for the Go programming language, written for [MediaMTX](https://github.com/bluenviron/mediamtx).

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

* [Examples](#examples)
* [API Documentation](#api-documentation)
* [Specifications](#specifications)
* [Related projects](#related-projects)

## Examples

* [playlist-parser](examples/playlist-parser/main.go)
* [client](examples/client/main.go)
* [client-absolute-timestamp](examples/client-absolute-timestamp/main.go)
* [client-codec-h264-save-to-disk](examples/client-codec-h264-save-to-disk/main.go)
* [client-codec-h264-convert-to-jpeg](examples/client-codec-h264-convert-to-jpeg/main.go)
* [client-codec-mpeg4audio-save-to-disk](examples/client-codec-mpeg4audio-save-to-disk/main.go)
* [muxer](examples/muxer/main.go)

## API Documentation

[Click to open the API Documentation](https://pkg.go.dev/nexus_hls#pkg-index)

## Specifications

|name|area|
|----|----|
|[RFC2616, HTTP 1.1](https://datatracker.ietf.org/doc/html/rfc2616)|protocol|
|[RFC8216, HLS](https://datatracker.ietf.org/doc/html/rfc8216)|protocol|
|[HLS v2](https://datatracker.ietf.org/doc/html/draft-pantos-hls-rfc8216bis)|protocol|
|[HTTP Live Streaming by Apple](https://developer.apple.com/documentation/http-live-streaming)|protocol|
|[Codec specifications](https://github.com/bluenviron/mediacommon#specifications)|codecs|
|[Golang project layout](https://github.com/golang-standards/project-layout)|project layout|

## Related projects

* [MediaMTX](https://github.com/bluenviron/mediamtx)
* [gortsplib](https://github.com/bluenviron/gortsplib)
* [mediacommon](https://github.com/bluenviron/mediacommon)