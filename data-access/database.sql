DROP TABLE IF EXISTS message;
DROP TABLE IF EXISTS contact;
DROP TABLE IF EXISTS conversation;
DROP TABLE IF EXISTS group_member;

CREATE TABLE message
(
    message_id SERIAL PRIMARY KEY,
    from_number INTEGER varchar(10),
    to_number INTEGER varchar(10),
    message_text NOT NULL TEXT,
    sent_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    conversation_id FOREIGN KEY REFERENCES conversation(conversation_id)
);

CREATE TABLE contact
(
    contact_id SERIAL PRIMARY KEY,
    first_name NOT NULL TEXT,
    last_name NOT NULL TEXT,
    -- profile_photo,
    phone_number INTEGER varchar(10)
);

CREATE TABLE conversation
(
    conversation_id SERIAL PRIMARY KEY,
    conversation_name TEXT NOT NULL DEFAULT "Empty",
    messageId_list TEXT[]
);

CREATE TABLE group_member
(
    contact_id FOREIGN KEY REFERENCES contact(contact_id),
    conversation_id FOREIGN KEY REFERENCES conversation(conversation_id),
    joined_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    left_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
