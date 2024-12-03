COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/sure/blob/10a10cb5c07f6239fc8e310459fb6a240dd117ce/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
