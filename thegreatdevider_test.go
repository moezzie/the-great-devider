package thegreatdevider

import (
    "testing"
)

func TestCreateGrid(t *testing.T) {
    
    expected_num_chunks := 7500
    sub_images, err     := GridFromFile("testdata/testimg.png" ,8, 8);

    if err != nil {
        t.Error(err)
    }

    if len(sub_images) < expected_num_chunks {
        t.Error("Incorrect number of sub_images. Expected", expected_num_chunks, "got", len(sub_images))
    }
}

func TestGridFromData(t *testing.T) {
   
   expected_num_images  := 7500
   sub_images, err      := GridFromData()

    if err != nil {
        t.Error(err)
    }

    if len(sub_images) < expected_num_chunks {
        t.Error("Incorrect number of sub_images. Expected", expected_num_chunks, "got", len(sub_images))
    }

}

func TestCreateSubimage(t *testing.T) {

    width       := 8
    height      := 8
    origin_x    := 1
    origin_y    := 1

    sub_image, err := SubImageFromFile("testdata/testimg.png", width, height, origin_x, origin_y)
    if err != nil {
        t.Error(err)
    }

    if sub_image.Bounds().Max.X != width || sub_image.Bounds().Max.Y != height {
        t.Error("Sub image is not the size it was expected to be")
    }

}