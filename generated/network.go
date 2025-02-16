package service

type Network struct{}


// Create a new network for the given environment
func (s *Network) AddEnvironmentNetwork () {}

// Delete a given network
func (s *Network) DeleteNetwork () {}

// Get a network by ID
func (s *Network) GetNetwork () {}

// Get current status of a network
func (s *Network) GetStatusNetwork () {}

// Return a list of all networks belonging to an environment
func (s *Network) ListEnvironmentNetworks () {}

// Update an existing Network
func (s *Network) UpdateNetwork () {}
