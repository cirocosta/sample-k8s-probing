image:
	docker build -t cirocosta/prober .

push: image
	docker push cirocosta/prober
