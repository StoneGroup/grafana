package util

import (
    "fmt"
    "strings"
    "regexp"
    "strconv"
    "time"
)

func AdjustSqlTimeRange(sql string) string {
    sqls := strings.Split(sql, ";");
    l := len(sqls);
    for i := 0; i < l; i++ {
        if strings.Index(sqls[i], "time < ") == -1 {
            sqls[i] = strings.Replace(sqls[i], " GROUP", " AND time < now() GROUP", 1);
        }
        re := regexp.MustCompile(`GROUP\sBY\stime\(([0-9]+)m\)`)
        if re.MatchString(sqls[i]) {
            interval, err := strconv.Atoi(re.FindStringSubmatch(sqls[i])[1])
            if err != nil {
                continue;
            }
            minute := time.Now().Minute()
            diff := (minute + 1) % interval
            if diff > 0 {
                sqls[i] = strings.Replace(sqls[i], "time < now()", fmt.Sprintf("time < now() - %dm", diff), 1)
            }
        }
    }

    return strings.Join(sqls, ";");
}
