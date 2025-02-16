package service

type Pipeline struct{}


// Add a Pipeline to a resource
func (s *Pipeline) AddPipeline () {}

// Create an action to run a Pipeline
func (s *Pipeline) AddPipelineAction () {}

// Create a hook to a specific pipeline
func (s *Pipeline) AddPipelineHook () {}

// Add a step to a pipeline
func (s *Pipeline) AddPipelineStep () {}

// Delete a given Pipeline
func (s *Pipeline) DeletePipeline () {}

// Get a Pipeline by ID
func (s *Pipeline) GetPipeline () {}

// Link a step to a Pipeline
func (s *Pipeline) LinkStepToPipeline () {}

// Return a list of pipeline's actions
func (s *Pipeline) ListPipelineActions () {}

// List all hooks in a pipeline
func (s *Pipeline) ListPipelineHooks () {}

// Return a list of pipelines to a resource
func (s *Pipeline) ListPipelinesByResourceType () {}

// Unlink a step from a Pipeline
func (s *Pipeline) UnlinkStepFromPipeline () {}

// Update an existing Pipeline
func (s *Pipeline) UpdatePipeline () {}

// Update an existing Pipeline Step
func (s *Pipeline) UpdatePipelineStep () {}
