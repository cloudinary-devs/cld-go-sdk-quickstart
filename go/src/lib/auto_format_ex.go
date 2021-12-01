package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Auto_format(cld *cloudinary.Cloudinary) string {
    // Instantiate an object that stores information for asset with public ID "pond_reflect" in folder "docs/sdk/go"
    img_lake, err := cld.Image("docs/sdk/go/pond_reflect")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    img_lake.Transformation = "f_auto/q_auto/w_300"

    // Generate and print the delivery URL
    url, err := img_lake.String()
    fmt.Println(url)
    if err != nil {
        fmt.Println("error")
    }
    return url
}