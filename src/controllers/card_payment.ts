import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";

require("dotenv").config();

export const cardPayment = async (
  req: express.Request,
  res: express.Response,
) => {
  // take ewallet id from the request
  const { ewalletID } = req.body;

  // request body
  const data = {
    amount: 999,
    currency: "INR",
    merchant_reference_id: "first",
    payment_method: {
      type: "in_debit_visa_card",
      fields: {
        number: "4111111111111111",
        expiration_month: "11",
        expiration_year: "21",
        cvv: "123",
        name: "Jane Doe",
      },
    },
    ewallets: [
      {
        ewallet: ewalletID,
        percentage: 100,
      },
    ],
    metadeta: {
      merchant_defined: "created",
    },
    capture: true,
  };
  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "post",
    "/v1/payments",
    salt,
    accessKey,
    secretKey,
    JSON.stringify(data),
  );

  try {
    const result = await axios.post(
      "https://sandboxapi.rapyd.net/v1/payments",
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
      message: "Card Payment successfully executed",
    });
  } catch (err) {
    await res.json({
      error: err,
      message: "Failed to create customer",
    });
  }
};
