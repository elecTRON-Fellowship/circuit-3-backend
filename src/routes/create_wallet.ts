import express from "express"
import { createWallet } from "../controllers/create_wallet"

const router = express.Router()

router.post("/cwallet", createWallet)

export default module.exports = {router}