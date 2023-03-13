
.PHONY: examples
examples:
	# go run -mod=mod ../../.. --dev=docker://postgres/10 --dir=./schema > postgres.hcl
	# go run -mod=mod ../../.. --dev=docker://mysql/8/test --dir=./schema > mysql.hcl
	go run -mod=mod ../../.. --dev=sqlite3://.tmp.db?_fk=1 --dir=./schema > sqlite3.hcl
