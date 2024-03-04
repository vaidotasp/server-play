import sqlite3 from "sqlite3";

export function initDB() {
  sqlite3.verbose();
  const db = new sqlite3.Database("db");
  console.log("created");

  db.run(
    "create table if not exists users (id integer primary key autoincrement, name text not null, email text not null)"
  );

  console.log("seeding");
  db.run("insert into users (name, email) values ('vp', 'vp@email.com')");
  db.run("insert into users (name, email) values ('kk', 'kk@email.com')");
  db.close();
  return { db };
}
