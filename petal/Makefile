# We use our registry
REGISTRY ?= foobar.io

# We keep the structure of the original forked repo
IMAGE_PATH ?= apache/camel-k

# We are covering both bases for tag retrieval, not assuming that the tag was annotated (which
# would be better)
get-version:
	@git describe --tag

# Return our own registry and image path as the image name
get-image-name:
	@echo ${REGISTRY}/${IMAGE_PATH}
