const users = JSON.stringify([{ name: "vp", email: "vp@vp.com" }]);

export function handler(req, res, query) {
  console.log("query received", query); // {abc: 123, bva: 282}
  res.statusCode = 200;
  res.setHeader("Content-Type", "application/json");
  res.end(users);
}
