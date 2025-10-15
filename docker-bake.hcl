variable "IMAGE_TAG" {
  type = string
  default = "0.0.0"
}

target "default" {
	context = "."
	dockerFile = "Dockerfile"
	tags = ["seddy.com/go-getting-started:${IMAGE_TAG}"]
}

