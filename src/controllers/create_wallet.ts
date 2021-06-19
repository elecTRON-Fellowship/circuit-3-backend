import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";
import { db } from "../util/firebase";

require("dotenv").config();

export const createWallet = async (
  req: express.Request,
  res: express.Response,
) => {
  // get data from req
  const data = req.body;
  // get user id from req header
  const userID = await req.header("user_id");

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "post",
    "/v1/user",
    salt,
    accessKey,
    secretKey,
    JSON.stringify(data),
  );
  try {
    const result: any = await axios.post(
      "https://sandboxapi.rapyd.net/v1/user",
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
    await res.json({
      data: result.data,
      message: "Successfully created",
    });
    if (!userID) {
      res.send("there was an error storing data in db");
      return;
    }
    console.log(result.data.data.id);
    try {
      await db.collection("users").doc(userID).update({
        walletData: result.data.data,
      });
    } catch (err) {
      console.log(err);
    }
  } catch (err) {
    await res.status(400).json({
      error: err,
      message: "Invalid data passed",
    });
  }
};
