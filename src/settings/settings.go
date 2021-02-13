package settings

type Settings struct {
	PluginName             string
	PolicyMinVulnerability string
	PolicyRegistry         []string
	PolicyImageName        []string
	PolicyNonCompliant     bool
	IgnoreRegistry         []string
	IgnoreImageName        []string

	AggregateIssuesNumber   int
	AggregateTimeoutSeconds int
	IsScheduleRun           bool
	PolicyOnlyFixAvailable  bool
	PolicyShowAll           bool
	AquaServer              string
}

func GetDefaultSettings() *Settings {
	return &Settings{
		PluginName:             "",
		PolicyMinVulnerability: "",
		PolicyRegistry:         []string{},
		PolicyImageName:        []string{},
		PolicyShowAll:          false,
		PolicyNonCompliant:     false,
		IgnoreRegistry:         []string{},
		IgnoreImageName:        []string{},
	}
}
