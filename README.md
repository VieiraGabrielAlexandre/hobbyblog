# 🧙‍♂️ **HobbyBlog – API e Blog Geek**

### *Seu universo de Anime, Mangá, Filmes, HQs e Escape Rooms em um só lugar.*

<p align="center">
  <img src="assets/blog-capa.png" width="650" />
</p>

<p align="center">
  <i>“Um blog para quem vive entre páginas de mangás, salas secretas de escape rooms, cinemas escuros e mundos fantásticos.”</i>
</p>

---

## 🚀 **O que é o HobbyBlog?**

O **HobbyBlog** é uma API + backend construída em **Go (Golang)**, usando **Uber Fx** para DI e ciclo de vida, e **Gin** como framework web.
Ele foi pensado para armazenar posts sobre:

* 🎬 Filmes
* 📺 Séries & Animes
* 📚 Mangás & HQs
* 🕹️ Games
* 🧩 Escape Rooms
* 🎨 Obras e conteúdos gerais

E cada post pode conter **HTML**, **Markdown**, **metadados**, imagens, notas, autores, etc.

---

## 🌌 **Por que esse projeto existe?**

Porque um blog sobre cultura pop **não precisa ser genérico**.
A ideia aqui é ter um backend limpo, moderno e extensível, para futuramente:

* Um blog estático
* Uma dashboard de publicação
* Um app mobile
* Uma wiki geek personalizável

Tudo isso com uma base sólida de API.

---

## 🛠️ **Tecnologias**

| Categoria          | Stack                                          |
| ------------------ | ---------------------------------------------- |
| Linguagem          | Go 1.22+                                       |
| Server             | Gin                                            |
| DI / Ciclo de vida | Uber Fx                                        |
| Logs               | Zap                                            |
| Data               | DynamoDB (em produção) / Repo em Memória (dev) |
| Infra              | AWS Lambda / API Gateway (futuro)              |
| Spec               | OpenAPI 3.0                                    |

---

## 🧩 **Arquitetura (Atual – MVP)**

```
hobbyblog/
  cmd/api/main.go
  internal/app/module.go
  internal/config/config.go
  internal/log/logger.go
  internal/http/router.go
  internal/server/httpserver.go

  internal/health/
      module.go

  internal/posts/
      module.go
      handler.go
      repo.go
      dto.go
      slug.go
```

* **Uber Fx** gerencia todas as dependências.
* **Gin** é o router principal.
* **Repo em memória** para dev.
* Rotas de `/health` e `/v1/posts` já disponíveis.

---

## 🧪 **Endpoints atuais**

### 🔍 Healthcheck

```
GET /health
```

### ✍️ Criar post sobre anime/filme/mangá/etc.

```
POST /v1/posts
Content-Type: application/json
```

Exemplo de corpo:

```json
{
  "type": "anime",
  "title": "Akira",
  "slug": "Akira Neo-Tokyo",
  "contentFormat": "html",
  "content": "<h1>Kaneda! Tetsuo!</h1><p>Neo-Tokyo é insana.</p>",
  "tags": ["cyberpunk","classic"],
  "rating": 9.5
}
```

### 🔎 Buscar por ID

```
GET /v1/posts/{id}
```

### 🔎 Buscar por Slug

```
GET /v1/posts/slug/{slug}
```

---

## 🧭 **Como rodar o projeto**

### 🔧 Pré-requisitos

* Go 1.22+
* Git

### ▶️ Rodar local

```bash
go mod tidy
PORT=8080 go run ./cmd/api
```

Abra:

```
http://localhost:8080/health
```

---

## 🐉 **Imagens Temáticas**

Anime vibes para decorar seu README e ilustrar o espírito do HobbyBlog:

<p align="center">
  <img src="assets/akira-capa.jpg" width="500" />
  <br>
  <i>Neo-Tokyo está prestes a explodir.</i>
</p>

<p align="center">
  <img src="assets/mk-shang-tsung-your-soul-is-mine.gif" width="500" />
  <br>
  <i>Tem jogos também, resenha.</i>
</p>

<p align="center">
  <img src="assets/superman-hq.jpg" width="500" />
  <br>
  <i>Mangás, HQs e mundos fantásticos.</i>
</p>

---

## 📘 **OpenAPI (Spec)**

O projeto mantém um arquivo:

```
openapi.yaml
```

Já documentando:

* criação de posts
* healthcheck
* buscas por ID/slug
* metadados extensíveis

---

## 🧱 **Roadmap**

### ✔️ MVP Atual

* [x] Healthcheck
* [x] Módulo de posts
* [x] Repo em memória
* [x] Sanitização HTML
* [x] Slug normalizado
* [x] Uber Fx configurado

### 🔜 Próximos passos

* [ ] Listagem com paginação
* [ ] Update / Delete / Publish
* [ ] Implementar repositório DynamoDB
* [ ] CI/CD (GitHub Actions)
* [ ] Deploy na AWS (Lambda + API Gateway)
* [ ] Painel de administração (S3 + CloudFront)
* [ ] Blog estático com frontend simples
* [ ] Sistema de rating/comentários
* [ ] Suporte a uploads (imagens, vídeos, etc.)
* [ ] Tags/Busca inteligente

---

## 🎮 **Sobre o Dev**

Criado por um dev geek, viciado em:

* 🧩 Escape Rooms
* 🍜 Cultura japonesa
* 📚 Mangás de ação
* 🎬 Filmes clássicos
* 🕹️ Games retrô
* 👨‍💻 Arquitetura e código limpo

---

## 🧡 **Contribuição**

Ideias são bem-vindas!
Se quiser sugerir categorias, features, novos tipos de metadados ou padrões de postagem, abra uma issue ou mande mensagem.

---

## 🏆 **Licença**

MIT — livre para usar, modificar e criar sua própria versão.

---

Se quiser, posso também:

* gerar **badges animados**,
* criar uma **logo “HobbyBlog” estilo neon anime**,
* adicionar um **diagrama da arquitetura**,
* ou criar uma **cover** estilo banner para colocar no topo do README.
