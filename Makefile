ifdef RUN
	RUNFUNC := -run $(RUN)
endif

deploy:
	gcloud app deploy app.yaml --version 1
