package service

type Action struct{}


// Get an action by ID
func (s *Action) GetAction () {}

// Get action target step log
func (s *Action) GetActionLog () {}

// Return a list of all actions belonging to current user
func (s *Action) ListActions () {}

// List resource actions of an action type
func (s *Action) ListActionsByResourceType () {}

// List actions triggered to a given action target resource
func (s *Action) ListActionsByTargetResourceType () {}

// List environment actions
func (s *Action) ListEnvironmentActions () {}

// List environment actions of a resource type
func (s *Action) ListEnvironmentActionsByResourceType () {}

// List project actions of a resource type
func (s *Action) ListProjectActionsByResourceType () {}

// Retry an action
func (s *Action) RetryAction () {}
