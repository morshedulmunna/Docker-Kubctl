require("dotenv").config();
const express = require("express");
const path = require("path");
const cookieParser = require("cookie-parser");
const { globalErrorHandler, AppError } = require("./src/libs/error");
const morgan = require("morgan");
const errorLogStream = require("./src/libs");
const routes = require("./src/routes");
const { DBconnectionHandling } = require("./src/configs/DB.config");

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use("/uploads", express.static("uploads"));

routes(app);

// Middleware to catch undefined routes
app.all("/*", (req, res, next) => {
   const error = new AppError(`Can't find ${req.originalUrl} on this server!`, 404);
   next(error);
});

// Configure Morgan to log errors
app.use(
   morgan("combined", {
      stream: errorLogStream,
      skip: function (req, res) {
         return res.statusCode < 400;
      },
   })
);

// Global error handling middleware
app.use(globalErrorHandler);
DBconnectionHandling();

module.exports = app;
