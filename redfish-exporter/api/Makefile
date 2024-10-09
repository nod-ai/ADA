.PHONY: gen-slurm-rest-api gen-clean clean

gen-slurm-rest-api:
	docker run --rm --user $(shell id -u):$(shell id -g) -v ${PWD}:/api openapitools/openapi-generator-cli generate -i /api/slurm-openapi.json -g go -o /api/generated/slurmrestdapi --package-name slurmrestdapi --additional-properties=withGoMod=false --global-property apiTests=false,modelTests=false
gen-clean:
	rm -rf generated/*

clean: gen-clean
