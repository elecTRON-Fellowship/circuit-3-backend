import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";

require("dotenv").config();

export const listPayMethods = async (
  req: express.Request,
  res: express.Response,
) => {
  // get data from req
  //const { category } = req.body;
  const { beneficiaryCountry } = req.body;
  const { payoutCurrency } = req.body;

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "get",
    `/v1/payouts/supported_types?beneficiary_country=${beneficiaryCountry}&payout_currency=${payoutCurrency}`,
    salt,
    accessKey,
    secretKey,
    "",
  );
  try {
    const result = await axios.get(`https://sandboxapi.rapyd.net/v1/payouts/supported_types?beneficiary_country=${beneficiaryCountry}&payout_currency=${payoutCurrency}`, {
      headers: {
        "content-type": "application/json",
        access_key: accessKey,
        salt: salt,
        timestamp: timestamp,
        signature: signature,
      },
    });
    await res.status(200).json({
      data: result.data,
      message: "List payout methods fetch successfully",
    })
  } catch (err) {
    console.log(err);
  }
};
