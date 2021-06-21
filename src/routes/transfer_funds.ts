import express from "express";
import {
  createTransfer,
  setTransferResponse,
} from "../controllers/transfer_funds";

const router = express.Router();

router.post("/ctransfer", createTransfer);

router.post("/sresponse", setTransferResponse);

export default module.exports = { router };
