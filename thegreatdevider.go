package thegreatdevider

import(
    "image"
    "image/png"
    "image/jpeg"
    "os"
    "errors"
    "bytes"
)

type Image struct {
    Matrix image.Image
}

type SubImager interface {
    SubImage(r image.Rectangle) image.Image
}

func init() {
    // damn important or else At(), Bounds() functions will
    // caused memory pointer error!!
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
    image.RegisterFormat("jpg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
    image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func LoadFromFile(file_path string) (image.Image, error) {

    var matrix image.Image

    // Open the file for reading
    reader, err := os.Open(file_path)
    if err != nil {
        return matrix, err
    }
    defer reader.Close()

    // Make sense of image file data
    matrix, _, err = image.Decode(reader)
    if err != nil {
        return matrix, err
    }

    return matrix, nil
}

func GridFromFile(file_path string, width, height int) ([]image.Image, error) {

    var subimages   []image.Image   = make([]image.Image, 0)

    matrix, err := LoadFromFile(file_path)
    if err != nil {
        return subimages, err
    }

    return Grid(matrix, width, height)
}

func Grid(matrix image.Image, width, height int) ([]image.Image, error) {

    var subimages   []image.Image   = make([]image.Image, 0)
    var max_x int = matrix.Bounds().Max.X
    var max_y int = matrix.Bounds().Max.Y


    // Make sure we have acceptabe input values
    if width < 1 || height < 1 {
        return subimages, errors.New("Both height and width of grid subimages must be greater than 0")
    }

    if max_x < width && max_y < height {
        return subimages, errors.New("Size of subimages must be less than or equal to the size of the source image")
    } 

    num_subs_x  :=  max_x/ width
    num_subs_y  :=  max_y / height
    subimages   =   make([]image.Image, num_subs_x * num_subs_y )
    count       :=  0

    var rectangle   image.Rectangle

    // Devide the image into x*y adjecent images
    for y := 0; y < num_subs_y; y += height {
        for x := 0; x < num_subs_x; x += width {

            rectangle = image.Rect(x, y, (x * width), (y * height))

            // Make sure we don't go out of bounds
            if rectangle.Min.X >= 0 && rectangle.Min.Y >= 0 && rectangle.Max.X <= max_x && rectangle.Max.Y <= max_y {
                subimages[count] = matrix.(SubImager).SubImage( rectangle )
            }

            count++
        }
    }

    return subimages, nil
}

func SubImageFromFile(file_path string, width, height, origin_x, origin_y int) (image.Image, error) {

    var sub_image image.Image

    matrix, err := LoadFromFile(file_path)
    if err != nil {
        return sub_image, err
    }

    return SubImage(matrix, width, height, origin_x, origin_y)
}

func SubImage(matrix image.Image, width, height, origin_x, origin_y int) (image.Image, error) {

    var img     image.Image

    if (width < 1 || width > matrix.Bounds().Max.X || height < 1 || height > matrix.Bounds().Max.Y ) {
        return img, errors.New("Width or height of the subimage must be larger than 0 but less than the size of the source image")
    }

    if (origin_x < 0 || origin_x > matrix.Bounds().Max.X || origin_y < 0 || origin_y > matrix.Bounds().Max.Y) {
        return img, errors.New("Origins X and Y of the subimage must be larger than or equal to 0 but less than the size of the source image")
    }

    rectangle := image.Rect(origin_x, origin_y, origin_x + width, origin_y + height)

    return matrix.(SubImager).SubImage(rectangle), nil

}

func ImageToBytes(matrix image.Image) ([]byte, error){
    buf := new(bytes.Buffer)
    err := png.Encode(buf, matrix)
    return buf.Bytes(), err
}

