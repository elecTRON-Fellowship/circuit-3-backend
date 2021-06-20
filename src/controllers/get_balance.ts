import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";

require("dotenv").config();

export const getBalance = async (
  req: express.Request,
  res: express.Response,
) => {
  // get data from request body
  const { ewalletID } = req.body;

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "get",
    `/v1/user/${ewalletID}/accounts`,
    salt,
    accessKey,
    secretKey,
    "",
  );
  try {
    // get balance
    const result: any = await axios.get(
      `https://sandboxapi.rapyd.net/v1/user/${ewalletID}/accounts`,
      {
        headers: {
          "content-type": "application/json",
          access_key: accessKey,
          salt: salt,
          timestamp: timestamp,
          signature: signature,
        },
      },
    );
    await res.json({
      data: result.data,
      message: "Successfully fetched balance",
    });
  } catch (err) {
    await res.status(400).json({
      error: err,
      message: "Invalid data passed",
    });
  }
};
