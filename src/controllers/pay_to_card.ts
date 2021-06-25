import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";

export const payToCard = async (
  req: express.Request,
  res: express.Response,
) => {
  // get data from request body
  const { ewallet } = req.body;

  const body = {
    beneficiary: {
      email: "sandboxtest@rapyd.net",
      card_number: "4111111111111111",
      card_expiration_month: "11",
      card_expiration_year: "2021",
      company_name: "Test Company",
      postcode: "56789",
    },
    beneficiary_country: "US",
    beneficiary_entity_type: "individual",
    description: "desc1562234632",
    payout_method_type: "us_atmdebit_card",
    // sender's data
    ewallet: ewallet,
    metadata: {
      merchant_defined: true,
    },
    payout_amount: "2.66",
    payout_currency: "USD",
    sender: {
      first_name: "John",
      last_name: "Doe",
      identification_type: "work_permit",
      identification_value: "asdasd123123",
      phone_number: "19019019011",
      occupation: "plumber",
      source_of_income: "business",
      date_of_birth: "11/12/1913",
      address: "1 Main Street",
      purpose_code: "investment_income",
      beneficiary_relationship: "spouse",
    },
    sender_country: "US",
    sender_currency: "USD",
    sender_entity_type: "individual",
  };

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "post",
    "/v1/payouts",
    salt,
    accessKey,
    secretKey,
    JSON.stringify(body),
  );
  try {
    const result = await axios.post(
      "https://sandboxapi.rapyd.net/v1/payouts",
      body,
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
      message: "Funds added successfully",
    });
  } catch (err) {
    await res.status(400).json({
      error: err,
      message: "Invalid data passed",
    });
  }
};
