// Pure Node server implementation
// App listens on 1337 and waits for connections
// connects to the local DB (sqlite or mysql)

import { handler } from "./handler.js";
import { initDB } from "./db.js";
import url from "url";
import http from "node:http";

const hostname = "127.0.0.1";
const port = 1337;

//init DB

console.log("main");
const { db } = initDB();
// clean tables
// db.run("truncate if exists table users");
// // seed tables
// db.run("insert into users (name, email) values ('vp', 'vp@email.com')");
// // close db
// seedDB();
// db.close();

const server = http.createServer((req, res) => {
  const { pathname, query } = url.parse(req.url, { parseQueryString: true });
  if (pathname === "/users") {
    handler(req, res, query);
  } else if (pathname === "/") {
    res.statusCode = 200;
    res.setHeader("Content-Type", "application/json");
    res.end(`{"msg": "ok-root"}`);
  } else {
    res.statusCode = 404;
    res.end(`{"err": "not-found"}`);
  }
});

server.listen(port, hostname, () => {
  console.log("server running");
});

console.log("hi!");
