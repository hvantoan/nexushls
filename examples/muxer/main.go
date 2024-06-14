package main

import (
	_ "embed"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/bluenviron/mediacommon/pkg/formats/mpegts"

	"github.com/hvantoan/nexushls"
	"github.com/hvantoan/nexushls/pkg/codecs"
)

// This example shows how to:
// 1. generate a MPEG-TS/H264 stream with GStreamer
// 2. re-encode the stream into HLS and serve it with an HTTP server

//go:embed index.html
var index []byte

// handleIndex wraps an HTTP handler and serves the home page
func handleIndex(wrapped http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(index))
			return
		}

		wrapped(w, r)
	}
}

func main() {
	// create the HLS muxer
	mux := &nexushls.Muxer{
		VideoTrack: &nexushls.Track{
			Codec: &codecs.H264{},
		},
	}
	err := mux.Start()
	if err != nil {
		panic(err)
	}

	// create an HTTP server and link it to the HLS muxer
	s := &http.Server{
		Addr:    ":8080",
		Handler: handleIndex(mux.Handle),
	}
	log.Println("HTTP server created on :8080")
	go s.ListenAndServe()

	// create a socket to receive MPEG-TS packets
	pc, err := net.ListenPacket("udp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	log.Println("Waiting for a MPEG-TS/H264 stream on UDP port 9000 - you can send one with ffmpeg :\n" +
		"ffmpeg -re -i 'D://video.mp4' -c:v libx264 -preset ultrafast -b:v 3000k -x264-params keyint=60 " +
		"-profile:v high -f mpegts udp://127.0.0.1:9000")

	// create a MPEG-TS reader
	r, err := mpegts.NewReader(mpegts.NewBufferedReader(newPacketConnReader(pc)))
	if err != nil {
		panic(err)
	}

	var timeDec *mpegts.TimeDecoder

	// find the H264 track
	found := false
	for _, track := range r.Tracks() {
		if _, ok := track.Codec.(*mpegts.CodecH264); ok {
			// setup a callback that is called once a H264 access unit is received
			r.OnDataH26x(track, func(rawPTS int64, _ int64, au [][]byte) error {
				// decode the time
				if timeDec == nil {
					timeDec = mpegts.NewTimeDecoder(rawPTS)
				}
				pts := timeDec.Decode(rawPTS)

				// pass the access unit to the HLS muxer
				log.Printf("visit http://localhost:8080 - encoding access unit with PTS = %v", pts)
				mux.WriteH264(time.Now(), pts, au)

				return nil
			})
			found = true
			break
		}
	}

	if !found {
		panic("H264 track not found")
	}

	// read from the MPEG-TS stream
	for {
		err := r.Read()
		if err != nil {
			panic(err)
		}
	}
}
