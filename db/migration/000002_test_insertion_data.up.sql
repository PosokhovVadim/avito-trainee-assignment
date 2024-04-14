INSERT INTO feature DEFAULT VALUES;
INSERT INTO feature DEFAULT VALUES;
INSERT INTO feature DEFAULT VALUES;
INSERT INTO feature DEFAULT VALUES;
INSERT INTO feature DEFAULT VALUES;

INSERT INTO tag DEFAULT VALUES;
INSERT INTO tag DEFAULT VALUES;
INSERT INTO tag DEFAULT VALUES;
INSERT INTO tag DEFAULT VALUES;
INSERT INTO tag DEFAULT VALUES;



INSERT INTO banner (content, is_active) VALUES ('{"title": "Banner 1", "message": "banner 1 content"}', true);
INSERT INTO banner (content, is_active) VALUES ('{"title": "Banner 2", "message": "banner 2 content"}', true);
INSERT INTO banner (content, is_active) VALUES ('{"title": "Banner 3", "message": "banner 3 content"}', true);

INSERT INTO banner_feature (banner_id, feature_id) VALUES (1, 1), (1, 2), (2, 2), (3, 3);

INSERT INTO banner_tag (banner_id, tag_id) VALUES (1, 1), (1, 2), (2, 2), (2, 3), (3, 4), (3, 5);
