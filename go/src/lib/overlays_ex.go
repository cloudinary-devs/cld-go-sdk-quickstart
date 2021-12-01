package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Overlays(cld *cloudinary.Cloudinary) string {
    img_coffee, err := cld.Image("docs/sdk/go/coffee_cup")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    img_coffee.Transformation = "w_400,h_250,c_fill,g_south/l_nice_couple,w_1.3,h_1.3,g_faces,c_crop,fl_region_relative/e_saturation:50/e_vignette/fl_layer_apply,w_100,r_max,g_center,y_20,x_-20/l_balloon,h_55/e_hue:-20,a_5/fl_layer_apply,x_30,y_5/l_text:Cookie_40_bold:Love,co_rgb:f08/a_20,fl_layer_apply,x_-45,y_44/c_crop,w_300,h_250,x_30/r_60"

    // Generate and print the delivery URL
    url, err := img_coffee.String()
    fmt.Println(url)
    if err != nil {
        fmt.Println("error")
    }
    return url
}