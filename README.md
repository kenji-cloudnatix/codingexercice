# Coding Exercise: Estimate the Cluster Cost

## Introduction

The goal of this coding exercise is slightly different from
traditional whiteboard coding interview questions. It won't take a lot
of time to implement a function in question, but we would like to
focus on writing high-quality code so that we can push to production
with high confidence. Key points include:

- Handle corner cases.
- Write unit tests to cover the above corner cases.
- Make code easier to understand for other.

## Problem

Suppose that we have a cluster consisting of multiple EC2
instances. The number of instances in the cluster changes
dynamically. Instances can get terminated and restarted constantly.

We would like to implement a simple function that estimates the cost
of the cluster in a given time window. We periodically dump a snapshot
of the cluster and processes the data to estimate the cost.

The snapshot of the cluster is a slice of instances:

```go
type instance struct {
    id string
    // state is one of the "running", "terminated", or "pending".
    state string
}

type clusterSnapshot struct {
    // timestamp is the time when the snapshot was taken.
    timestamp int64
    instances []*instance
}
```

We would like to define a function that takes the following inputs and returns the cost:
- A list of `clusterSnapshot`s,
- Time window for cost estimation (start time and end time)
- Instance price

```
func estimateClusterCost(snapshots []*clusterSnapshot, startTime, endTime int64, pricePerTimestamp float64) float64
```

Please note that AWS charges money only when an instance is running.

Please also note that the time when a snapshot was taken is not always
the same time when the state transition happened. Let's consider the following case:

- A snapshot at time T0 has an instance whose state is "running".
- Another snapshot at time T1 (T1 > T0) has the same instance. Its state is "terminated".

In the above case, the state of the instance changed from "running" to
"terminated" between T0 and T1, not exactly at T1.

## Example Test Cases

Example test cases are defined in `cost_test_stub.go`. The first test
case is a basic scenario where snapshots contain only a single
instance. The instance is running during the target time window ([200,
300]). The cost is simply `(300 - 200) * pricePerTimestamp`.

The other test case is a more complex scenario where snapshots contain
there instances. The output cost will depend on how the heuristic for
estimating state transition is implemented.
