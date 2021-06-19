import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";

require("dotenv").config();

export const retreiveBalance = async (
  _req: express.Request,
  res: express.Response,
) => {
  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "get",
    "/v1/user/ewallet_29b0a00f91541ca6783cd610029c5de6/accounts",
    salt,
    accessKey,
    secretKey,
    "",
  );
  try {
    const result = await axios.get(
      "https://sandboxapi.rapyd.net/v1/user/ewallet_29b0a00f91541ca6783cd610029c5de6/accounts",
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
    });
  } catch (err) {
    console.log(err);
    await res.status(400).json({
      error: err,
      message: "Invalid wallet id passed",
    });
  }
};
