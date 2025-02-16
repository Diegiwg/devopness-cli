package service

type NetworkRule struct{}


// Add a Network Rule to the given environment
func (s *NetworkRule) AddEnvironmentNetworkRule () {}

// Delete a given Network Rule
func (s *NetworkRule) DeleteNetworkRule () {}

// Get a Network Rule by ID
func (s *NetworkRule) GetNetworkRule () {}

// Return a list of all Network Rules belonging to an environment
func (s *NetworkRule) ListEnvironmentNetworkRules () {}

// Update an existing Network Rule
func (s *NetworkRule) UpdateNetworkRule () {}
