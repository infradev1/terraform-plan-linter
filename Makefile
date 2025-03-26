PLAN_OUT=tfplan.binary
PLAN_JSON=testdata/sample-plan.json

.PHONY: plan lint test clean

plan:
	cd terraform-sample && terraform init -input=false
	cd terraform-sample && terraform plan -out=$(PLAN_OUT)
	cd terraform-sample && terraform show -json $(PLAN_OUT) > ../$(PLAN_JSON)

lint:
	go run main.go --file $(PLAN_JSON)

test:
	go test ./... -v

clean:
	rm -f $(PLAN_JSON)
	rm -f terraform-sample/$(PLAN_OUT)
