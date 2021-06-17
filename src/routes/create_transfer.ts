import express from "express";
import { createTransfer } from "../controllers/create_transfer";

const router = express.Router();

router.post("/ctransfer", createTransfer);

export default module.exports = { router };
