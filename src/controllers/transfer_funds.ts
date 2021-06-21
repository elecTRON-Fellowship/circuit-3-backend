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
    // initiate transfer
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

    // automagically accept the transaction
    let nreq = express.request;
    let nres = express.response;
    let txID = result.data.data.id;
    nreq.body = {
      id: txID,
      metadata: {
        merchant_defined: "accepted",
      },
      status: "accept",
    };
    try {
      await setTransferResponseInternal(nreq, nres);
    } catch (err) {
      console.log(err);
      await res.status(400).json({
        error: err,
        message: "There was an error transferring funds",
      });
      return;
    }
    await res.status(200).json({
      data: result.data,
      message: "Successfully transferred",
    });
  } catch (err) {
    await res.status(400).json({
      error: err,
      message: "Invalid data passed",
    });
    return;
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
    const result = await axios.post(
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
    await res.status(200).json({
      data: result.data,
      message: "Transfer response set successfully",
    });
  } catch (err) {
    await res.status(400).json({
      error: err,
      message: "There was an error updating the status",
    });
  }
};

export const setTransferResponseInternal = async (
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
    await axios.post(
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
  } catch (err) {
    await res.json({
      error: err,
      message: "There was an error updating the status",
    });
  }
};
