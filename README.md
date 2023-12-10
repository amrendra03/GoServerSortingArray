# Go Server Sorting Array

This project demonstrates sequential and concurrent array sorting using a Go server.

## Overview

The Go server exposes two endpoints, `/process-single` and `/process-concurrent`, for sorting arrays sequentially and concurrently, respectively.

## Packages Used

- **fmt:** Standard Go package for formatted I/O.
- **net/http:** Go package for building HTTP servers and handling HTTP requests.
- **sort:** Go package for sorting slices.
- **sync:** Go package for synchronization primitives.
- **time:** Go package for measuring time durations.
- **github.com/gin-gonic/gin:** Third-party HTTP web framework for Go.

## Running the Server

1. Clone the repository:

   ```bash
   git clone https://github.com/amrendra03/GoServerSortingArray.git
   
2. The server will be accessible at http://localhost:8080.

# Endpoints
1. **/process-single (Sequential)**
- Method: POST

- Input:

   json
   Copy code
   {
      "to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
   }
- Output:

   json
   Copy code
   {
      "sorted_arrays": [[1, 2, 3], [4, 5, 6], [7, 8, 9]],
      "time_ns": 123456789
   }
2. **/process-concurrent (Concurrent)**
- Method: POST

- Input:

   json
   Copy code
   {
      "to_sort": [[3, 2, 1], [6, 5, 4], [9, 8, 7]]
   }
- Output:

   json
   Copy code
   {
      "sorted_arrays": [[1, 2, 3], [4, 5, 6], [7, 8, 9]],
      "time_ns": 987654321
   }
## Number of Routes
   The server has two routes:

1. /process-single (Method: POST)
2. /process-concurrent (Method: POST)