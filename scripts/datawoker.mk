# sample:
# make run-dataworker query="select\\ 55\\ AS\\ data1,2\\ AS\\ data2,1\\ AS\\ keyfield1,2\\ AS\\ keyfield2,3\\ AS\\ keyfield3\\ from\\ campaigns\\ where\\ id=1\\;" keyfields="campaign_id,hour_id,daily_id" targetfields="cpc,cpm" targetTable=campaign_detail

run-dataworker:
	$(BUILD) clickyab.com/crab/cmd/dataworker
	$(BIN)/dataworker --query=$(query) --keyfields=$(keyfields) --targetfields=$(targetfields) --targetTable=$(targetTable)

test-dataworker:
	$(BUILD) clickyab.com/crab/cmd/dataworker
	$(BIN)/dataworker --query="select 55 AS data1,2 AS data2,1 AS keyfield1,2 AS keyfield2,3 AS keyfield3 from campaigns where id=1;" --keyfields="campaign_id,hour_id,daily_id" --targetfields="cpc,cpm" --targetTable="campaign_detail"