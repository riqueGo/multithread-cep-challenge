# Multithreading

## Challenge
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.
As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://cdn.apicep.com/file/apicep/" + cep + ".json

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:
- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- O resultado da request deverá ser exibido no command line, bem como qual API a enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## How to Run
<code>go run main.go</code>

Faça uma requisição GET na porta 8080 colocando no path o CEP desejado.</br>
Dentre as duas API's o serviço retornará a resposta mais rápida.</br>
A resposta será disponibilizada no command line.</br>
No cabeçalho da mensagem conterá a API responsável pela resposta, e no corpo o retorno dela.</br>

### Exemplo

<code>curl http://localhost:8080/30250-130</code>

