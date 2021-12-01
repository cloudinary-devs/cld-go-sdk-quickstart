package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Resize_crop(cld *cloudinary.Cloudinary) string {
    // Instantiate an object that stores information for asset with public ID "black_coat_portrait"
    img_fam, err := cld.Image("docs/sdk/go/family_video")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    img_fam.Transformation = "w_350,h_350,c_fill,g_faces"

    // Generate and print the delivery URL
    url, err := img_fam.String()
    fmt.Println(url)
    if err != nil {
        fmt.Println("error")
    }
    return url
}