import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";

require("dotenv").config();

export const createCheckout = async (
  _req: express.Request,
  _res: express.Response,
) => {
  // get data from req
  const data = _req.body;

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "post",
    "/v1/checkout",
    salt,
    accessKey,
    secretKey,
    JSON.stringify(data),
  );

  try {
    const result: any = await axios.post(
      "https://sandboxapi.rapyd.net/v1/checkout",
      data,
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
    await _res.json({
      data: result.data,
      message: "Checkout created successfully",
    });
  } catch (err) {
    await _res.json({
      error: err,
      message: "There was an error creating checkout",
    });
  }
};
