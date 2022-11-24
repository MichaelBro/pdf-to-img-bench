up:
	docker compose up -d --build
down:
	docker compose down --remove-orphans
run:
	docker compose exec app sh ./bench.sh
