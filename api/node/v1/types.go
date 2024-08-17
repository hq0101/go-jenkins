package v1

type Node struct {
	Class          string `json:"_class" xml:"_class,attr"`
	BusyExecutors  int    `json:"busyExecutors" xml:"busyExecutors"`
	DisplayName    string `json:"displayName" xml:"displayName"`
	TotalExecutors int    `json:"totalExecutors" xml:"totalExecutors"`
	Computer       []struct {
		Class          string              `json:"_class" xml:"_class,attr"`
		Actions        []map[string]string `json:"actions" xml:"actions"`
		AssignedLabels []struct {
			Name string `json:"name" xml:"name"`
		} `json:"assignedLabels" xml:"assignedLabels"`
		Description string `json:"description" xml:"description"`
		DisplayName string `json:"displayName" xml:"displayName"`
		Executors   []struct {
		} `json:"executors" xml:"executors"`
		Icon            string `json:"icon" xml:"icon"`
		IconClassName   string `json:"iconClassName" xml:"iconClassName"`
		Idle            bool   `json:"idle" xml:"idle"`
		JnlpAgent       bool   `json:"jnlpAgent" xml:"jnlpAgent"`
		LaunchSupported bool   `json:"launchSupported" xml:"launchSupported"`
		LoadStatistics  struct {
			Class string `json:"_class" xml:"_class,attr"`
		} `json:"loadStatistics" xml:"loadStatistics"`
		ManualLaunchAllowed bool `json:"manualLaunchAllowed" xml:"manualLaunchAllowed"`
		MonitorData         struct {
			HudsonNodeMonitorsSwapSpaceMonitor struct {
				Class                   string `json:"_class" xml:"_class,attr"`
				AvailablePhysicalMemory int    `json:"availablePhysicalMemory" xml:"availablePhysicalMemory"`
				AvailableSwapSpace      int    `json:"availableSwapSpace" xml:"availableSwapSpace"`
				TotalPhysicalMemory     int    `json:"totalPhysicalMemory" xml:"totalPhysicalMemory"`
				SwapSpaceUsed           int    `json:"swapSpaceUsed" xml:"swapSpaceUsed"`
			} `json:"hudson.node_monitors.SwapSpaceMonitor" xml:"hudson.node_monitors.SwapSpaceMonitor"`
			HudsonNodeMonitorsTemporarySpaceMonitor struct {
				Class            string `json:"_class" xml:"_class,attr"`
				Timestamp        int64  `json:"timestamp" xml:"timestamp"`
				Path             string `json:"path" xml:"path"`
				Size             int64  `json:"size" xml:"size"`
				ThreShold        int64  `json:"threshold" xml:"threshold"`
				TotalSize        int64  `json:"totalSize" xml:"totalSize"`
				WarningThreshold int64  `json:"warningThreshold" xml:"warningThreshold"`
			} `json:"hudson.node_monitors.TemporarySpaceMonitor" xml:"hudson.node_monitors.TemporarySpaceMonitor"`
			HudsonNodeMonitorsDiskSpaceMonitor struct {
				Class            string `json:"_class" xml:"_class,attr"`
				Timestamp        int64  `json:"timestamp" xml:"timestamp"`
				Path             string `json:"path" xml:"path"`
				Size             int64  `json:"size" xml:"size"`
				ThreShold        int64  `json:"threshold" xml:"threshold"`
				TotalSize        int64  `json:"totalSize" xml:"totalSize"`
				WarningThreshold int64  `json:"warningThreshold" xml:"warningThreshold"`
			} `json:"hudson.node_monitors.DiskSpaceMonitor" xml:"hudson.node_monitors.DiskSpaceMonitor"`
			HudsonNodeMonitorsArchitectureMonitor string `json:"hudson.node_monitors.ArchitectureMonitor" xml:"hudson.node_monitors.ArchitectureMonitor"`
			HudsonNodeMonitorsResponseTimeMonitor struct {
				Class     string `json:"_class" xml:"_class,attr"`
				Timestamp int64  `json:"timestamp" xml:"timestamp"`
				Average   int    `json:"average" xml:"average"`
			} `json:"hudson.node_monitors.ResponseTimeMonitor" xml:"hudson.node_monitors.ResponseTimeMonitor"`
			HudsonNodeMonitorsClockMonitor struct {
				Class string `json:"_class" xml:"_class,attr"`
				Diff  int    `json:"diff" xml:"diff"`
			} `json:"hudson.node_monitors.ClockMonitor" xml:"hudson.node_monitors.ClockMonitor"`
		} `json:"monitorData" xml:"monitorData"`
		NumExecutors int  `json:"numExecutors" xml:"numExecutors"`
		Offline      bool `json:"offline" xml:"offline"`
		OfflineCause struct {
		} `json:"offlineCause" xml:"offlineCause"`
		OfflineCauseReason string     `json:"offlineCauseReason" xml:"offlineCauseReason"`
		OneOffExecutors    []struct{} `json:"oneOffExecutors" xml:"oneOffExecutors"`
		TemporarilyOffline bool       `json:"temporarilyOffline" xml:"temporarilyOffline"`
	} `json:"computer" xml:"computer"`
}
