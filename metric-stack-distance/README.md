# Stack Distance Study

In this investigation, we aim to find a fast and reliable way of calculating the optimal capacity for a cache service with a given workload. Specifically, we use the Stack Distances of an item in the LRU stack throughout the execution of a cache workload.

## Processing data: Stack Distance List

First, we calculate the Stack Distance for all ocurrences of an ID in our workload. This should result in a file with a sorted decrescing list of Stack Distances in this format:

```csv
9223372036854775807
9223372036854775807
9223372036854775807
9223372036854775807
100000
100000
93497
92125
90125
...
0
0
0
0
```

The number `9223372036854775807` is a representation of infinity Stack Distance. We need an indication of infinity distance because when an ID appears in a workload for the first time, there is no previous occurrence of such ID, so the distance between the LRU Stack Height of both occurrences is inexistent, or, to facilitate calculus, infinite.

## Calculating Hit Ratio

We use the **Stack Distance List** to calculate the Hit Ratio of the workload for a wide range of capacities with a simple script. 

To do so, move the **Stack Distance List** file to the `stack-distance-study/data/`. The file should be named `stack-distance-list.csv`, and the scripts assume the file has precisely `60917095` lines. This should improve in the future to accept a **Stack Distance List** with an arbitrary name and number of lines.

Finally, run the script with `parse-stack-distance-hitratio.go`:

```sh
go run parse-stack-distance-gitratio.go > data/stack-distance-hitratio.csv
```

The scripts' output is a CSV located in `data/stack-distance-hitratio.csv` with the columns `stackDist`, which can be interpreted as a cache capacity, and `hitRatio`, the expected Hit Ratio for this capacity. In this file, `-1.0` results from processing IDs with infinity Stack Distance (this should be removed or improved in the future).

## Visualization

To visualize the Hit Ratio as the capacity increases, open the R-project in this directory on RStudio. Next, open the file `stack-distance-hitratio.Rmd` and run the chunks of code.

