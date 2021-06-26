import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";

require("dotenv").config();

export const createCheckout = async (
  req: express.Request,
  res: express.Response,
) => {
  // get data from req
  const { customer } = req.body;
  const { ewallet } = req.body;
  const { amount } = req.body;
  const data = {
    amount: amount,
    country: "IN",
    currency: "INR",
    customer: customer,
    cardholder_preferred_currency: true,
    language: "en",
    metadata: {
      merchant_defined: true,
    },
    payment_method_types_include: [
      "in_visa_credit_card",
      "in_visa_debit_card",
      "in_phonepe_ewallet",
      "in_airtelmoney_ewallet",
      "in_paytm_ewallet",
      "in_rupay_debit_card",
      "in_credit_mastercard_card",
    ],
    ewallet: ewallet,
    payment_method_types_exclude: []
  };

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
    await res.status(200).json({
      data: result.data.data.redirect_url,
      message: "Checkout created successfully",
    });
  } catch (err) {
    await res.status(400).json({
      error: err,
      message: "There was an error creating checkout",
    });
  }
};
