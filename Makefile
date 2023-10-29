.PHONY: buildLinux run dockerrun

buildLinux:
	./buildScripts/buildLinux.sh
	docker build -t player-service .

run:
	./buildScripts/run.sh

dockerRun: buildLinux
	docker run -e PORT="8080" -e CSV_FILE_PATH="/opt/players.csv" -p 8080:8080 player-service
