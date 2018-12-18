image:
	docker build -t cirocosta/prober .

push: image
	docker push cirocosta/prober

k8s:
	kubectl apply --namespace=dev -f ./prober-deployment.yaml
