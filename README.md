# Monitoramento de sites

## Objetivo

Aprimorar conhecimentos básicos de Go com foco em automação e Devops, neste caso, foi feita uma CLI básica para monitorar sites.

## Como rodar

Para rodar esta aplicação, será preciso ter docker instalado em sua máquinac, para isso, entre [aqui](https://docs.docker.com/get-docker/). Tendo docker, basta executar os comandos abaixo e usar a aplicação como desejar.

Primeiro, clone o repositório com o seguinte comando:
```bash
$ git clone https://github.com/CrysLef/site-monitoramento.git 
```
Após isso, faça o build da imagem e rode o container, executando os seguintes comando em sequência:
```bash
$ docker build -t site-monitoramento .
```

```bash
$ docker container run -it site-monitoramento
```

Caso tenha instalado em sua máquina Golang ou prefira rodar diretamento por ele, primeiro instale o Golang clicando [aqui](https://go.dev/dl/) e rode este código:

```bash
go run cmd/main.go
```