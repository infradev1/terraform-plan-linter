PLAN_OUT=tfplan.binary
PLAN_JSON=testdata/tf-plan.json

.PHONY: plan lint test clean

plan:
	cd terraform
	terraform init
	terraform plan -out=$(PLAN_OUT)
	terraform show $(PLAN_OUT)

lint:
	go run main.go --file $(PLAN_JSON)

test:
	go test ./... -v

clean:
	rm -f $(PLAN_JSON)
	rm -f terraform/$(PLAN_OUT)
