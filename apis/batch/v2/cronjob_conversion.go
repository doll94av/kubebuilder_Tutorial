package v2

import (
    "fmt"
    "strings"

    "sigs.k8s.io/controller-runtime/pkg/conversion"

    "tutorial.kubebuilder.io/project/apis/batch/v1"
)


// ConvertTo converts this CronJob to the Hub version (v1).
func (src *CronJob) ConvertTo(dstRaw conversion.Hub) error {
    dst := dstRaw.(*v1.CronJob)

    sched := src.Spec.Schedule
    scheduleParts := []string{"*", "*", "*", "*", "*"}
    if sched.Minute != nil {
        scheduleParts[0] = string(*sched.Minute)
    }
    if sched.Hour != nil {
        scheduleParts[1] = string(*sched.Hour)
    }
    if sched.DayOfMonth != nil {
        scheduleParts[2] = string(*sched.DayOfMonth)
    }
    if sched.Month != nil {
        scheduleParts[3] = string(*sched.Month)
    }
    if sched.DayOfWeek != nil {
        scheduleParts[4] = string(*sched.DayOfWeek)
    }
    dst.Spec.Schedule = strings.Join(scheduleParts, " ")



    return nil
}


// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *CronJob) ConvertFrom(srcRaw conversion.Hub) error {
    src := srcRaw.(*v1.CronJob)

    schedParts := strings.Split(src.Spec.Schedule, " ")
    if len(schedParts) != 5 {
        return fmt.Errorf("invalid schedule: not a standard 5-field schedule")
    }
    partIfNeeded := func(raw string) *CronField {
        if raw == "*" {
            return nil
        }
        part := CronField(raw)
        return &part
    }
    dst.Spec.Schedule.Minute = partIfNeeded(schedParts[0])
    dst.Spec.Schedule.Hour = partIfNeeded(schedParts[1])
    dst.Spec.Schedule.DayOfMonth = partIfNeeded(schedParts[2])
    dst.Spec.Schedule.Month = partIfNeeded(schedParts[3])
    dst.Spec.Schedule.DayOfWeek = partIfNeeded(schedParts[4])

    return nil
}
