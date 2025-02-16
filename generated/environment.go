package service

type Environment struct{}


// Create a new environment on the current project
func (s *Environment) AddProjectEnvironment () {}

// Archive an environment
func (s *Environment) ArchiveEnvironment () {}

// Get an environment by ID
func (s *Environment) GetEnvironment () {}

// Return a list of all archived environments belonging to a project
func (s *Environment) ListProjectArchivedEnvironments () {}

// Return a list of all environments belonging to a project
func (s *Environment) ListProjectEnvironments () {}

// Unarchive an environment
func (s *Environment) UnarchiveEnvironment () {}

// Update a given environment
func (s *Environment) UpdateEnvironment () {}
