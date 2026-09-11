# Cardinality Cost Model

## The multiplication

Time series count = product of each label's distinct value count.

    {method, status, route}
    = 5 methods x 6 status codes x 40 routes
    = 1,200 time series

Adding one high-cardinality label (e.g. user_id) multiplies this by the
number of active users:

    1,200 x 50,000 active users
    = 60,000,000 time series

One label addition: 5x10^4 multiplier.

## Memory estimate

Active-series memory cost varies by Prometheus version, label length,
and scrape interval, so no single number applies universally. Using an
illustrative estimate of ~2KB per active series:

    60,000,000 series x 2KB = 120GB

This assumes 50,000 active users and 2KB/series. Any number quoted from
this model should carry its assumptions rather than stand alone.

## Why this matters for detection design

This is why C001/C002 are ERROR severity: the operational cost of a
single bad label is not linear, it is multiplicative with every other
label already on the metric. A reviewer looking at a one-line diff
(`+ "user_id"`) has no way to see this multiplier without already
knowing the existing label cardinalities on that metric - which is
exactly the kind of check a CI gate can do and a code reviewer cannot.