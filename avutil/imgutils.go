package avutil

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <libavutil/imgutils.h>
import "C"
import "unsafe"

type (
//AvSampleFormat C.enum_AVSampleFormat
)

const (
//AV_SAMPLE_FMT_NONE = int64(C.AV_SAMPLE_FMT_NONE)
)

/**
 * Return the name of sample_fmt, or NULL if sample_fmt is not
 * recognized.
 */
func AvImageFillArrays(f *Frame, srcData [8]*uint8, dst_linesize [8]int32, src *uint8, pix_fmt AvPixelFormat, width, height int, align int) int {
	//cdata := (**C.uint8_t)(unsafe.Pointer(&srcData[0]))
	//cdst_linesize := (*C.int)(unsafe.Pointer(&dst_linesize[0]))
	cdata := (**C.uint8_t)(unsafe.Pointer(&f.data[0]))
	cdst_linesize := (*C.int)(unsafe.Pointer(&f.linesize[0]))
	csrc := (*C.uint8_t)(unsafe.Pointer(src))
	return int(C.av_image_fill_arrays(cdata, cdst_linesize, csrc, (C.enum_AVPixelFormat)(pix_fmt), C.int(width), C.int(height), C.int(align)))
}
