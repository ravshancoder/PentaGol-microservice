INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/swagger/*', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/login', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/posts', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/post/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/liga/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/ligas', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/game/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/games', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/posts', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/post/{id}', 'DELETE');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/admin/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/liga', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/liga/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/ligas', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/liga/{id}', 'DELETE');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/post/{id}', 'PUT');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/post/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/post', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/game', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/game/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/games', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/game/{id}', 'DELETE');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/club', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/club/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/clubs', 'GET');


INSERT INTO admins(name, email, password, refresh_token) VALUES('ravshan', 'mavlonovr555@gmail.com', '$2a$14$synOItEs8oaSXPgA13wFpetq1pJZBNf9mqTvd4XyIhsZy5R98v0UG', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.exw43QM2_SlQ_MC3zIZ1XyUvnOldQy7CxqRYsvUI4A0');
