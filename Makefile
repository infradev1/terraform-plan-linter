PLAN_OUT=tfplan.binary
PLAN_JSON=tf-plan.json

#.PHONY: plan lint test clean

plan:
	cd terraform && terraform init
	cd terraform && terraform plan -out=$(PLAN_OUT)
	cd terraform && cat $(PLAN_OUT)
	terraform show -json terraform/$(PLAN_OUT) > $(PLAN_JSON)
	sleep 1 && cat $(PLAN_JSON)

lint:
	go run main.go --file $(PLAN_JSON)

test:
	go test ./... -v

clean:
	rm -f terraform/$(PLAN_JSON)
	rm -f terraform/$(PLAN_OUT)
