import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";
import { db } from "../util/firebase";
import { addFundsInternal } from "./add_funds";

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
    // create wallet
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

    // add funds to play with
    let nreq = express.request;
    let nres = express.response;
    nreq.body = {
      ewallet: result.data.data.id,
      amount: "10000",
      currency: "INR",
      metadata: {
        merchant_defined: true,
      },
    };
    try {
      await addFundsInternal(nreq, nres);
    } catch (err) {
      await res.status(400).json({
        error: err,
        message: "Invalid data passed",
      });
      return;
    }
    if (!userID) {
      res.json("Provide a userID to store ewallet to firebase");
      return;
    }
    try {
      await db.collection("users").doc(userID).update({
        walletID: result.data.data.id,
      });
    } catch (err) {
      await res.status(400).json({
        error: err,
        message: "Error storing wallet id to firebase",
      });
      return;
    }
    await res.json({
      data: result.data,
      message: "Successfully created",
    });
  } catch (err) {
    await res.status(400).json({
      error: err,
      message: "Invalid data passed",
    });
    return;
  }
};
