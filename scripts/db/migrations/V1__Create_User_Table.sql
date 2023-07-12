CREATE TABLE Users (
    userId UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(60) NOT NULL,
    phone CHAR(13),
    region VARCHAR(100),
    gender ENUM('M', 'F', 'N') DEFAULT 'N',
    title_id INTEGER NOT NULL,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    age TINYINT,
    CONSTRAINT fk_title FOREIGN KEY (title_id) REFERENCES Titles(title_id)
);

CREATE TABLE Titles (
    title_id SERIAL PRIMARY KEY,
    title VARCHAR(10) NOT NULL
);
