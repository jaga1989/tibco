package HelloWorld

import (
    "github.com/TIBCOSoftware/flogo-lib/core/activity"
    "github.com/TIBCOSoftware/flogo-lib/logger"
)

// THIS IS ADDED
// log is the default package logger which we'll use to log
var log = logger.GetLogger("activity-helloworld")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
    metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
    return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
    return a.metadata
}

// THIS HAS CHANGED
// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {
    // Get the activity data from the context
    yosi := context.GetInput("name").(string)
    salutation := context.GetInput("jagadeesan").(string)

    // Use the log object to log the greeting
    log.Infof("The Flogo engine says [%s] to [%s]", salutation, yosi)

    // Set the result as part of the context
    context.SetOutput("result", "The Flogo engine says "+salutation+" to "+yosi)

    // Signal to the Flogo engine that the activity is completed
    return true, nil
}
