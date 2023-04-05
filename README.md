# DTS BATCH 7

## DEMO

[![demo](https://img.youtube.com/vi/isFsEDRgGvE/0.jpg)](https://www.youtube.com/watch?v=isFsEDRgGvE)

## DOCS

### Setup Project

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/lib/pq
```

### Migration Sql

```sql
CREATE TABLE users(
    id serial NOT NULL PRIMARY KEY,
    email text UNIQUE,
    name VARCHAR(255),
    dob DATE,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz
);

-- memiliki relasi dengan users table
-- dihubungkan dengan user_id
CREATE TABLE user_photos(
    id serial NOT NULL PRIMARY KEY,
    url text,
    user_id int,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz,
    Foreign Key (user_id)
        REFERENCES users(id)
);

CREATE TABLE user_photos_no_fk(
    id serial NOT NULL PRIMARY KEY,
    url text,
    user_id int,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz
);

-- Insert dummy data into the users table
INSERT INTO users (email, name, dob) VALUES
    ('john.doe@example.com', 'John Doe', '1980-01-01'),
    ('jane.doe@example.com', 'Jane Doe', '1990-01-01'),
    ('bob.smith@example.com', 'Bob Smith', '2000-01-01');

-- Insert dummy data into the user_photos table
INSERT INTO user_photos (url, user_id) VALUES
    ('https://example.com/user1.jpg', 1),
    ('https://example.com/user2.jpg', 2),
    ('https://example.com/user3.jpg', 3);

SELECT * FROM users

CREATE TABLE books(
    id serial NOT NULL PRIMARY KEY,
    name_book text,
    author VARCHAR(255),
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz
);

INSERT INTO books(name_book, author)
VALUES ('The Great Gatsby', 'F. Scott Fitzgerald'),
       ('To Kill a Mockingbird', 'Harper Lee'),
       ('1984', 'George Orwell');
INSERT INTO books(name_book, author)
VALUES ('The Catcher in the Rye', 'J.D. Salinger'),
       ('Pride and Prejudice', 'Jane Austen'),
       ('The Lord of the Rings', 'J.R.R. Tolkien');

SELECT id, name_book, author FROM books u
```
