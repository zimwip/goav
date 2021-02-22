// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avformat provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package avutil

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <libavutil/channel_layout.h>
import "C"

const (
	AV_CH_FRONT_LEFT            = int(C.AV_CH_FRONT_LEFT)
	AV_CH_FRONT_RIGHT           = int(C.AV_CH_FRONT_RIGHT)
	AV_CH_FRONT_CENTER          = int(C.AV_CH_FRONT_CENTER)
	AV_CH_LOW_FREQUENCY         = int(C.AV_CH_LOW_FREQUENCY)
	AV_CH_BACK_LEFT             = int(C.AV_CH_BACK_LEFT)
	AV_CH_BACK_RIGHT            = int(C.AV_CH_BACK_RIGHT)
	AV_CH_FRONT_LEFT_OF_CENTER  = int(C.AV_CH_FRONT_LEFT_OF_CENTER)
	AV_CH_FRONT_RIGHT_OF_CENTER = int(C.AV_CH_FRONT_RIGHT_OF_CENTER)
	AV_CH_BACK_CENTER           = int(C.AV_CH_BACK_CENTER)
	AV_Ch_SIDE_LEFT             = int(C.AV_CH_SIDE_LEFT)
	AV_CH_SIDE_RIGHT            = int(C.AV_CH_SIDE_RIGHT)
	AV_CH_TOP_CENTER            = int(C.AV_CH_TOP_CENTER)
	AV_CH_TOP_FRONT_LEFT        = int(C.AV_CH_TOP_FRONT_LEFT)
	AV_CH_TOP_FRONT_CENTER      = int(C.AV_CH_TOP_FRONT_CENTER)
	AV_CH_TOP_FRONT_RIGHT       = int(C.AV_CH_TOP_FRONT_RIGHT)
	AV_CH_TOP_BACK_LEFT         = int(C.AV_CH_TOP_BACK_LEFT)
	AV_CH_TOP_BACK_CENTER       = int(C.AV_CH_TOP_BACK_CENTER)
	AV_CH_TOP_BACK_RIGHT        = int(C.AV_CH_TOP_BACK_RIGHT)
	AV_CH_STEREO_LEFT           = int(C.AV_CH_STEREO_LEFT)
	AV_CH_STEREO_RIGHT          = int(C.AV_CH_STEREO_RIGHT)
	AV_CH_WIDE_LEFT             = int(C.AV_CH_WIDE_LEFT)
	AV_CH_WIDE_RIGHT            = int(C.AV_CH_WIDE_RIGHT)
	AV_CH_SURROUND_DIRECT_LEFT  = int(C.AV_CH_SURROUND_DIRECT_LEFT)
	AV_CH_SURROUND_DIRECT_RIGHT = int(C.AV_CH_SURROUND_DIRECT_RIGHT)
	AV_CH_LOW_FREQUENCY_2       = int(C.AV_CH_LOW_FREQUENCY_2)
	//AV_CH_TOP_SIDE_LEFT         = int(C.AV_CH_TOP_SIDE_LEFT)
	//AV_CH_TOP_SIDE_RIGHT        = int(C.AV_CH_TOP_SIDE_RIGHT)
	//AV_CH_BOTTOM_FRONT_CENTER   = int(C.AV_CH_BOTTOM_FRONT_CENTER)
	//AV_CH_BOTTOM_FRONT_LEFT     = int(C.AV_CH_BOTTOM_FRONT_LEFT)
	//AV_CH_BOTTOM_FRONT_RIGHT    = int(C.AV_CH_BOTTOM_FRONT_RIGHT)

	/** Channel mask value used for AVCodecContext.request_channel_layout
	  to indicate that the user requests the channel order of the decoder output
	  to be the native codec channel order. */
	AV_CH_LAYOUT_NATIVE = uint64(C.AV_CH_LAYOUT_NATIVE)

	/**
	 * @}
	 * @defgroup channel_mask_c Audio channel layouts
	 * @{
	 * */
	AV_CH_LAYOUT_MONO              = int(C.AV_CH_LAYOUT_MONO)
	AV_CH_LAYOUT_STEREO            = int(C.AV_CH_LAYOUT_STEREO)
	AV_CH_LAYOUT_2POINT1           = int(C.AV_CH_LAYOUT_2POINT1)
	AV_CH_LAYOUT_2_1               = int(C.AV_CH_LAYOUT_2_1)
	AV_CH_LAYOUT_SURROUND          = int(C.AV_CH_LAYOUT_SURROUND)
	AV_CH_LAYOUT_3POINT1           = int(C.AV_CH_LAYOUT_3POINT1)
	AV_CH_LAYOUT_4POINT0           = int(C.AV_CH_LAYOUT_4POINT0)
	AV_CH_LAYOUT_4POINT1           = int(C.AV_CH_LAYOUT_4POINT1)
	AV_CH_LAYOUT_282               = int(C.AV_CH_LAYOUT_2_2)
	AV_CH_LAYOUT_QUAD              = int(C.AV_CH_LAYOUT_QUAD)
	AV_CH_LAYOUT_5POINT0           = int(C.AV_CH_LAYOUT_5POINT0)
	AV_CH_LAYOUT_5POINT1           = int(C.AV_CH_LAYOUT_5POINT1)
	AV_CH_LAYOUT_5POINT0_BACK      = int(C.AV_CH_LAYOUT_5POINT0_BACK)
	AV_CH_LAYOUT_5POINT1_BACK      = int(C.AV_CH_LAYOUT_5POINT1_BACK)
	AV_CH_LAYOUT_6POINT0           = int(C.AV_CH_LAYOUT_6POINT0)
	AV_CH_LAYOUT_6POINT0_FRONT     = int(C.AV_CH_LAYOUT_6POINT0_FRONT)
	AV_CH_LAYOUT_HEXAGONAL         = int(C.AV_CH_LAYOUT_HEXAGONAL)
	AV_CH_LAYOUT_6POINT1           = int(C.AV_CH_LAYOUT_6POINT1)
	AV_CH_LAYOUT_6POINT1_BACK      = int(C.AV_CH_LAYOUT_6POINT1_BACK)
	AV_CH_LAYOUT_6POINT1_FRONT     = int(C.AV_CH_LAYOUT_6POINT1_FRONT)
	AV_CH_LAYOUT_7POINT0           = int(C.AV_CH_LAYOUT_7POINT0)
	AV_CH_LAYOUT_7POINT0_FRONT     = int(C.AV_CH_LAYOUT_7POINT0_FRONT)
	AV_CH_LAYOUT_7POINT1           = int(C.AV_CH_LAYOUT_7POINT1)
	AV_CH_LAYOUT_7POINT1_WIDE      = int(C.AV_CH_LAYOUT_7POINT1_WIDE)
	AV_CH_LAYOUT_7POINT1_WIDE_BACK = int(C.AV_CH_LAYOUT_7POINT1_WIDE_BACK)
	AV_CH_LAYOUT_OCTOGONAL         = int(C.AV_CH_LAYOUT_OCTAGONAL)
	AV_CH_LAYOUT_HEXADECAGONAL     = int(C.AV_CH_LAYOUT_HEXADECAGONAL)
	AV_CH_LAYOUT_STEREO_DOWWNMIX   = int(C.AV_CH_LAYOUT_STEREO_DOWNMIX)
	//AV_CH_LAYOUT_22POINT2          = int(C.AV_CH_LAYOUT_22POINT2)
)
