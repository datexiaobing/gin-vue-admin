package ffmpegBinding

import (
	"fmt"
)

type Options struct {
	Parame []string
}



//自定义绑定参数
func (e *Options)WithCustomParame(parame string,value interface{},c bool)  {

	if !c {
		for i,v := range e.Parame{
			if v == parame{
				e.Parame =append(e.Parame[:i], e.Parame[i+1:]...)
				if _, ok := value.(string); ok && value!="" {
					e.Parame =append(e.Parame[:i], e.Parame[i+1:]...)
				}
				break
			}
		}
	}


	if _, ok := value.(bool); ok {
		e.Parame = append(e.Parame, parame)
	}

	if sv, ok := value.(string); ok {
		e.Parame = append(e.Parame,parame,  sv)
	}


	if va, ok := value.([]string); ok {
		e.Parame = append(e.Parame,parame)
		for i := 0; i < len(va); i++ {
			item := va[i]
			e.Parame = append(e.Parame,item)
		}
	}


	if vm, ok := value.(map[string]string); ok {
		for k, v := range vm {
			e.Parame = append(e.Parame, parame, fmt.Sprintf("%v:%v", k, v))
		}
	}

	if vi, ok := value.(*int); ok {
		e.Parame = append(e.Parame, parame, fmt.Sprintf("%d", *vi))
	}

}


