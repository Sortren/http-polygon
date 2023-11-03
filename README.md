# http-polygon

## Purpose of the app
1. Use the api endpoint to upload a file and draw any polygon in any RGBA color on top of it, then download the processed file
2. Import the Draw method from `pkg/polygon` to simply draw polygons by yourself

## Running the app
In the root directory, simply type:
```shell
go run cmd/app.go
```
and the server will start working on localhost:8080, with available endpoint: `/draw-polygon`

## Endpoint request restrictions (`/draw-polygon` and `/v2/draw-polygon`)
1. Use POST request 
2. Use form-data
3. Three keys needs to be setup: "photo, vertices, color". **Everything is in the postman collection in the root folder in this repo**
4. One thing requires manual configuration - path to the photo, simply click at the value and change it to the repository's initial.png file (or any other, as you wish)
5. Right top corner, click the arrow and then *Send and Download*
6. v2 endpoint is concurrent version of a regular `/draw-polygon`, service constructor requires max number of goroutines to process the request

## Used algorithm for drawing
**scanline filling with linear interpolation** implemented in `Draw()` function
1. iteration starts from the minY value of vertices (top of a polygon), `y` indicates a linear function which is a straight horizontal line on the cartesian - `y = i`
2. vars x1, x2 define range between pixels should be colored with the fillColor arg
3. inner for loop goes through all provided vertices
4. `j` is the index of a `next` vertex to treat as a pair - mathematically we could write it as a pair of points:
   $$(x_i, y_i)(x_j, y_j)$$
    Modular arithmetic is used due to have a full circle of pairing vertices. Last vertex should be paired with the first one (index 0), so in the last iteration, the expression will return [0]
5. check if our horizontal line goes through the edge
6. using linear interpolation, calculate the `x` based on the `y` function crossing the pair of points (vertices)
7. following math formula is being used to calculate it:
   $$x = x_i + { x_j - x_i \over y_j - y_i} * (y - y_i)$$
   which is just a simple transformation of regular formula for calculating the function crosses two points on cartesian:
   $$y - y_i = { y_j - y_i \over x_j - x_i} * (x - x_i)$$
8. the purpose of it is to find the `x` value which is a product of vertices intersection with `y = i` that corresponds to the end of the painting area (horizontal). Following img will help to understand it.
9. finally, iterate over horizontal line (from x1 to x2) and set the pixel color


## Example shapes 
1. four vertices
   ```json
     [
        {
           "x": 200,
           "y": 50
        },
        {
           "x": 350,
           "y": 250
        },
        {
           "x": 50,
           "y": 250
        },
        {
           "x": 50,
           "y": 100
        }
     ]
    ```
2. five vertices
   ```json
     [
        {
           "x": 200,
           "y": 50
        },
        {
           "x": 300,
           "y": 100
        },
        {
           "x": 250,
           "y": 250
        },
        {
           "x": 150,
           "y": 250
        }, 
        {
           "x": 100,
           "y": 100
        }
     ]
   ```
