package main

import (
	"strconv"
)

func UptimeMonitorMonitorToBaseMonitorMapper(uptimeMonitor UptimeMonitorMonitor) *Monitor {
	var m Monitor

	m.name = uptimeMonitor.FriendlyName
	m.url = uptimeMonitor.URL
	m.id = strconv.Itoa(uptimeMonitor.ID)

	return &m
}

func UptimeMonitorMonitorsToBaseMonitorsMapper(uptimeMonitors []UptimeMonitorMonitor) []Monitor {
	var monitors []Monitor

	for index := 0; index < len(uptimeMonitors); index++ {
		monitors = append(monitors, *UptimeMonitorMonitorToBaseMonitorMapper(uptimeMonitors[index]))
	}

	return monitors
}