func (e *Options)Aspect(value string)  {
	e.WithCustomParame("-aspect", value,false)
}
func (e *Options)Resolution(value string)  {
	e.WithCustomParame("-s", value,false)
}
func (e *Options)VideoBitRate(value string)  {
	e.WithCustomParame("-b:v", value,false)
}
func (e *Options)VideoBitRateTolerance(value int)  {
	e.WithCustomParame("-bt", value,false)
}
func (e *Options)VideoMaxBitRate(value int)  {
	e.WithCustomParame("-maxrate", value,false)
}
func (e *Options)VideoMinBitrate(value int)  {
	e.WithCustomParame("-minrate", value,false)
}
func (e *Options)VideoCodec(value string)  {
	e.WithCustomParame("-c:v", value,false)
}
func (e *Options)Vframes(value int)  {
	e.WithCustomParame("-vframes", value,false)
}
func (e *Options)FrameRate(value int)  {
	e.WithCustomParame("-r", value,false)
}
func (e *Options)AudioRate(value int)  {
	e.WithCustomParame("-ar", value,false)
}
func (e *Options)KeyframeInterval(value int)  {
	e.WithCustomParame("-g", value,false)
}
func (e *Options)AudioCodec(value string)  {
	e.WithCustomParame("-c:a", value,false)
}
func (e *Options)AudioBitrate(value string)  {
	e.WithCustomParame("-ab", value,false)
}
func (e *Options)AudioChannels(value int)  {
	e.WithCustomParame("-ac", value,false)
}
func (e *Options)AudioVariableBitrate(value bool)  {
	e.WithCustomParame("-q:a", value,false)
}
func (e *Options)BufferSize(value int)  {
	e.WithCustomParame("-bufsize", value,false)
}
func (e *Options)Threadset(value bool)  {
	e.WithCustomParame("-threads", value,false)
}
func (e *Options)Threads(value int)  {
	e.WithCustomParame("-threads", value,false)
}
func (e *Options)Preset(value string)  {
	e.WithCustomParame("-preset", value,false)
}
func (e *Options)Tune(value string)  {
	e.WithCustomParame("-tune", value,false)
}
func (e *Options)AudioProfile(value string)  {
	e.WithCustomParame("-profile:a", value,false)
}
func (e *Options)VideoProfile(value string)  {
	e.WithCustomParame("-profile:v", value,false)
}
func (e *Options)Target(value string)  {
	e.WithCustomParame("-target", value,false)
}
func (e *Options)Duration(value string)  {
	e.WithCustomParame("-t", value,false)
}
func (e *Options)Qscale(value uint32)  {
	e.WithCustomParame("-qscale", value,false)
}
func (e *Options)Crf(value uint32)  {
	e.WithCustomParame("-crf", value,false)
}
func (e *Options)Strict(value int)  {
	e.WithCustomParame("-strict", value,false)
}
func (e *Options)MuxDelay(value string)  {
	e.WithCustomParame("-muxdelay", value,false)
}
func (e *Options)SeekTime(value string)  {
	e.WithCustomParame("-ss", value,false)
}
func (e *Options)SeekUsingTimestamp(value bool)  {
	e.WithCustomParame("-seek_timestamp", value,false)
}
func (e *Options)MovFlags(value string)  {
	e.WithCustomParame("-movflags", value,false)
}
func (e *Options)HideBanner(value bool)  {
	e.WithCustomParame("-hide_banner", value,false)
}
func (e *Options)OutputFormat(value string)  {
	e.WithCustomParame("-f", value,false)
}
func (e *Options)CopyTs(value bool)  {
	e.WithCustomParame("-copyts", value,false)
}
func (e *Options)NativeFramerateInput(value bool)  {
	e.WithCustomParame("-re", value,false)
}
func (e *Options)InputInitialOffset(value string)  {
	e.WithCustomParame("-itsoffset", value,false)
}
func (e *Options)RtmpLive(value string)  {
	e.WithCustomParame("-rtmp_live", value,false)
}
func (e *Options)HlsPlaylistType(value string)  {
	e.WithCustomParame("-hls_playlist_type", value,false)
}
func (e *Options)HlsListSize(value int)  {
	e.WithCustomParame("-hls_list_size", value,false)
}
func (e *Options)HlsSegmentDuration(value int)  {
	e.WithCustomParame("-hls_time", value,false)
}
func (e *Options)HlsMasterPlaylistName(value string)  {
	e.WithCustomParame("-master_pl_name", value,false)
}
func (e *Options)HlsSegmentFilename(value string)  {
	e.WithCustomParame("-hls_segment_filename", value,false)
}
func (e *Options)HTTPMethod(value string)  {
	e.WithCustomParame("-method", value,false)
}
func (e *Options)HTTPKeepAlive(value bool)  {
	e.WithCustomParame("-multiple_requests", value,false)
}
func (e *Options)Hwaccel(value string)  {
	e.WithCustomParame("-hwaccel", value,false)
}
func (e *Options)StreamIds(value map[string]string)  {
	e.WithCustomParame("-streamid", value,false)
}
func (e *Options)VideoFilter(value string)  {
	e.WithCustomParame("-vf", value,false)
}
func (e *Options)AudioFilter(value string)  {
	e.WithCustomParame("-af", value,false)
}
func (e *Options)SkipVideo(value bool)  {
	e.WithCustomParame("-vn", value,false)
}
func (e *Options)SkipAudio(value bool)  {
	e.WithCustomParame("-an", value,false)
}
func (e *Options)CompressionLevel(value int)  {
	e.WithCustomParame("-compression_level", value,false)
}
func (e *Options)MapMetadata(value string)  {
	e.WithCustomParame("-map_metadata", value,false)
}
func (e *Options)Metadata(value map[string]string)  {
	e.WithCustomParame("-metadata", value,false)
}
func (e *Options)EncryptionKey(value string)  {
	e.WithCustomParame("-hls_key_info_file", value,false)
}
func (e *Options)Bframe(value int)  {
	e.WithCustomParame("-bf", value,false)
}
func (e *Options)PixFmt(value string)  {
	e.WithCustomParame("-pix_fmt", value,false)
}
func (e *Options)WhiteListProtocols(value []string)  {
	e.WithCustomParame("-protocol_whitelist", value,false)
}
func (e *Options)Overwrite(value bool)  {
	e.WithCustomParame("-y", value,false)
}
func (e *Options)CodecV(value string)  {
	e.WithCustomParame("-codec:v", value,false)
}
func (e *Options)CodecA(value string)  {
	e.WithCustomParame("-codec:a", value,false)
}
func (e *Options)Filter(value string)  {
	e.WithCustomParame("-filter:v", value,false)
}
func (e *Options)Lavfi(value string)  {
	e.WithCustomParame("-lavfi", value,false)
}
func (e *Options)I(value string)  {
	e.WithCustomParame("-i", value,true)
}
func (e *Options)ForceKeyFrames(value string)  {
	e.WithCustomParame("-force_key_frames", value,false)
}
func (e *Options)ExtraArgs(value map[string]interface {})  {
	e.WithCustomParame("", value,true)
}

