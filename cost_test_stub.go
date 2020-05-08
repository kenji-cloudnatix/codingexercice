package cost

import "testing"

func TestInstanceAlwaysRunning(t *testing.T) {
	snapshots := []*clusterSnapshot{
		{
			timestamp: 100,
			instances: []*instance{
				{
					"id0",
					"running",
				},
			},
		},
		{
			timestamp: 400,
			instances: []*instance{
				{
					"id0",
					"running",
				},
			},
		},
	}

	startTime := int64(200)
	endTime := int64(300)
	pricePerTimestamp := 1.0
	actual := estimateClusterCost(snapshots, startTime, endTime, pricePerTimestamp)
	expected := float64(endTime-startTime) * pricePerTimestamp
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}

func TestDynamicChange(t *testing.T) {
	snapshots := []*clusterSnapshot{
		{
			timestamp: 100,
			instances: []*instance{
				{
					"id0",
					"running",
				},
				{
					"id1",
					"running",
				},
			},
		},
		{
			timestamp: 200,
			instances: []*instance{
				// id0 got got terminated between timestamp 100 and 200.
				{
					"id0",
					"terminated",
				},
				{
					"id1",
					"running",
				},
				// id2 was created between timestamp 100 and 200.
				{
					"id2",
					"running",
				},
			},
		},
		{
			timestamp: 400,
			instances: []*instance{
				// id0 started running again between timestamp 200 and 400.
				{
					"id0",
					"running",
				},
				{
					"id1",
					"running",
				},
				// "id2" was destroyed between timestamp 200 and 400.
			},
		},
	}

	startTime := int64(50)
	endTime := int64(350)
	pricePerTimestamp := 1.0

	c := estimateClusterCost(snapshots, startTime, endTime, pricePerTimestamp)
	expected := 0.0 // TODO: Fill this. The expected value depends on how we estimate the running time.
	if c != expected {
		t.Errorf("expected %v, but got %v", expected, c)
	}
}
