package driver

import (
	"testing"

	"github.com/datazip-inc/olake/constants"
	"github.com/datazip-inc/olake/types"
	"github.com/datazip-inc/olake/utils/testutils"
)

func TestOracleIntegration(t *testing.T) {
	t.Parallel()
	testConfig := &testutils.IntegrationTest{
		TestConfig:                       testutils.GetTestConfig(string(constants.Oracle)),
		Namespace:                        "MYUSER",
		ExpectedData:                     ExpectedOracleData,
		ExpectedUpdatedData:              ExpectedUpdatedOracleData,
		DestinationDataTypeSchema:        OracleToDestinationSchema,
		UpdatedDestinationDataTypeSchema: UpdatedOracleToDestinationSchema,
		ExecuteQuery:                     ExecuteQuery,
		DestinationDB:                    "oracle_myuser",
		CursorField:                      "COL_CURSOR:COL_SMALLINT",
		PartitionRegex:                   "/{id, identity}",
		// Filter configuration: keeps records where col_timestamp >= 2023-01-01 AND col_int > 0
		FilterInput: &types.FilterInput{
			LogicalOperator: "AND",
			Conditions: []types.FilterCondition{
				{Column: "COL_TIMESTAMP", Operator: ">=", Value: "2023-01-01T12:00:00"},
				{Column: "COL_INT", Operator: ">", Value: 0},
			},
		},
	}
	testConfig.TestIntegration(t)
}
