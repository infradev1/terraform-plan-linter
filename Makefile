PLAN_OUT=tfplan.binary
PLAN_JSON=tf-plan.json

#.PHONY: plan lint test clean

plan:
	cd terraform && terraform init -input=false
	cd terraform && terraform plan -out=$(PLAN_OUT) && terraform show -json $(PLAN_OUT) > $(PLAN_JSON)
	sleep 1
	cat terraform/$(PLAN_JSON)
	sleep 1

lint:
	go run main.go --file terraform/$(PLAN_JSON)

test:
	go test ./... -v

clean:
	rm -f terraform/$(PLAN_JSON)
	rm -f terraform/$(PLAN_OUT)
