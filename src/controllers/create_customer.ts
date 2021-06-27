import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";
import { db } from "../util/firebase";

require("dotenv").config();

export const createCustomer = async (
  req: express.Request,
  res: express.Response,
) => {
  // get data from req
  const { ewallet } = req.body;
  const { name } = req.body;
  const { phoneNumber } = req.body;

  // get user id from req header
  const userID = req.header("user_id");

  const data = {
    name: name,
    email: "hackathon@rapyd.net",
    ewallet: ewallet,
    invoice_prefix: "JD-",
    metadata: {
      merchant_defined: true,
    },
    payment_method: {
      type: "in_visa_debit_card",
      fields: {
        name: name,
        number: "4111111111111111",
        expiration_month: "10",
        expiration_year: "23",
        cvv: "123",
      },
      complete_payment_url: "https://raven.herokuapp.com/",
      error_payment_url: "https://raven.herokuapp.com/",
    },
    phone_number: phoneNumber,
  };

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "post",
    "/v1/customers",
    salt,
    accessKey,
    secretKey,
    JSON.stringify(data),
  );

  try {
    const result = await axios.post(
      "https://sandboxapi.rapyd.net/v1/customers",
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
    if (!userID) {
      res.json("Provide a userID to store ewallet to firebase");
      return;
    }
    try {
      await db.collection("users").doc(userID).update({
        customerID: result.data.data.id,
      });
    } catch (err) {
      await res.status(400).json({
        error: err,
        message: "Error storing wallet id to firebase",
      });
      return;
    }
    await res.status(200).json({
      data: result.data,
      message: "Customer created successfully",
    });
  } catch (err) {
    console.log(err);
    await res.status(400).json({
      error: err,
      message: "Failed to create customer",
    });
  }
};
