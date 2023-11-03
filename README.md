# http-polygon

## Purpose of the app
1. Use the api endpoint to upload a file and draw any polygon in any RGBA color on top of it, then download the processed file
2. Import the Draw method from `pkg/polygon` to simply draw polygons by yourself

## Used algorithm for drawing
**scanline filling with linear interpolation** implemented in `Draw()` function
1. iteration starts from the top of an image, `y` indicates a linear function which is a straight horizontal line on the cartesian - `y = i`
2. vars x1, x2 define range between pixels should be colored with the fillColor arg
3. inner for loop goes through all provided vertices
4. `j` is the index of a `next` vertex to treat as a pair - mathematically we could write it as a pair of points:
   $$(x_i, y_i)(x_j, y_j)$$
    Modular arithmetic is used due to have a full circle of pairing vertices. Last vertex should be paired with the first one (index 0), so in the last iteration, the expression will return [0]
5. check if our horizontal line is in the range of at least one vertex
6. using linear interpolation, calculate the `x` based on the `y` function crossing the pair of points (vertices)
7. following math formula is being used to calculate it:
   $$x = x_i + { x_j - x_i \over y_j - y_i} * (y - y_i)$$
   which is just a simple transformation of regular formula for calculating the function crosses two points on cartesian:
   $$y - y_i = { y_j - y_i \over x_j - x_i} * (x - x_i)$$
8. the purpose of it is to find the `x` value which is a product of vertices intersection with `y = i` that corresponds to the end of the painting area (horizontal). Following img will help to understand it.
9. finally, iterate over horizontal line (from x1 to x2) and set the pixel color
