package api

import "time"

type createTaskReq struct {
	Name        string    `json:"name" bin`
	Description string    `json:"description"` // todo: instead string use some text object
	Important   int       `json:"important"`
	Done        int       `json:"done"`
	Deadline    time.Time `json:"deadline"`
}
