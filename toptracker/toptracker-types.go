package main

type userInfo struct {
	Token string `json:"access_token"`
	User  struct {
		ID int `json:"id"`
	}
	Profiles []struct {
		ID int `json:"id"`
	}
}

type projectsInfo struct {
	Projects []struct {
		ID int
	}
}

type topTrackerInfo struct {
	SecondsPerMonth int
	SecondsPerDay   int
}

type workSummaryInfo struct {
	WorkSummary struct {
		WorkersProjects []struct {
			Projects []struct {
				Seconds int `json:"seconds"`
			} `json:"projects"`
		} `json:"workers_projects"`
	} `json:"work_summary"`
}
