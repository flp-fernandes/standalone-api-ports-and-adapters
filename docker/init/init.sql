-- Cria tabela de usuários
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

-- Cria tabela de pedidos
CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  amount NUMERIC NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Insere dados de exemplo
INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Carol');

INSERT INTO orders (user_id, amount) VALUES
  (1, 100.00),
  (2, 50.50),
  (3, 250.75);
