#!/bin/sh

# This sample application is part of the Timestream prerelease documentation. 
# The prerelease documentation is confidential and is provided under the terms
# of your nondisclosure agreement with Amazon Web Services (AWS) or other 
# agreement governing your receipt of AWS confidential information.


if test $# -ne 2; then
    echo "Usage:"
    echo "./$2.sh <database name> <table name>"
    echo "e.g. ./$2.sh DATABASE_NAME TABLE_NAME"
    exit -1
fi

# find the host with the maximum data points
go run query-sample.go -query "SELECT hostname, count(*) as hostcount
FROM $1.$2
WHERE measure_name = 'cpu_utilization'
    AND time > ago(2h)
GROUP BY hostname
ORDER BY count(*) DESC"
 
