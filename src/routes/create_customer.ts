import express from "express";
import { createCustomer } from "../controllers/create_customer";

const router = express.Router();

router.post("/ccustomer", createCustomer);

export default module.exports = { router };
