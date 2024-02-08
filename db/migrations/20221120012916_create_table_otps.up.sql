CREATE TABLE otp (
    id INT(11) NOT NULL,
    user_id INT(11) NOT NULL,
    code CHAR(4) NOT NULL,
    category VARCHAR(30) NOT NULL,
    expired TINYINT NOT NULL,
    timer_otp int(1) NULL,
    indentity varchar(255) NULL,
    platform varchar(255) NULL,
    used varchar(2) NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL
);

ALTER TABLE otp
ADD PRIMARY KEY (id);

ALTER TABLE otp
MODIFY id int(11) NOT NULL AUTO_INCREMENT;