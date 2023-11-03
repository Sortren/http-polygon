package api

import (
	"encoding/json"
	"http-polygon/pkg/polygon"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
)

type FormKey string

var (
	PhotoKey    FormKey = "photo"
	VerticesKey FormKey = "vertices"
	ColorKey    FormKey = "color"
)

type RGBAColorRequest struct {
	R uint8 `json:"r,omitempty"`
	G uint8 `json:"g,omitempty"`
	B uint8 `json:"b,omitempty"`
	A uint8 `json:"a,omitempty"`
}

func DrawPolygonOnFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var rgbaColor RGBAColorRequest
	if err := json.Unmarshal([]byte(r.FormValue(string(ColorKey))), &rgbaColor); err != nil {
		http.Error(w, "invalid RGBA Color", http.StatusBadRequest)
		return
	}

	fillColor := color.RGBA{
		R: rgbaColor.R,
		G: rgbaColor.G,
		B: rgbaColor.B,
		A: rgbaColor.A,
	}

	inputFile, _, err := r.FormFile(string(PhotoKey))
	if err != nil {
		http.Error(w, "invalid photo file", http.StatusBadRequest)
		return
	}
	defer inputFile.Close()

	decodedInputImage, _, err := image.Decode(inputFile)
	if err != nil {
		log.Printf("error while decoding image: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var vertices polygon.Vertices
	if err := json.Unmarshal([]byte(r.FormValue(string(VerticesKey))), &vertices); err != nil {
		http.Error(w, "invalid vertices format", http.StatusBadRequest)
		return
	}

	if len(vertices) < 2 {
		http.Error(w, "number of passed vertices should be greater than 2", http.StatusBadRequest)
		return
	}

	decodedInputImageBounds := decodedInputImage.Bounds()
	minVertex, maxVertex := polygon.ImageBoundsToMinMaxVertex(decodedInputImageBounds)

	if err := vertices.Validate(minVertex, maxVertex); err != nil {
		http.Error(w, "invalid coordinates", http.StatusBadRequest)
		return
	}

	resultImage := image.NewRGBA(decodedInputImageBounds)

	// copy source of the input image into our result image to be able to draw polygon on top of it
	draw.Draw(resultImage, decodedInputImageBounds, decodedInputImage, decodedInputImageBounds.Min, draw.Src)

	polygon.Draw(resultImage, vertices, fillColor)

	w.Header().Set("Content-Disposition", "attachment; filename=processed.jpg")
	w.Header().Set("Content-Type", "application/octet-stream")

	if err := jpeg.Encode(w, resultImage, nil); err != nil {
		log.Printf("can't encode jpeg: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
