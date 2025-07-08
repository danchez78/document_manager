CREATE SCHEMA document_manager;

CREATE TABLE document_manager.users
(
    id VARCHAR(255) NOT NULL,
    login VARCHAR NOT NULL,
    hashed_password VARCHAR NOT NULL,
    token VARCHAR NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(login)
);

CREATE TABLE document_manager.docs
(
    id VARCHAR(255) NOT NULL,
    data bytea NOT NULL,
    name VARCHAR NOT NULL,
    file bool NOT NULL,
    public bool NOT NULL,
    mime VARCHAR NOT NULL,
    created VARCHAR NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE document_manager.grants
(
    doc_id VARCHAR(255) NOT NULL,
    user_login VARCHAR NOT NULL,
    CONSTRAINT PK_Grant PRIMARY KEY (doc_id, user_login)
);

ALTER TABLE document_manager.grants
    ADD FOREIGN KEY (doc_id) REFERENCES document_manager.docs (id)
        ON UPDATE NO ACTION ON DELETE CASCADE
    ADD FOREIGN KEY (user_login) REFERENCES document_manager.users (login)
        ON UPDATE NO ACTION ON DELETE CASCADE;
