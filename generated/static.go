package service

type Static struct{}


// List `Application` resource options
func (s *Static) GetStaticApplicationOptions () {}

// Get details of a single `Cloud Provider Service`
func (s *Static) GetStaticCloudProviderService () {}

// List `Credential` resource options
func (s *Static) GetStaticCredentialOptions () {}

// List `CronJob` resource options
func (s *Static) GetStaticCronJobOptions () {}

// List `Environment` options
func (s *Static) GetStaticEnvironmentOptions () {}

// List `Network Rule` options
func (s *Static) GetStaticNetworkRuleOptions () {}

// List `Server` options
func (s *Static) GetStaticServerOptions () {}

// List `Service` resource options
func (s *Static) GetStaticServiceOptions () {}

// List `User profile` options
func (s *Static) GetStaticUserProfileOptions () {}

// List `Virtual Host` options
func (s *Static) GetStaticVirtualHostOptions () {}

// List `Cloud Provider Service` instance types by region
func (s *Static) ListStaticCloudInstancesByCloudProviderServiceCodeAndRegionCode () {}

// List available `Role` permissions
func (s *Static) ListStaticPermissions () {}

// List available resource types
func (s *Static) ListStaticResourceTypes () {}
