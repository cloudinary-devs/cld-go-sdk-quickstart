package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Sepia(cld *cloudinary.Cloudinary) string {
    // Instantiate an object that stores information for asset with public ID "pond_reflect" in folder "docs/sdk/go"
    i, err := cld.Image("docs/sdk/go/maroon_hat")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    i.Transformation = "f_auto/q_auto/c_fill,e_sepia:80,g_face,h_250,w_250"

    // Generate and print the delivery URL
    url, err := i.String()
    fmt.Println(url)
    if err != nil {
        fmt.Println("error")
    }
    return url
}