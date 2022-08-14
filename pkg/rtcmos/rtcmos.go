package rtcmos

import (
	"log"
	"math"
)

const (
	DefaultRoundTripTime = 50
	DefaultBufferDelay   = 50
)

// Stat defines the input parameter to calculate Score
type Stat struct {
	PacketLoss    float32
	Bitrate       float32
	RoundTripTime *int32
	BufferDelay   *int32
	AudioConfig   *AudioConfig
	VideoConfig   *VideoConfig
}

// Scores contains to MOS audio and video scores
type Scores struct {
	// AudioScore: score based on modified E-model
	AudioScore float64
	// VideoScore: score based on logarithmic regression
	VideoScore float64
}

// Score compute audio and video scores for the passed stats
//
// returns audio/video scores for each input
func Score(stats []Stat) []Scores {
	var scores []Scores
	for _, stat := range stats {
		if stat.AudioConfig != nil {
			scores = append(scores, AudioScore(stat))
		} else if stat.VideoConfig != nil {
			scores = append(scores, VideoScore(stat))
		} else {
			log.Println("invalid request, no audio or video config")
			scores = append(scores, Scores{})
		}
	}
	return scores
}

func clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(value, max))
}

func int32Ptr(x int32) *int32 {
	return &x
}

func float32Ptr(x float32) *float32 {
	return &x
}

func boolPtr(x bool) *bool {
	return &x
}
