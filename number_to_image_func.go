package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	//so I presume using some sort of function Im supposed to output an image seems like
	// so the way im thinking of how u first create a base unallocated double slice
	image := make([][]uint8, dy)
	// write an outerloop which goes over all the columns
	for i := 0; i < dy; i++ {
		// create a single slice of uint8 type
		a := make([]uint8, dx)
		//create another loop for each slice in there
		for j := 0; j < dx; j++ {
			//just wrote a randomn math function for each pixel of the slice
			grayscale := uint8((i ^ j) / 2)
			// allocate it to that number it wont overflow since its uint8 only has 255 possible vals
			//so will just loop over
			a[j] = grayscale
		}
		//keep adding each row of the image to our row we had made
		image[i] = a

	}
	return image
}

func main() {
	//this simply outputs our image
	pic.Show(Pic)
}
