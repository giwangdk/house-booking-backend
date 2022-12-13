CREATE DATABASE house_booking_giwang;

CREATE TABLE cities (
   	id BIGSERIAL PRIMARY KEY,
	name varchar(50) NOT NULL,
   	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
  	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
   	deleted_at TIMESTAMPTZ
);

CREATE TABLE categories (
   	id BIGSERIAL PRIMARY KEY,
	name varchar(50) NOT NULL,
   	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
  	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
   	deleted_at TIMESTAMPTZ
);

CREATE TABLE users (
   	id BIGSERIAL PRIMARY KEY,
   	email VARCHAR(50) UNIQUE NOT NULL,
	fullname varchar(50) NOT NULL,
   	password VARCHAR(225) NOT NULL,
    address VARCHAR(100),
    role VARCHAR(20) NOT NULL,
  	city_id INT,
   	FOREIGN KEY (city_id)
     REFERENCES cities (id),
   	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
  	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
   	deleted_at TIMESTAMPTZ
);

CREATE TABLE games(
    id BIGSERIAL PRIMARY KEY,
    user_id INT,
    chance INT,
    FOREIGN KEY (user_id)
     REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE houses (
   	id BIGSERIAL PRIMARY KEY,
    name varchar(50) NOT NULL,
   	city_id INT,
   	category_id INT,
    user_id INT,
   	price INT,
    description VARCHAR(100),
    max_guest INT NOT NULL,
   	FOREIGN KEY (city_id)
     REFERENCES cities (id),
   	FOREIGN KEY (user_id)
     REFERENCES users (id),
   	FOREIGN KEY (category_id)
     REFERENCES categories (id),
   	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
  	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
   	deleted_at TIMESTAMPTZ
);

CREATE TABLE houses_photos(
    id BIGSERIAL PRIMARY KEY,
    house_id INT,
    photo VARCHAR(100),
    FOREIGN KEY (house_id)
     REFERENCES houses (id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE reservations(
    id BIGSERIAL PRIMARY KEY,
    house_id INT,
    user_id INT,
    check_in DATE,
    check_out DATE,
    total_price INT,
    FOREIGN KEY (house_id)
     REFERENCES houses (id),
    FOREIGN KEY (user_id)
     REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE transactions(
    id BIGSERIAL PRIMARY KEY,
    reservation_id INT,
    user_id INT,
    house_id INT,
    payment_method VARCHAR(50),
    FOREIGN KEY (reservation_id)
     REFERENCES reservations (id),
    FOREIGN KEY (user_id)
     REFERENCES users (id),
    FOREIGN KEY (hosue_id)
     REFERENCES houses (id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE pickup_status(
    id BIGSERIAL PRIMARY KEY,
    status VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE pickups(
    id BIGSERIAL PRIMARY KEY,
    reservation_id INT,
    user_id INT,
    pickup_status_id INT,
    FOREIGN KEY (reservation_id)
     REFERENCES reservations (id),
    FOREIGN KEY (user_id)
     REFERENCES users (id),
    FOREIGN KEY (pickup_status_id)
     REFERENCES pickup_status (id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);



