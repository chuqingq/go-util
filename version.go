package util

import (
	"fmt"
	"runtime/debug"
	"time"
)

func GetVersion() string {
	const MAIN_VERSION = "v1.1.0"

	vcsTime := "unknown"
	vcsRevision := "unknown"
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			fmt.Printf("key: %v, value: %v\n", setting.Key, setting.Value)
			if setting.Value == "" {
				continue
			}

			switch setting.Key {
			case "vcs.revision":
				vcsRevision = setting.Value
				if len(vcsRevision) > 8 {
					vcsRevision = vcsRevision[:8]
				}
			case "vcs.time":
				vcsTime = setting.Value
				// 2023-08-10T13:49:44Z
				t, err := time.Parse(time.RFC3339, vcsTime)
				if err != nil {
					fmt.Printf("parse vcs.time[%v] err: %v", vcsTime, err)
					continue
				} else {
					vcsTime = t.Format("20060102150405")
				}
			}
		}
	}

	return MAIN_VERSION + "-" + vcsTime + "-" + vcsRevision
}

