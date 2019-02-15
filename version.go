package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	Manual   = flag.String("manual", "", "Set Version Manually: [release.major.minor]")
	Minor    = flag.Bool("minor", true, "Increment Minor Version: [release.major.MINOR]")
	Major    = flag.Bool("major", false, "Increment Major Version [release.MAJOR.minor]")
	Release  = flag.Bool("release", false, "Increment Release Version [RELEASE.major.minor]")
	Verbose  = flag.Bool("verbose", false, "Print Out to STDOUT the version change")
	Snapshot = flag.Bool("snapshot", false, "Append `-SNAPSHOT` to the version (doesn't change the version)")

	validationRegexp = regexp.MustCompile(`^([0-9])+\.([0-9])+\.([0-9])+(\-SNAPSHOT)*$`)
)

func UpdateVersion(oldVersion string) (newVersion string, err error) {

	if !validationRegexp.MatchString(oldVersion) {
		err = fmt.Errorf("Version not formatted as 'release.major.minor' - please check the version file to continue.")
		return
	}

	oldVersion = strings.Replace(oldVersion, "-SNAPSHOT", "", -1)

	version := strings.Split(oldVersion, ".")

	releaseVersion, e := strconv.Atoi(version[0])

	if e != nil {
		err = fmt.Errorf("Release Version is non-number (%s)", e.Error())
		return
	}

	majorVersion, e := strconv.Atoi(version[1])

	if e != nil {
		err = fmt.Errorf("Major Version is non-number (%s)", e.Error())
		return
	}

	minorVersion, e := strconv.Atoi(version[2])

	if e != nil {
		err = fmt.Errorf("Minor Version is non-number (%s)", e.Error())
		return
	}

	if *Minor {
		minorVersion = minorVersion + 1
	}

	if *Major {
		minorVersion = 0
		majorVersion = majorVersion + 1
	}

	if *Release {
		minorVersion = 0
		majorVersion = 0
		releaseVersion = releaseVersion + 1
	}

	newVersion = fmt.Sprintf("%v.%v.%v", releaseVersion, majorVersion, minorVersion)

	if *Manual != "" {
		manual := *Manual
		if validationRegexp.MatchString(manual) {
			newVersion = manual
		} else {
			err = fmt.Errorf("Manual Version %s did not match pattern %v", manual, validationRegexp)
		}
	}

	if *Snapshot {
		newVersion = fmt.Sprintf("%s-SNAPSHOT", oldVersion)
	}

	if *Verbose {
		fmt.Println(fmt.Sprintf("%s -> %s", oldVersion, newVersion))
	}

	return
}
