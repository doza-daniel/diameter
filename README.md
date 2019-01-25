# diameter
Implementation of an algorithm that given a set of points, finds two most distant ones

## Running
Make sure you have `go` (version >= 1.11) installed.
```
$ git clone https://github.com/doza-daniel/diameter.git
$ cd diameter
$ make
$ bin/pointgen -n 500 -u 1000 -l -1000 | bin/smart
```
Both 'smart' and 'naive' implementations will take JSON array of points from `stdin`. You can use pointgen program to generate
random points in a suitable format for these programs. In the above example, `pointgen` will create 500 random points, with
upper bound for `x` and `y` being set to 1000 and lower bound being set to -1000.
