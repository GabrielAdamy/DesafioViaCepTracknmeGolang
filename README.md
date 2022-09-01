# DesafioViaCepTracknmeGolang
<h1 align="center">
  <br>
    <a href="https://github.com/GabrielAdamy/DesafioViaCepTracknmeGolang"><img src="https://media-exp1.licdn.com/dms/image/C4E0BAQFrhiNubVpL_A/company-logo_200_200/0/1635518084901?e=2147483647&v=beta&t=YqhegBy3yBJsH4_0Q-odEiInleCGhYlgDZrg21tYWis"</a>
  <br>
  Desafio crud em golang 
  <br>
</h1>

<p align="left">
   <p>
   ENDPOINTS
   <br>
   <br>
POST
http://localhost:8080/employee/create
<br>
Body: <br> { <br>"name": String, (OBRIGATÓRIO) <br>"age": int,(OBRIGATÓRIO) <br>"sex": String, (OBRIGATÓRIO) <br>"cep": String (OBRIGATÓRIO)<br> }
<br>
<br>
GET
http://localhost:8080/employee
<br>
<br>
GET
http://localhost:8080/employee/cep/{cep String}
<br>
<br>
GET
http://localhost:8080/viacep/cep/{cep String}
<br>
<br>
DELETE
http://localhost:8080/employee/{id}
<br>
<br>
PUT
http://localhost:8080/employee/update/{id}
<br>
Body: <br>{ <br> "name": String, (OBRIGATÓRIO) <br> "age": int, (OBRIGATÓRIO) <br> "sex": String, <br> "cep": String, (OBRIGATÓRIO) <br> "address": String, <br> "district": String, <br> "city": String, <br>"state": String  <br>}
<br>
<br>
PATCH
http://localhost:8080/employee/{id}
<br>
Body:<br> { <br>"name": String, (OBRIGATÓRIO) <br>"age": int, (OBRIGATÓRIO) <br>"sex": String, <br>"cep": String (OBRIGATÓRIO)<br> }</p>
</p>

<h6 align="center">

  👤 **Gabriel Adamy**

Github:   [@GabrielAdamy](https://github.com/GabrielAdamy)
LinkedIn: [@gabriel-adamy](https://www.linkedin.com/in/gabriel-adamy) 
</h6>
