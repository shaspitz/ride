package types

// Returns the pseudo average of historical values by appending a new value
// to an average with 10% weighting. This type of heuristic average is useful
// in that you only have to store a single average to keep an accurate running average of a data stream.
// This kind of average also prioritizes the newest appended values, as opposed to a simple moving average.
// Lastly, this type of average makes new participants have a lower rating (starting at 0) until they build it up.
func ComputePseudoAverage(current, new float32) float32 {
	return 0.9*current + 0.1*new
}
