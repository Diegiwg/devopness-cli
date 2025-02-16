package service

type Hook struct{}


// Delete a given hook
func (s *Hook) DeleteHook () {}

// Get a hook by ID
func (s *Hook) GetHook () {}

// Get a hook request by ID
func (s *Hook) GetHookRequest () {}

// Returns a list of all hook requests belonging to a hook
func (s *Hook) ListHookRequestsByHookType () {}

// Trigger an incoming hook associated action
func (s *Hook) TriggerHook () {}

// Update an existing hook
func (s *Hook) UpdateHook () {}
