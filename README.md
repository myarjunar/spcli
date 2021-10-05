# spcli
## What is this about
A command line interface application to find shortest path and path length from a graph.
```
spcli graph.csv --sp AB --pl ABCDE
```

## Instructions
- Prepare a csv file that contains nodes connections and it's weighted value, e.g.
```
start_id,end_id,travel_time
A,B,15
B,C,4
C,D,8
D,C,10
D,E,6
A,D,50
C,E,22
E,B,13
A,E,37
```
- This command line have two flags
    - `--sp`: To find the shortest distance between two nodes
    - `--pl`: To find the total distance across the path

## Assumptions and limitations
- Error message will be thrown if there are node from the input parameter that are not exist or not connected with any other node

## How to use
### Manual
1. Clone this repository
2. From the project root directory, run below commands to setup and build the application
```
make setup
make build
```
3. After successfully run previous commands, the output binary would be available inside `bin/spcli`
4. Run the command e.g.
```
bin/spcli <graph_file> --sp <{start_node}{end_node}>
```

### Test scenarios
To run all test scenarios of different kind of input, run below command
```
make test
```
