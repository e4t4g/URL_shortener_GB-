
CREATE TABLE IF NOT EXISTS url (
                                   id INTEGER PRIMARY KEY AUTOINCREMENT,
                                   full_url TEXT NOT NULL,
                                   short_url TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS stat (
                                    url_id INTEGER NOT NULL,
                                    counter INTEGER,
                                    FOREIGN KEY (url_id) REFERENCES url(id)
);

