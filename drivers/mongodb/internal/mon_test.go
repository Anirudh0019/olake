package driver

import (
	"testing"

	"github.com/datazip-inc/olake/constants"
	"github.com/datazip-inc/olake/types"
	"github.com/datazip-inc/olake/utils/testutils"
)

func TestMongodbIntegration(t *testing.T) {
	t.Parallel()
	testConfig := &testutils.IntegrationTest{
		TestConfig:                       testutils.GetTestConfig(string(constants.MongoDB)),
		Namespace:                        "olake_mongodb_test",
		ExpectedData:                     ExpectedMongoData,
		ExpectedUpdatedData:              ExpectedUpdatedData,
		DestinationDataTypeSchema:        MongoToDestinationSchema,
		UpdatedDestinationDataTypeSchema: UpdatedMongoToDestinationSchema,
		ExecuteQuery:                     ExecuteQuery,
		DestinationDB:                    "mongodb_olake_mongodb_test",
		CursorField:                      "id_cursor:id_int",
		PartitionRegex:                   "/{_id,identity}",
		// Filter configuration: keeps records where id_timestamp >= 2023-01-01 AND id_int > 0
		FilterInput: &types.FilterInput{
			LogicalOperator: "AND",
			Conditions: []types.FilterCondition{
				{Column: "id_timestamp", Operator: ">=", Value: "2023-01-01T12:00:00"},
				{Column: "id_int", Operator: ">", Value: 0},
			},
		},
	}
	testConfig.TestIntegration(t)
}

func TestMongodbPerformance(t *testing.T) {
	config := &testutils.PerformanceTest{
		TestConfig:      testutils.GetTestConfig(string(constants.MongoDB)),
		Namespace:       "twitter_data",
		BackfillStreams: []string{"tweets"},
		CDCStreams:      []string{"tweets_cdc"},
		ExecuteQuery:    ExecuteQuery,
	}

	config.TestPerformance(t)
}
