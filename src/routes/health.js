const express = require("express");
const { AppError } = require("../libs/error");
const router = express.Router();

router.get("/", (req, res) => {
   throw new AppError("Server Error", 500);
});

router.get("/health", (req, res) => {
   res.status(200).json({ message: "Server is healthy 100%", version: "1.0.0" });
});

module.exports = router;
