
### Ver registros logeados
```bash
cat ~/.docker/config.json
```

### Crear un repo en ecr
```bash
aws ecr create-repository \
    --repository-name odihnx-poc-grpc-app-3 \
    --region us-east-1 \
    --profile odihnx-dev
```

### Login en ECR desde Docker
```bash
aws ecr get-login-password --region us-east-1 --profile odihnx-dev | \
docker login --username AWS --password-stdin 590184058323.dkr.ecr.us-east-1.amazonaws.com
```

### Construir imagen
```bash
docker build -t poc-grpc-app-3 .
```

### Etiquetar (tag) la imagen para ECR 
```bash
docker tag poc-grpc-app-3:latest 590184058323.dkr.ecr.us-east-1.amazonaws.com/odihnx-poc-grpc-app-3:latest
```

### Subir imagen
```bash
docker push 590184058323.dkr.ecr.us-east-1.amazonaws.com/odihnx-poc-grpc-app-3:latest
```