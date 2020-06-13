package gigasecond

// import path for the time package from the standard library
import "time"

const gigaSecond = time.Duration(1e9) * time.Second

// AddGigasecond adds 10^9 seconds to the specified time
func AddGigasecond(startTime time.Time) time.Time {
	return startTime.Add(gigaSecond)
}
