package service

type CronJob struct{}


// Add a Cron Job to the given environment
func (s *CronJob) AddEnvironmentCronJob () {}

// Delete a given Cron Job
func (s *CronJob) DeleteCronJob () {}

// Get a Cron Job by ID
func (s *CronJob) GetCronJob () {}

// Return a list of all Cron Jobs belonging to an environment
func (s *CronJob) ListEnvironmentCronJobs () {}

// Update an existing Cron Job
func (s *CronJob) UpdateCronJob () {}
