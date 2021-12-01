package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Faces_gravity(cld *cloudinary.Cloudinary) string {
        // Instantiate an object that stores information for asset with public ID "black_coat_portrait" in folder "docs/sdk/go"
    img, err := cld.Image("docs/sdk/go/black_coat_portrait")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    img.Transformation = "f_auto/q_auto/c_crop,h_350,w_300,x_1400,y_1200/c_fill,h_100,w_130"

    // Generate and print the delivery URL
    url, err := img.String()
    fmt.Println(url)
    if err != nil {
        fmt.Println("error")
    }
    return url
}