# rubarb

Each row of pixels in the input image is converted into a line. The offset at each point is controlled by the `-x` and `-y` flag values which are multiplied by each pixel's brightness. Line width and seperation can be changed with the `-l` and `-s` flags.

```
  -c	use image color (default true)
  -i string
    	input file (default "in.png")
  -l int
    	line width (default 1)
  -o string
    	output file (default "out.png")
  -s int
    	line separation (default 4)
  -x int
    	x offset
  -y int
    	y offset (default -20)
```
