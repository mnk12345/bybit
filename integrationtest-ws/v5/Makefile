BYBIT_TEST_UPDATED=false

test:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/mnk12345/bybit/v2/integrationtest-ws/v5 -v -tags=integrationtestwsv5

test-spec:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/mnk12345/bybit/v2/integrationtest-ws/v5 -v -tags=integrationtestwsv5 -run ${BYBIT_TEST_METHOD}

