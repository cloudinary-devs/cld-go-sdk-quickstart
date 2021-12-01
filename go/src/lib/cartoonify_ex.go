package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Cartoonify(cld *cloudinary.Cloudinary) string {
   // Instantiate an object that stores information for asset with public ID "actor"
    img_actor, err := cld.Image("docs/sdk/go/actor")
    if err != nil {
      fmt.Println("error")
    }

    // Add the transformation
    img_actor.Transformation = "e_cartoonify/r_max/e_outline:100,co_lightblue/b_lightblue,h_300,c_pad"

    // Generate and print the delivery URL
    url, err := img_actor.String()
    fmt.Println(url)
    if err != nil {
      fmt.Println("error")
    }
    return url
}