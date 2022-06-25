DELETE FROM todos;
INSERT INTO todos (`id`, `text`, `done`)
VALUES
  ("test_todo_id_001", "test todo 001", 0),
  ("test_todo_id_002", "test todo 002", 0),
  ("test_todo_id_003", "test todo 003", 1)
;