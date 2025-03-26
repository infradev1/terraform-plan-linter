PLAN_OUT=tfplan.binary
PLAN_JSON=tf-plan.json

#.PHONY: plan lint test clean

plan:
	cd terraform && terraform init -input=false
	cd terraform && terraform plan -out=$(PLAN_OUT)
	cat terraform/$(PLAN_OUT)
	terraform show -json terraform/$(PLAN_OUT) | tee terraform/$(PLAN_JSON) > /dev/null
	sleep 1 && cat terraform/$(PLAN_JSON)

lint:
	go run main.go --file terraform/$(PLAN_JSON)

test:
	go test ./... -v

clean:
	rm -f terraform/$(PLAN_JSON)
	rm -f terraform/$(PLAN_OUT)
