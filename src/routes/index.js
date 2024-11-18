const checkHealth = require("./health");

const version = "/api/v1";

module.exports = function (app) {
   app.use(`${version}`, checkHealth);
};
