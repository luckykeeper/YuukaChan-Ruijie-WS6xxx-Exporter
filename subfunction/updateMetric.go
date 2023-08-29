// 优香酱锐捷 AC Exporter ，适配 WS6xxx 系列设备 - 指标发布
// @CreateTime : 2023/8/28 10:52
// @LastModified : 2023/8/28 10:52
// @Author : Luckykeeper
// @Contact : luckykeeper@luckykeeper.site | https://github.com/luckykeeper | https://luckykeeper.site
// @ProgramEntry: yuukaExporter.go
// @Project : yuukaChan-Ruijie-WS6xxx-Exporter

package subFunction

import (
	"fmt"
	"sort"
	"strings"

	"github.com/VictoriaMetrics/metrics"
)

// 更新指标
func UpdateMetric(name string, labels map[string]string, value float64) {
	// Construct `metric{labels}`
	var labelValues []string
	for k, v := range labels {
		labelValues = append(labelValues, fmt.Sprintf("%s=%q", k, v))
	}
	sort.Strings(labelValues)
	metricName := fmt.Sprintf("%s{%s}", name, strings.Join(labelValues, ","))

	// Update the counter
	metrics.GetOrCreateFloatCounter(metricName).Set(value)
}
