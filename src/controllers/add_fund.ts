import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";

require("dotenv").config();

export const addFunds = async (req: express.Request, res: express.Response) => {
  // get data from req
  const data = req.body;

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "post",
    "/v1/account/deposit",
    salt,
    accessKey,
    secretKey,
    JSON.stringify(data),
  );

  try {
    await axios.post("https://sandboxapi.rapyd.net/v1/account/deposit", data, {
      headers: {
        "content-type": "application/json",
        access_key: accessKey,
        salt: salt,
        timestamp: timestamp,
        signature: signature,
      },
    });
  } catch (err) {
    await res.status(400).json({
      error: err,
      message: "Invalid data passed",
    });
  }
};
