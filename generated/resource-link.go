package service

type ResourceLink struct{}


// Link the given resources
func (s *ResourceLink) LinkResourceLinkToResourceLink () {}

// List linked resources of the given resource
func (s *ResourceLink) ListResourceLinksByResourceType () {}

// List linked resources of specified link type
func (s *ResourceLink) ListResourceLinksByResourceTypeAndLinkType () {}

// Delete a given resource link
func (s *ResourceLink) UnlinkResourceLinkFromResourceLink () {}
