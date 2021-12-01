package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Specified_format(cld *cloudinary.Cloudinary) string {
    // Instantiate an object that stores information for asset with public ID "cloud_castle" in folder "docs/sdk/go"
    img_fam, err := cld.Image("docs/sdk/go/cloud_castle")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    img_fam.Transformation = "f_png/w_300"

    // Generate and print the delivery URL
    url, err := img_fam.String()
    fmt.Println(url)
    if err != nil {
        fmt.Println("error")
    }
    return url
}