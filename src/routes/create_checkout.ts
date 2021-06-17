import express from "express";
import { createCheckout } from "../controllers/create_checkout";

const router = express.Router();

router.post("/ccheckout", createCheckout);

export default module.exports = { router };
