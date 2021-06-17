import express from "express";
import { retreiveBalance } from "../controllers/retreive_balance";

const router = express.Router();

router.get("/rbalance", retreiveBalance);

export default module.exports = { router };
