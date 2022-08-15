package rtcmos

import (
	"math"
)

// AudioConfig is used to specify audio configuration used
type AudioConfig struct {
	// Fec: flag to pass opus forward error correction status
	Fec *bool
	// Dtx: flag to pass opus discontinuous transmission status
	Dtx *bool
}

// AudioScore - MOS calculation based on E-Model algorithm
func AudioScore(input Stat) Scores {
	stat := normalizeAudioStat(input)
	const R0 = 100
	// Assume 20 packetization delay
	delay := float64(20 + *stat.BufferDelay + *stat.RoundTripTime/2)
	pl := float64(stat.PacketLoss)

	audioConfig := stat.AudioConfig

	// Ignore audio bitrate in dtx mode
	var Ie float64
	if *audioConfig.Dtx {
		Ie = 8
	} else {
		if stat.Bitrate > 0 {
			Ie = clamp(55-4.6*math.Log(float64(stat.Bitrate)), 0, 30)
		} else {
			Ie = 6
		}
	}

	Bpl := float64(10)
	if *audioConfig.Fec {
		Bpl = 20
	}

	Ipl := Ie + (100-Ie)*(pl/(pl+Bpl))

	delayFactor := float64(0)
	if delay > 150 {
		delayFactor = 0.1 * (delay - 150)
	}
	Id := delay*0.03 + delayFactor

	R := clamp(R0-Ipl-Id, 0, 100)
	MOS := 1 + 0.035*R + (R*(R-60)*(100-R)*7)/1000000

	return Scores{AudioScore: clamp(math.Round(MOS*100)/100, 1, 5)}
}

func normalizeAudioStat(input Stat) Stat {
	if input.RoundTripTime == nil {
		input.RoundTripTime = int32Ptr(DefaultRoundTripTime)
	}

	if input.BufferDelay == nil {
		input.BufferDelay = int32Ptr(DefaultBufferDelay)
	}

	if input.AudioConfig.Fec == nil {
		input.AudioConfig.Fec = boolPtr(true)
	}
	if input.AudioConfig.Dtx == nil {
		input.AudioConfig.Dtx = boolPtr(false)
	}

	return input
}
