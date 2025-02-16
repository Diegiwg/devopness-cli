package service

type Credential struct{}


// Add a Credential to the given environment
func (s *Credential) AddEnvironmentCredential () {}

// Delete a given credential
func (s *Credential) DeleteCredential () {}

// Get a credential by ID
func (s *Credential) GetCredential () {}

// Get details of a repository by its name
func (s *Credential) GetCredentialRepository () {}

// Return provider settings
func (s *Credential) GetEnvironmentCredentialSettings () {}

// Get current status of a credential on its provider
func (s *Credential) GetStatusCredential () {}

// Return a list of all repositories belonging to the source provider linked to the credential
func (s *Credential) ListCredentialRepositories () {}

// Return a list of all Credentials belonging to an environment
func (s *Credential) ListEnvironmentCredentials () {}

// Update an existing Credential
func (s *Credential) UpdateCredential () {}
