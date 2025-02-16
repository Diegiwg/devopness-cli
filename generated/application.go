package service

type Application struct{}


// Trigger a new deployment for current application
func (s *Application) AddApplicationDeployment () {}

// Create a new variable linked to an application
func (s *Application) AddApplicationVariable () {}

// Create a new application
func (s *Application) AddEnvironmentApplication () {}

// Delete a given application
func (s *Application) DeleteApplication () {}

// Get an application by ID
func (s *Application) GetApplication () {}

// List all hooks in an application
func (s *Application) ListApplicationHooks () {}

// Return a list of variables belonging to an application
func (s *Application) ListApplicationVariables () {}

// Return a list of all Applications belonging to an environment
func (s *Application) ListEnvironmentApplications () {}

// Update an existing application
func (s *Application) UpdateApplication () {}
