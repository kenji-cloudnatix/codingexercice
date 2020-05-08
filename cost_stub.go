package cost

type instance struct {
	id string
	// state is one of the "running", "terminated", or "pending".
	state string
}

type clusterSnapshot struct {
	instances []*instance
	// timestamp is the time when the snapshot was taken.
	timestamp int64
}

func estimateClusterCost(snapshots []*clusterSnapshot, startTime, endTime int64, pricePerTimestamp float64) float64 {
	// TODO: Implement this.
	return 0.0
}
