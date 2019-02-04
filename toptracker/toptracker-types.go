package main

type Profile struct {
	ID int `json:"id"`
}
type User struct {
	Token string `json:"access_token"`
	User  struct {
		ID int `json:"id"`
	}
	Profiles []Profile
}

type Project struct {
	ID int
}
type Projects struct {
	Projects []Project
}

type TopTrackerInfo struct {
	SecondsPerMonth int
	SecondsPerDay   int
}

/**
{
  "work_summary": {
    "workers_projects": [
      {
        "id": 85725,
        "name": "Andrey",
        "projects": [
          {
            "id": 280887,
            "name": "StreamLayer",
            "seconds": 45291,
            "amount": null,
            "currency": null
          }
        ]
      }
    ]
  }
}
*/

type project struct {
	Seconds int `json:"seconds"`
}

type workersProject struct {
	Projects []project `json:"projects"`
}

type workSummary struct {
	WorkersProjects []workersProject `json:"workers_projects"`
}

type workSummaryInfo struct {
	WorkSummary workSummary `json:"work_summary"`
}
