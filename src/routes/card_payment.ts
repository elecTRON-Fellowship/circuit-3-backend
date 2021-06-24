import express from "express";
import { cardPayment } from "../controllers/card_payment";

const router = express.Router();

router.post("/cpay", cardPayment);

export default module.exports = { router };
