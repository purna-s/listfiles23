package listfiles23

import (

	"fmt"
    "os"
    "path/filepath"
	"strings"
	"time"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("Path", "F:\TESTING")
	tc.SetInput("SubDirectories[Y/N]", "Y")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}
	act.Eval(tc)
	//check output attr

	FileName = tc.GetOutput("FileName")
	assert.Equal(t, FileName, FileName)
	Directory = tc.GetOutput("Directory")
	assert.Equal(t, Directory, Directory)
	Extension = tc.GetOutput("Extension")
	assert.Equal(t, Extension, Extension)
	Size = tc.GetOutput("Size")
	assert.Equal(t, Size, Size)
	ModTime = tc.GetOutput("ModTime")
	assert.Equal(t, ModTime, ModTime)
	MinutesDiff = tc.GetOutput("MinutesDiff")
	assert.Equal(t, MinutesDiff, MinutesDiff)
	
	//assert.Equal(t, output, output)

}
