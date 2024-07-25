make-migration :
	migrate create -ext sql -dir migration -seq $(name)