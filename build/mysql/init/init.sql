CREATE TABLE pokemons (
    id INT PRIMARY KEY AUTO_INCREMENT,
    pokemon_number INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    type1 VARCHAR(50) NOT NULL,
    type2 VARCHAR(50),
    height DECIMAL(20, 10) NOT NULL,
    weight DECIMAL(20, 10) NOT NULL,
    base_experience INT NOT NULL,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE parties (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    pokemon_id INT NOT NULL,
    nickname VARCHAR(255),
    level INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE exchanges (
    id INT PRIMARY KEY AUTO_INCREMENT,
    sender_user_id INT NOT NULL,
    receiver_user_id INT NOT NULL,
    sender_party_id INT NOT NULL,
    receiver_party_id INT NOT NULL,
    exchange_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);