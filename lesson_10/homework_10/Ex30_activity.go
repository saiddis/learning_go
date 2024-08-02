package main

import "time"

type Activity struct {
	activityType string
	timeStamp    time.Time
}

type UserActivityTracker struct {
	Activities []Activity
}

func (uat *UserActivityTracker) AddActivity(activityType string, timestamp time.Time) {
	activity := Activity{
		activityType: activityType,
		timeStamp:    timestamp,
	}

	uat.Activities = append(uat.Activities, activity)
}

func (uat UserActivityTracker) ActivityAfterTime(timestamp time.Time) []Activity {
	var actsAfterTime []Activity

	for _, v := range uat.Activities {
		if v.timeStamp.Compare(timestamp) == 1 {
			actsAfterTime = append(actsAfterTime, v)
		}
	}

	return actsAfterTime
}
