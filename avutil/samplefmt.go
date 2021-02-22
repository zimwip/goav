package avutil

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <libavutil/samplefmt.h>
import "C"

type (
	AvSampleFormat C.enum_AVSampleFormat
)

const (
	AV_SAMPLE_FMT_NONE = int64(C.AV_SAMPLE_FMT_NONE)
	AV_SAMPLE_FMT_U8   = int64(C.AV_SAMPLE_FMT_U8)
	AV_SAMPLE_FMT_S16  = int64(C.AV_SAMPLE_FMT_S16)
	AV_SAMPLE_FMT_S32  = int64(C.AV_SAMPLE_FMT_S32)
	AV_SAMPLE_FMT_FLT  = int64(C.AV_SAMPLE_FMT_FLT)
	AV_SAMPLE_FMT_DBL  = int64(C.AV_SAMPLE_FMT_DBL)

	AV_SAMPLE_FMT_U8P  = int64(C.AV_SAMPLE_FMT_U8P)
	AV_SAMPLE_FMT_S16P = int64(C.AV_SAMPLE_FMT_S16P)
	AV_SAMPLE_FMT_S32P = int64(C.AV_SAMPLE_FMT_S32P)
	AV_SAMPLE_FMT_FLTP = int64(C.AV_SAMPLE_FMT_FLTP)
	AV_SAMPLE_FMT_DBLP = int64(C.AV_SAMPLE_FMT_DBLP)
	AV_SAMPLE_FMT_S64  = int64(C.AV_SAMPLE_FMT_S64)
	AV_SAMPLE_FMT_S64P = int64(C.AV_SAMPLE_FMT_S64P)

	AV_SAMPLE_FMT_NB = AvSampleFormat(C.AV_SAMPLE_FMT_NB)
)

/**
 * Return the name of sample_fmt, or NULL if sample_fmt is not
 * recognized.
 */
func (fmt AvSampleFormat) AvGetSampleFmtName() string {
	return C.GoString(C.av_get_sample_fmt_name((C.enum_AVSampleFormat)(fmt)))
}

/**
 * Allocate a data pointers array, samples buffer for nb_samples
 * samples, and fill data pointers and linesize accordingly.
 *
 * This is the same as av_samples_alloc(), but also allocates the data
 * pointers array.
 *
 * @see av_samples_alloc()
 */
func AvSamplesAllocArrayAndSamples() {

}

//int av_samples_alloc_array_and_samples(uint8_t ***audio_data, int *linesize, int nb_channels,
//int nb_samples, enum AVSampleFormat sample_fmt, int align);
