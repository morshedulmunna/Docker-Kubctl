const { AppError } = require("../../libs/error");

const router = require("express").Router();

router.post("/register", (req, res) => {
  const { email, password } = req.body;
});
