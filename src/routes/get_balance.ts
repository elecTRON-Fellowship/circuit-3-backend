import express from "express";
import { getBalance } from "../controllers/get_balance";

const router = express.Router();

router.post("/gbalance", getBalance);

export default module.exports = { router };
