init:
	docker compose down
	docker compose up -d
	go run . import-deck /decks/bmc_1.json
	go run . import-deck /decks/bmc_2.json
	go run . import-deck /decks/cao.json