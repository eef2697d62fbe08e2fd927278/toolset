-- insert test data to see if my code works

-- password is 'password' in SHA-256
INSERT INTO toolset.tbl_user (email, username, password) VALUES ("email@email.com", "username", "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8");
INSERT INTO toolset.tbl_user (email, username, password) VALUES ("email@email2.com", "username2", "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8");

INSERT INTO toolset.tbl_note (title, content, time, author) VALUES ("title1", "content1", "2020-12-09 12:59:59", 1);
INSERT INTO toolset.tbl_note (content, time, author) VALUES ("content2", "2020-12-09 12:59:59", 2);

INSERT INTO toolset.tbl_tag (name) VALUES ("tag1");
INSERT INTO toolset.tbl_tag (name) VALUES ("tag2");

INSERT INTO toolset.lktbl_tag (note, tag) VALUES (1, 1);
INSERT INTO toolset.lktbl_tag (note, tag) VALUES (2, 2);
