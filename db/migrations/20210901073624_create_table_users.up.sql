CREATE TABLE users (
    id INT(11) NOT NULL,
    name VARCHAR(255),
    username VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL,
    password VARCHAR(150),
    address TEXT,
    phone_number VARCHAR(45),
    status_active INT(1) NOT NULL,
    is_partner INT(1) NOT NULL,
    profile_picture TEXT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL
);

ALTER TABLE users
ADD PRIMARY KEY (id);

ALTER TABLE users
MODIFY id int(11) NOT NULL AUTO_INCREMENT;