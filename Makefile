init:
	docker compose down
	docker compose up -d
	sleep 5
	go run . deck import ./decks/fixture_deck.json
	go run . deck import ./decks/fixture_deck_2.json