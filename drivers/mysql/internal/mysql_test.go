package driver

import (
	"testing"

	"github.com/datazip-inc/olake/constants"
	"github.com/datazip-inc/olake/types"
	"github.com/datazip-inc/olake/utils/testutils"
)

func TestMySQLIntegration(t *testing.T) {
	t.Parallel()
	testConfig := &testutils.IntegrationTest{
		TestConfig:                       testutils.GetTestConfig(string(constants.MySQL)),
		Namespace:                        "olake_mysql_test",
		ExpectedData:                     ExpectedMySQLData,
		ExpectedUpdatedData:              ExpectedUpdatedData,
		DestinationDataTypeSchema:        MySQLToDestinationSchema,
		UpdatedDestinationDataTypeSchema: EvolvedMySQLToDestinationSchema,
		ExecuteQuery:                     ExecuteQuery,
		DestinationDB:                    "mysql_olake_mysql_test",
		CursorField:                      "id_cursor:id_smallint",
		PartitionRegex:                   "/{id,identity}",
		// Filter configuration: keeps records where created_timestamp >= 2023-01-01 AND id_int > 0
		FilterInput: &types.FilterInput{
			LogicalOperator: "AND",
			Conditions: []types.FilterCondition{
				{Column: "created_timestamp", Operator: ">=", Value: "2023-01-01T12:00:00"},
				{Column: "id_int", Operator: ">", Value: 0},
			},
		},
	}
	testConfig.TestIntegration(t)
}

func TestMySQLPerformance(t *testing.T) {
	config := &testutils.PerformanceTest{
		TestConfig:      testutils.GetTestConfig(string(constants.MySQL)),
		Namespace:       "benchmark",
		BackfillStreams: []string{"trips", "fhv_trips"},
		CDCStreams:      []string{"trips_cdc", "fhv_trips_cdc"},
		ExecuteQuery:    ExecuteQuery,
	}

	config.TestPerformance(t)
}
