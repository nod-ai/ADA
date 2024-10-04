.PHONY: gen-slurm-rest-api gen-clean clean

gen-slurm-rest-api:
	docker run --rm --user $(shell id -u):$(shell id -g) -v ${PWD}:/api openapitools/openapi-generator-cli generate -i /api/slurm-openapi.json -g go -o /api/generated/slurmrestdapi --package-name slurmrestdapi
	rm -f generated/slurmrestdapi/go.sum
	rm -f generated/slurmrestdapi/go.mod

gen-clean:
	rm -rf generated/*

clean: gen-clean
