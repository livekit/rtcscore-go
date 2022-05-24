package rtcmos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	Jitter = 0
	Rtt    = 0
)

func TestSpatialLayer(t *testing.T) {
	{
		stat := Stat{
			Bitrate: 620000,
			VideoConfig: &VideoConfig{Width: int32Ptr(640), Height: int32Ptr(360), FrameRate: int32Ptr(15),
				ExpectedWidth: int32Ptr(1280), ExpectedHeight: int32Ptr(720)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("spatial high-mid", scores[0].VideoScore)
		//require.GreaterOrEqual(t, scores[0].VideoScore, 4.9)
	}
	{
		stat := Stat{
			Bitrate: 310000,
			VideoConfig: &VideoConfig{Width: int32Ptr(320), Height: int32Ptr(180), FrameRate: int32Ptr(15),
				ExpectedWidth: int32Ptr(640), ExpectedHeight: int32Ptr(360)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("spatial mid-low", scores[0].VideoScore)
		//require.GreaterOrEqual(t, scores[0].VideoScore, 4.9)
	}
	{
		stat := Stat{
			Bitrate: 310000,
			VideoConfig: &VideoConfig{Width: int32Ptr(320), Height: int32Ptr(180), FrameRate: int32Ptr(15),
				ExpectedWidth: int32Ptr(1280), ExpectedHeight: int32Ptr(720)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}

		scores := Score([]Stat{stat})
		t.Log("spatial high-low", scores[0].VideoScore)
		//require.GreaterOrEqual(t, scores[0].VideoScore, 4.9)
	}

}

func TestIdealScores(t *testing.T) {
	{
		stat := Stat{
			Bitrate:       2100000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(1280), Height: int32Ptr(720), FrameRate: int32Ptr(30)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("720x1280x30fpsx2.1Mbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.9)
	}
	{
		stat := Stat{
			Bitrate:       2000000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(1280), Height: int32Ptr(720), FrameRate: int32Ptr(30)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("720x1280x30fpsx2Mbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.5)
	}
	{
		stat := Stat{
			Bitrate:       1500000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(1280), Height: int32Ptr(720), FrameRate: int32Ptr(30)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("720x1280x30fpsx1.5Mbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.0)
	}
}

func TestMediumLayer(t *testing.T) {
	{
		stat := Stat{
			Bitrate:       620000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(640), Height: int32Ptr(360), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("640x360x15fpsx620Kbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.9)
	}
	{
		stat := Stat{
			Bitrate:       550000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(640), Height: int32Ptr(360), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("640x360x15fpsx550Kbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.5)
	}
	{
		stat := Stat{
			Bitrate:       450000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(640), Height: int32Ptr(360), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("640x360x15fpsx450Kbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.0)
	}
	{
		stat := Stat{
			Bitrate:       350000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(640), Height: int32Ptr(360), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("640x360x15fpsx350Kbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 3.5)
	}
	{
		stat := Stat{
			Bitrate:       250000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(640), Height: int32Ptr(360), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("640x360x15fpsx", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 2.5)
	}
}

func TestLowLayer(t *testing.T) {
	{
		stat := Stat{
			Bitrate:       310000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(320), Height: int32Ptr(180), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("380x180x15fpsx310Kbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.9)
	}
	{
		stat := Stat{
			Bitrate:       260000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(320), Height: int32Ptr(180), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("380x180x15fpsx260Kbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.5)
	}
	{
		stat := Stat{
			Bitrate:       210000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(320), Height: int32Ptr(180), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("380x180x15fpsx210Kbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.0)
	}
	{
		stat := Stat{
			Bitrate:       170000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(320), Height: int32Ptr(180), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("380x180x15fpsx170Kbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 3.5)
	}
	{
		stat := Stat{
			Bitrate:       110000,
			VideoConfig:   &VideoConfig{Width: int32Ptr(320), Height: int32Ptr(180), FrameRate: int32Ptr(15)},
			BufferDelay:   int32Ptr(Jitter),
			RoundTripTime: int32Ptr(Rtt),
		}
		scores := Score([]Stat{stat})
		t.Log("380x180x15fpsx110Kbps", scores[0].VideoScore)
		require.GreaterOrEqual(t, scores[0].VideoScore, 2.5)
	}
}

func TestScore(t *testing.T) {
	{
		// score of audio is close to 4.5 in perfect conditions
		stat := Stat{
			PacketLoss:    0,
			Bitrate:       0,
			RoundTripTime: int32Ptr(0),
			BufferDelay:   int32Ptr(0),
			AudioConfig:   &AudioConfig{},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].AudioScore, 4.4)
		require.LessOrEqual(t, scores[0].AudioScore, 4.5)
	}
	{
		// score of audio is 1 in worst conditions
		stat := Stat{
			PacketLoss:  100,
			AudioConfig: &AudioConfig{},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].AudioScore, 1.0)
		// TODO: 1.1
		require.LessOrEqual(t, scores[0].AudioScore, 1.2)
	}
	{
		// score of audio is 1 with huge delay
		stat := Stat{
			PacketLoss:    100,
			RoundTripTime: int32Ptr(1000000000),
			AudioConfig:   &AudioConfig{},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].AudioScore, 1.0)
		require.LessOrEqual(t, scores[0].AudioScore, 1.1)
	}
	{
		// score of audio depends on packet loss
		stat1 := Stat{
			PacketLoss:  10,
			AudioConfig: &AudioConfig{},
		}
		stat2 := Stat{
			PacketLoss:  20,
			AudioConfig: &AudioConfig{},
		}

		scores := Score([]Stat{stat1, stat2})
		require.Len(t, scores, 2)
		require.Greater(t, scores[0].AudioScore, scores[1].AudioScore)
	}
	{
		// score of audio depends on bitrate
		stat1 := Stat{
			Bitrate:     100000,
			AudioConfig: &AudioConfig{},
		}

		stat2 := Stat{
			PacketLoss:  50000,
			AudioConfig: &AudioConfig{},
		}

		scores := Score([]Stat{stat1, stat2})
		require.Len(t, scores, 2)
		require.Greater(t, scores[0].AudioScore, scores[1].AudioScore)
	}

	{
		// score of audio depends on fec
		stat1 := Stat{
			PacketLoss:  10,
			AudioConfig: &AudioConfig{Fec: boolPtr(true)},
		}

		stat2 := Stat{
			PacketLoss:  10,
			AudioConfig: &AudioConfig{Fec: boolPtr(false)},
		}

		scores := Score([]Stat{stat1, stat2})
		require.Len(t, scores, 2)
		require.Greater(t, scores[0].AudioScore, scores[1].AudioScore)
	}

	{
		//score of audio depends on buffer delay
		stat1 := Stat{
			BufferDelay: int32Ptr(10),
			AudioConfig: &AudioConfig{},
		}
		stat2 := Stat{
			BufferDelay: int32Ptr(100),
			AudioConfig: &AudioConfig{},
		}

		scores := Score([]Stat{stat1, stat2})
		require.Len(t, scores, 2)
		require.Greater(t, scores[0].AudioScore, scores[1].AudioScore)
	}
	{
		//score of audio is average on control conditions one
		stat := Stat{
			PacketLoss:  15,
			AudioConfig: &AudioConfig{},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].AudioScore, 2.5)
		require.LessOrEqual(t, scores[0].AudioScore, 3.0)
	}
	{
		// score of audio is average on control conditions two
		stat := Stat{
			PacketLoss:  30,
			AudioConfig: &AudioConfig{},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].AudioScore, 1.5)
		require.LessOrEqual(t, scores[0].AudioScore, 2.0)
	}
	{
		// score of audio is average on control conditions three
		stat := Stat{
			PacketLoss:  50,
			AudioConfig: &AudioConfig{},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		// TODO: 1.5
		require.GreaterOrEqual(t, scores[0].AudioScore, 1.3)
		require.LessOrEqual(t, scores[0].AudioScore, 2.0)
	}
	{
		// score of video is 4.5 in perfect conditions
		stat := Stat{
			Bitrate:     13000000,
			VideoConfig: &VideoConfig{Width: int32Ptr(1280), Height: int32Ptr(720), FrameRate: int32Ptr(30)},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].VideoScore, 4.8)
		require.LessOrEqual(t, scores[0].VideoScore, 5.0)
	}
	{
		// score of video is 1 in worst bitrate conditions
		stat := Stat{
			Bitrate:     1000,
			VideoConfig: &VideoConfig{},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].VideoScore, 1.0)
		require.LessOrEqual(t, scores[0].VideoScore, 1.1)
	}
	{
		// score of video is 1 in worst framerate conditions
		stat := Stat{
			Bitrate:     10000000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(1), ExpectedFrameRate: int32Ptr(30)},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.Equal(t, scores[0].VideoScore, 1.0)
	}

	{
		// score of video is 1 if no framerate is received
		stat := Stat{
			Bitrate:     100000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(0)},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.Equal(t, scores[0].VideoScore, 1.0)
	}

	{
		// score is average on average bitrate conditions
		stat := Stat{
			Bitrate:     400000,
			VideoConfig: &VideoConfig{},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].VideoScore, 3.0)
		require.LessOrEqual(t, scores[0].VideoScore, 4.0)
	}
	{
		// score is not good on low bitrate conditions
		stat := Stat{
			Bitrate:     200000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(25)},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].VideoScore, 1.5)
		require.LessOrEqual(t, scores[0].VideoScore, 2.5)
	}
	{
		// score is not good on average bitrate conditions but low framerate
		stat := Stat{
			Bitrate:     500000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(8), ExpectedFrameRate: int32Ptr(25)},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].VideoScore, 2.0)
		require.LessOrEqual(t, scores[0].VideoScore, 3.0)
	}
	{
		// score is average on average framerate conditions
		stat := Stat{
			Bitrate:     400000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(25), ExpectedFrameRate: int32Ptr(30)},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].VideoScore, 3.0)
		require.LessOrEqual(t, scores[0].VideoScore, 4.0)
	}
	{
		// score is average on control conditions one
		stat := Stat{
			Bitrate:     300000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(20), Width: int32Ptr(640), Height: int32Ptr(480)},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].VideoScore, 3.0)
		require.LessOrEqual(t, scores[0].VideoScore, 4.0)
	}
	{
		// score is average on control conditions two
		stat := Stat{
			Bitrate:     400000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(20), Width: int32Ptr(640), Height: int32Ptr(480)},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].VideoScore, 3.5)
		require.LessOrEqual(t, scores[0].VideoScore, 4.5)
	}
	{
		// score of video depends on bitrate
		stat1 := Stat{
			Bitrate:     200000,
			VideoConfig: &VideoConfig{},
		}
		stat2 := Stat{
			Bitrate:     100000,
			VideoConfig: &VideoConfig{},
		}
		scores := Score([]Stat{stat1, stat2})
		require.Len(t, scores, 2)
		require.Greater(t, scores[0].VideoScore, scores[1].VideoScore)
	}
	{
		// score of video depends on codec
		stat1 := Stat{
			Bitrate:     200000,
			VideoConfig: &VideoConfig{Codec: "vp9"},
		}
		stat2 := Stat{
			Bitrate:     200000,
			VideoConfig: &VideoConfig{Codec: "vp8"},
		}
		scores := Score([]Stat{stat1, stat2})
		require.Len(t, scores, 2)
		require.Greater(t, scores[0].VideoScore, scores[1].VideoScore)
	}
	{
		// score of video depends on framerate
		stat1 := Stat{
			Bitrate:     200000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(15), ExpectedFrameRate: int32Ptr(15)},
		}
		stat2 := Stat{
			Bitrate:     200000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(15), ExpectedFrameRate: int32Ptr(30)},
		}
		scores := Score([]Stat{stat1, stat2})
		require.Len(t, scores, 2)
		require.Greater(t, scores[0].VideoScore, scores[1].VideoScore)
	}
	{
		// score of video depends on resolution
		stat1 := Stat{
			Bitrate:     200000,
			VideoConfig: &VideoConfig{Width: int32Ptr(100), Height: int32Ptr(100), ExpectedWidth: int32Ptr(100), ExpectedHeight: int32Ptr(100)},
		}
		stat2 := Stat{
			Bitrate:     200000,
			VideoConfig: &VideoConfig{Width: int32Ptr(640), Height: int32Ptr(480), ExpectedWidth: int32Ptr(640), ExpectedHeight: int32Ptr(480)},
		}
		scores := Score([]Stat{stat1, stat2})
		require.Len(t, scores, 2)
		require.Greater(t, scores[0].VideoScore, scores[1].VideoScore)
	}
	{
		// score of video is 1 for 0 framerate
		stat := Stat{
			Bitrate:     200000,
			VideoConfig: &VideoConfig{FrameRate: int32Ptr(0), ExpectedFrameRate: int32Ptr(0)},
		}
		scores := Score([]Stat{stat})
		require.Len(t, scores, 1)
		require.GreaterOrEqual(t, scores[0].VideoScore, 1.0)
		require.LessOrEqual(t, scores[0].VideoScore, 1.1)
	}

}
