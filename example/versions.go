package main

import (
	"log"

	"github.com/zimwip/goav/avcodec"
	"github.com/zimwip/goav/avdevice"
	"github.com/zimwip/goav/avfilter"
	"github.com/zimwip/goav/avformat"
	"github.com/zimwip/goav/avutil"
	"github.com/zimwip/goav/swresample"
	"github.com/zimwip/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
