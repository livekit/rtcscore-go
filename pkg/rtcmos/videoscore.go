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
	// Height: Resolution of the video received
	Height *int32
	// FrameRate: FrameRate of the video received
	FrameRate *float32
	// ExpectedFrameRate: FrameRate of the video source
	ExpectedFrameRate *float32
}

// VideoScore - MOS calculation based on logarithmic regression
func VideoScore(input Stat) Scores {
	stat := normalizeVideoStat(input)
	videoConfig := stat.VideoConfig
	if videoConfig == nil {
		return Scores{}
	}
	codecFactor := 1.0
	if strings.ToLower(videoConfig.Codec) == "vp9" {
		// assuming approximately 83% of vp8/h.264 bitrate for same quality
		codecFactor = 1.2
	}
	if strings.ToLower(videoConfig.Codec) == "av1" {
		// assuming approximately 70% of vp8/h.264 bitrate for same quality
		codecFactor = 1.43
	}

	delay := float64(*stat.BufferDelay + *stat.RoundTripTime/2)

	var score Scores
	// These parameters are generated with a logarithmic regression
	// on some very limited test data for now
	// They are based on the bits per pixel per frame (bPPPF)
	if *videoConfig.FrameRate != 0 {
		frameRate := float64(*videoConfig.FrameRate)
		pixels := float64(*videoConfig.Width * *videoConfig.Height)
		bPPPF := (codecFactor * float64(stat.Bitrate)) / pixels / frameRate

		//
		// A bit of speculation on logarithmic regression equation from https://github.com/ggarber/rtcscore
		// base := clamp(0.56*math.Log(bPPPF)+5.36, 1, 5)
		//
		// Assuming that derivation is based on Chrome (libwebrtc) simulcast default settings.
		// That would be 2.5 Mbps for 1280 x 720 (https://chromium.googlesource.com/external/webrtc/+/master/media/engine/simulcast.cc#83).
		// That piece of code does not specify frame rate.
		// But, assuming a frame rate of 30 fps, the equation above would yield a score of approximately 4.01
		// under perfect conditions, i. e. no delay or jitter and expected frame rate matching actual frame rate.
		//
		// LK clients by default use 1.7 Mbps for 720p30 for vp8/h.264.
		// That yields a score of approximately 3.8 using the above equation again under perfect conditions.
		// The perceived quality is good at that bit rate (based on user perception),
		// So, using a theshold like 3.5 MOS for declaring good quality should be fine.
		//
		base := clamp(0.56*math.Log(bPPPF)+5.36, 1, 5)

		score.VideoScore = clamp(base-1.9*math.Log(float64(*videoConfig.ExpectedFrameRate)/frameRate)-delay*0.002, 1, 5)
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
		input.VideoConfig.FrameRate = float32Ptr(DefaultFrameRate)
	}

	if input.VideoConfig.ExpectedFrameRate == nil {
		input.VideoConfig.ExpectedFrameRate = input.VideoConfig.FrameRate
	}

	return input
}
