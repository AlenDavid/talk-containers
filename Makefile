container-app:
	docker build -t talks/bakery-app ./services/bakery-app
container-web:
	docker build -t talks/bakery-web ./services/bakery-web

containers: container-app container-web

publish-app:
	docker tag talks/bakery-app registry.heroku.com/bakery-app/web
	docker push registry.heroku.com/bakery-app/web
publish-web:
	docker tag talks/bakery-web registry.heroku.com/bakery-web/web
	docker push registry.heroku.com/bakery-web/web

publish: publish-app publish-web

release:
	heroku container:release web -a bakery-app
	heroku container:release web -a bakery-web

build: containers publish release
