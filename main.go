package main


import (
    
    "html/template"
    "net/http"
    "github.com/cloudinary/cloudinary-go"
    ec "lib"
    "os"
)

type Transformation struct {
    Description string
    Asset string
    Done  bool
}

type TransformationPageData struct {
    PageTitle string
    Transformations     []Transformation
}

func main() {
    // Add your Cloudinary credentials.
    cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

    url_cartoon := ec.Cartoonify(cld) 
    url_resize := ec.Resize(cld)
    url_resize_crop := ec.Resize_crop(cld)
    url_overlays := ec.Overlays(cld) 
    url_faces_gravity := ec.Faces_gravity(cld)
    url_auto_format := ec.Auto_format(cld)
    url_specified_format := ec.Specified_format(cld)
    url_sepia := ec.Sepia(cld)
  
    tmpl := template.Must(template.ParseFiles("transformations.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TransformationPageData{
          PageTitle: "Cloudinary Transformations: Go SDK",
            Transformations: []Transformation{
                {Description: "Cartoonified", Asset: url_cartoon, Done: false},
                {Description: "Resized", Asset: url_resize, Done: true},
                {Description: "Resized & Cropped", Asset: url_resize_crop, Done: true},
                {Description: "Text and Image Overlays", Asset: url_overlays, Done: true},
                {Description: "Gravity with Face Detection", Asset: url_faces_gravity, Done: true},
                {Description: "Auto format", Asset: url_auto_format, Done: true},
                {Description: "Specified format", Asset: url_specified_format, Done: true},
                {Description: "Sepia effect", Asset: url_sepia, Done: true},
            },
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":8000", nil)
}

