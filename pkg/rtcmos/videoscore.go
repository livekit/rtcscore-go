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

func ScaleFactor(ratio float64) float64 {
	if ratio >= 0.8 {
		return ratio
	}

	if ratio >= 0.5 {
		return 0.75
	}

	if ratio >= 0.2 {
		return 0.7
	}
	return 0.55
}

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
	actualPixels := float64(*videoConfig.Width * *videoConfig.Height)
	scaleFactor := 1.0
	if *videoConfig.ExpectedWidth != *videoConfig.Width || *videoConfig.ExpectedHeight != *videoConfig.Height {
		expectedPixels := float64(*videoConfig.ExpectedWidth * *videoConfig.ExpectedHeight)
		if expectedPixels > actualPixels {
			ratio := actualPixels / expectedPixels
			scaleFactor = ScaleFactor(ratio)
		}
	}
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
		bPPPF := (codecFactor * float64(stat.Bitrate)) / actualPixels / frameRate
		//base := clamp(2.3*math.Log(bPPPF*29)+2.3, 1, 5)
		//base := clamp(0.56*math.Log(bPPPF)+5.36, 1, 5)
		base := clamp(2.3*math.Log(bPPPF*29)+3.1, 1, 5)
		MOS := base - 1.9*math.Log(float64(*videoConfig.ExpectedFrameRate)/frameRate) - delay*0.002
		score.VideoScore = scaleFactor * clamp(math.Round(MOS*100)/100, 1, 5)
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

	if input.VideoConfig.Width == nil {
		input.VideoConfig.Width = int32Ptr(DefaultWidth)
	}

	if input.VideoConfig.Height == nil {
		input.VideoConfig.Height = int32Ptr(DefaultHeight)
	}

	if input.VideoConfig.FrameRate == nil {
		input.VideoConfig.FrameRate = int32Ptr(DefaultFrameRate)
	}

	if input.VideoConfig.ExpectedHeight == nil {
		input.VideoConfig.ExpectedHeight = input.VideoConfig.Height
	}

	if input.VideoConfig.ExpectedWidth == nil {
		input.VideoConfig.ExpectedWidth = input.VideoConfig.Width
	}

	if input.VideoConfig.ExpectedFrameRate == nil {
		input.VideoConfig.ExpectedFrameRate = input.VideoConfig.FrameRate
	}

	return input
}
