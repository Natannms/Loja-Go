CREATE TABLE produtos(  
    id SERIAL NOT NULL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    preco decimal(10,2) NOT NULL,
    quantidade INTEGER NOT NULL
);