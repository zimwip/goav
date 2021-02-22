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
func AvImageGetBufferSize(pix_fmt AvPixelFormat, width, height int, align int) int {
	return int(C.av_image_get_buffer_size((C.enum_AVPixelFormat)(pix_fmt), C.int(width), C.int(height), C.int(align)))
}

/**
 * Return the name of sample_fmt, or NULL if sample_fmt is not
 * recognized.
 */
func AvImageFillArrays(f *Frame, src *uint8, pix_fmt AvPixelFormat, width, height int, align int) int {
	cdata := (**C.uint8_t)(unsafe.Pointer(&f.data[0]))
	cdst_linesize := (*C.int)(unsafe.Pointer(&f.linesize[0]))
	csrc := (*C.uint8_t)(unsafe.Pointer(src))
	return int(C.av_image_fill_arrays(cdata, cdst_linesize, csrc, (C.enum_AVPixelFormat)(pix_fmt), C.int(width), C.int(height), C.int(align)))
}
