package utility

type Utility struct{}

const IMAGE_URL = "https://siakadu.unila.ac.id/uploads/fotomhs/"

var POSSIBLE_IMAGE_EXT []string = []string{".jpg", ".jpeg", ".png"}
var BASE_IMAGE_STORAGE = "./result/images/"

func New() *Utility {
	return &Utility{}
}
