package config

type Ffmpegconfig struct {
	Ffmpeg  string `yaml:"ffmpeg"`
	Ffprobe string `yaml:"ffprobe"`
	Verbose bool   `yaml:"verbose"`
}

var Ffmpegconf = new(Ffmpegconfig)
