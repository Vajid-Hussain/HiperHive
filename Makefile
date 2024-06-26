autoserver:
	nodemon --watch './api-gateway/**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go
	nodemon --watch './auth-service/**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go

deploy:
	docker compose up --detach

deploydown:
	docker compose down

dockerPostgres:
	docker exec -it hyperhive_postgres bash -c 'psql -U postgres'

swaggo:
	swag init --parseDependency --parseInternal --parseDepth 5  -g ./api-gateway/cmd/main.go -o api-gateway/

dockerall:
	docker compose down
	docker compose build
	docker compose up
	
uploadImgfriend:
	docker tag vajidhussain/hyperhive_friend_service:latest hyperhive_friend_service
	docker push vajidhussain/hyperhive_friend_service

uploadImgserver:
	docker tag vajidhussain/hyperhive_friend_service:latest hyperhive_friend_service
	docker push vajidhussain/hyperhive_friend_service


#cerr manager
# kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.7.1/cert-manager.yaml
# kubectl get pods --namespace cert-manager


# #ingress controller
# kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v<vnum>/deploy/static/provider/cloud/deploy.yaml
# kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.1/deploy/static/provider/cloud/deploy.yaml
# kubectl apply -f https://raw.githubusercontent.com/oracle-devrel/oci-oke-virtual-nodes/main/ingress-nginx/deploy.yaml
