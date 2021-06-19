import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";
// import { db } from "../util/firebase";

require("dotenv").config();

export const createTransfer = async (
  req: express.Request,
  res: express.Response,
) => {
  // get data from req
  const data = req.body;
  // const senderID = await req.header("sender_id");
  // const receiverID = await req.header("receiver_id");

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "post",
    "/v1/account/transfer",
    salt,
    accessKey,
    secretKey,
    JSON.stringify(data),
  );

  try {
    const result: any = await axios.post(
      "https://sandboxapi.rapyd.net/v1/account/transfer",
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
  } catch (err) {
    await res.status(400).json({
      error: err,
      message: "Invalid data passed",
    });
  }
};

export const setTransferResponse = async (
  req: express.Request,
  res: express.Response,
) => {
  // get data from req
  const data = req.body;

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "post",
    "/v1/account/transfer/response",
    salt,
    accessKey,
    secretKey,
    JSON.stringify(data),
  );

  try {
    const result: any = await axios.post(
      "https://sandboxapi.rapyd.net/v1/account/transfer/response",
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
      message: "Status updated",
    });
  } catch (err) {
    await res.json({
      error: err,
      message: "There was an error updating the status",
    });
  }
};
