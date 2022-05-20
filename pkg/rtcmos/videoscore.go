package rtcmos

import (
	"math"
	"strings"
)

const (
	DefaultHeight    = 640
	DefaultWidth     = 480
	DefaultFrameRate = 30
)

// VideoConfig is used to specify the video configuration used
type VideoConfig struct {
	// Codec: video codec used - opus / vp8 / vp9 / h264
	Codec string
	// Width: Resolution of the video received
	Width *int32
	// ExpectedWidth: Resolution of the rendering widget
	ExpectedWidth *int32
	// Height: Resolution of the video received
	Height *int32
	// ExpectedHeight: Resolution of the rendering widget
	ExpectedHeight *int32
	// FrameRate: FrameRate of the video received
	FrameRate *int32
	// ExpectedFrameRate: FrameRate of the video source
	ExpectedFrameRate *int32
}

// VideoScore - MOS calculation based on logarithmic regression
func VideoScore(input Stat) Scores {

	stat := normalizeVideoStat(input)
	videoConfig := stat.VideoConfig
	if videoConfig == nil {
		return Scores{}
	}
	pixels := float64((*videoConfig.ExpectedWidth) * (*videoConfig.ExpectedHeight))
	codecFactor := 1.0
	if strings.ToLower(videoConfig.Codec) == "vp9" {
		codecFactor = 1.2
	}

	delay := float64(*stat.BufferDelay + *stat.RoundTripTime/2)

	var score Scores
	// These parameters are generated with a logarithmic regression
	// on some very limited test data for now
	// They are based on the bits per pixel per frame (bPPPF)
	if *videoConfig.FrameRate != 0 {
		frameRate := float64(*videoConfig.FrameRate)
		bPPPF := (codecFactor * float64(stat.Bitrate)) / pixels / frameRate
		base := clamp(0.56*math.Log(bPPPF)+5.36, 1, 5)
		MOS := base - 1.9*math.Log(float64(*videoConfig.ExpectedFrameRate)/frameRate) - delay*0.002
		score.VideoScore = clamp(math.Round(MOS*100)/100, 1, 5)
	} else {
		score.VideoScore = 1
	}
	return score
}

func normalizeVideoStat(input Stat) Stat {
	if input.RoundTripTime == nil {
		input.RoundTripTime = int32Ptr(DefaultRoundTripTime)
	}

	if input.BufferDelay == nil {
		input.BufferDelay = int32Ptr(DefaultBufferDelay)
	}
	videoConfig := input.VideoConfig

	if videoConfig.ExpectedHeight == nil {
		videoConfig.ExpectedHeight = int32Ptr(DefaultHeight)
	}

	if videoConfig.ExpectedWidth == nil {
		if videoConfig.Width != nil {
			videoConfig.ExpectedWidth = videoConfig.Width
		} else {
			videoConfig.ExpectedWidth = int32Ptr(DefaultWidth)
		}
	}
	if videoConfig.FrameRate == nil {
		videoConfig.FrameRate = int32Ptr(DefaultFrameRate)
	}

	if videoConfig.ExpectedFrameRate == nil {
		if videoConfig.FrameRate != nil {
			videoConfig.ExpectedFrameRate = videoConfig.FrameRate
		} else {
			videoConfig.ExpectedFrameRate = int32Ptr(DefaultFrameRate)
		}
	}

	return input
}
