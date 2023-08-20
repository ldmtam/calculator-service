gen-proto:
	buf generate --path proto

docker-minikube:
	eval $(minikube docker-env)

build-client:
	docker-compose build calculator-client

build-server:
	docker-compose build calculator-server

deploy-server:
	kubectl apply -f deployment/calculator-server.yaml
	kubectl rollout restart deployment calculator-server

deploy-client:
	kubectl apply -f deployment/calculator-client.yaml
	kubectl rollout restart deployment calculator-client

allow-watch:
	kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default