import express from "express";
import { listPayMethods } from "../controllers/list_payout_method";

const router = express.Router();

router.post("/lpaymethods", listPayMethods);

export default module.exports = { router };
