package service

type Daemon struct{}


// Add a Daemon to the given environment
func (s *Daemon) AddEnvironmentDaemon () {}

// Delete a given Daemon
func (s *Daemon) DeleteDaemon () {}

// Get a Daemon by ID
func (s *Daemon) GetDaemon () {}

// Get current status of a daemon
func (s *Daemon) GetStatusDaemon () {}

// Return a list of all Daemons belonging to an environment
func (s *Daemon) ListEnvironmentDaemons () {}

// Restart a Daemon
func (s *Daemon) RestartDaemon () {}

// Start a Daemon
func (s *Daemon) StartDaemon () {}

// Stop a Daemon
func (s *Daemon) StopDaemon () {}

// Update an existing Daemon
func (s *Daemon) UpdateDaemon () {}
