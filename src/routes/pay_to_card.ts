import express from "express";
import { payToCard } from "../controllers/pay_to_card";

const router = express.Router();

router.post("/ptocard", payToCard);

export default module.exports = { router };
