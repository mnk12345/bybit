BYBIT_TEST_UPDATED=false

test:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/mnk12345/bybit/v2/integrationtest/v5/user -v -tags=integrationtestv5user

test-spec:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/mnk12345/bybit/v2/integrationtest/v5/user -v -tags=integrationtestv5user -run ${BYBIT_TEST_METHOD}

