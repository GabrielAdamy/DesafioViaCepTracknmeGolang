<h1 align="center"> DesafioViaCepTracknmeGolang</h1>
<h1 align="center">
  <br>
    <a href="https://github.com/GabrielAdamy/DesafioViaCepTracknmeGolang"><img src="https://media-exp1.licdn.com/dms/image/C4E0BAQFrhiNubVpL_A/company-logo_200_200/0/1635518084901?e=2147483647&v=beta&t=YqhegBy3yBJsH4_0Q-odEiInleCGhYlgDZrg21tYWis"</a>
 <br> Desafio crud em golang
</h1>
  
  <p>O projeto tem como finalidade testar o conhecimento e evolução do desenvolvedor. No projeto foi utilizado a conexão com o banco de dados Postgres, utilizando os métodos básicos de um crud.</p>

  <h2 aling="left">Conexão com o Postgres</h2>
  <br>
  <p>    
  "POSTGRES_HOST": "localhost", <br>
  "POSTGRES_PORT": "5432", <br>
  "POSTGRES_USER": "postgres", <br>
  "POSTGRES_PASSWORD": "postgres", <br>
  "POSTGRES_DBNAME": "tracknme" <br>
  </p>
  
<h3 align="left">  
   ENDPOINTS:
</h3>
<h5 align="left">  
POST <br>
http://localhost:8080/api/employee/create

Body: { <br>"name": String, (OBRIGATÓRIO) <br>"age": int,(OBRIGATÓRIO) <br>"sex": String, (OBRIGATÓRIO) <br>"cep": String (OBRIGATÓRIO)<br> }
</h5>
<h5 align="left">
GET <br>
http://localhost:8080/api/employee
</h5>
<h5 align="left">
GET <br>
http://localhost:8080/api/employee/cep/{cep String}
</h5>
<h5 align="left">
GET <br>
http://localhost:8080/api/viacep/cep/{cep String}
</h5>
<h5 align="left">
DELETE <br>
http://localhost:8080/api/employee/{id}
</h5>
<h5 align="left">
PUT <br>
http://localhost:8080/api/employee/update/{id}
<br>
Body: { <br> "name": String, (OBRIGATÓRIO) <br> "age": int, (OBRIGATÓRIO) <br> "sex": String, <br> "cep": String, (OBRIGATÓRIO) <br> "address": String, <br> "district": String, <br> "city": String, <br>"state": String  <br>}
</h5>
<h5 align="left">
PATCH <br>
http://localhost:8080/api/employee/{id}
<br>
Body: { <br>"name": String, (OBRIGATÓRIO) <br>"age": int, (OBRIGATÓRIO) <br>"sex": String, <br>"cep": String (OBRIGATÓRIO)<br> }
</h5>
</p>

<h6 align="center">
  👤 ***Gabriel Adamy***
</h6>
<h6 align="center">
GitHub: https://github.com/GabrielAdamy
  
LinkedIn: https://www.linkedin.com/in/gabriel-adamy
</h6>
