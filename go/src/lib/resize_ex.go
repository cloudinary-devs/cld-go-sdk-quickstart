package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Resize(cld *cloudinary.Cloudinary) string {
	  // Instantiate an object that stores information for asset with public ID "family_video" in folder "docs/sdk/go"
	  my_image, err := cld.Image("docs/sdk/go/apple")
	  if err != nil {
		  fmt.Println("error")
	  }

	  // Add the transformation
	  my_image.Transformation = "c_fill,h_250,w_250"

	  // Generate and print the delivery URL
	  url, err := my_image.String()
	  if err != nil {
		  fmt.Println("error")
	  }
    
    return url
}