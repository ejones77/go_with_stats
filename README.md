# go_with_stats
Using Go to run linear regression on the Anscombe Quartet

Submitted assignment for Northwestern MSDS 431

## Overview

This Go program performs a single linear regression on a series of coordinates that make up the Anscombe Quartet. 

Go doesn't provide a full out-of-the-box summary compared to Python or R. 

To achieve a printed summary, this program creates a new struct for the interept, slope, correlation coefficient, R squared value, and standard error.

## Core Functionality

There are three functions that split the responsibilities between tasks. 

- `CalculateDiffs` returns the sum of squared differences 
- `CalculateResiduals` returns the squared residuals 
- `LmSummary` returns the calculated summary statistics as defined in the initial struct. 

To run the model for each quadrant in the quartet, each list of x and y values are defined as a slice of slices to iterate through calculations easier. 

## Performance Benchmarking 

My methodology here was to comment out the plotting tasks in the Python and R programs, this was to mirror the functionality between languages as closely as possible. 

Each program was then executed and timed 3 times. The png file in this repository displays the average time in milliseconds. Specific execution times are provided in `stats_benchmark.xlsx`

In all, Go was the fastest by a wide margin, with an average time of 2 ms, compared to 19 ms in R, and 50 ms in Python. 

## Recommendation

The stats package for Go is robust, but limited in functionality when compared to R or Python. Go did return the same results as the other two languages, so we can confirm that Go performs the linear regression calculations accurately.

In the process of this assignment I tried using the package's LinearRegression method [shown here](https://pkg.go.dev/github.com/montanaflynn/stats#LinearRegression) but got stumped trying to work around constructing a summary from it. The function returns the list of coordinates from the linear model, but it's on the developer to identify which summary statistics are important.

The main concern one would have with transitioning statistics work to Go is the lack of features when compared to R. The R language alone provides a notable performance boost compared to Python in its own right. 

The decision on which language to use comes down to whether performance or breadth of information is most important. If the task is to investigate improvements to a small set of models, R would be the best blend of features and performance. 

But if the task is to manage a large swath of models and performance is most important, then it's worth considering a transition to Go.
